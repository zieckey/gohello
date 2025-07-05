package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

// 定义分块大小为 10000 字符，可根据实际情况调整
const chunkSize = 10000

func readLargeFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	contentBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	content := string(contentBytes)

	var chunks []string
	for i := 0; i < len(content); i += chunkSize {
		end := i + chunkSize
		if end > len(content) {
			end = len(content)
		}
		chunks = append(chunks, content[i:end])
	}
	return chunks, nil
}

func summarizeFileContent(client *arkruntime.Client, ctx context.Context, chunks []string) (string, error) {
	var summary string
	for _, chunk := range chunks {
		messages := []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String(fmt.Sprintf("请总结以下内容：\n%s", chunk)),
				},
			},
		}

		req := model.CreateChatCompletionRequest{
			Model:    "doubao-seed-1-6-250615",
			Messages: messages,
		}

		resp, err := client.CreateChatCompletion(ctx, req)
		if err != nil {
			return "", err
		}
		summary += *resp.Choices[0].Message.Content.StringValue + "\n"
	}
	return summary, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请提供要读取的文件路径作为参数")
		return
	}
	filePath := os.Args[1]

	client := arkruntime.NewClientWithApiKey(
		//通过 os.Getenv 从环境变量中获取 ARK_API_KEY
		os.Getenv("ARK_API_KEY"),
	)
	// 创建一个上下文，通常用于传递请求的上下文信息，如超时、取消等
	ctx := context.Background()

	// 读取大文件并分块
	chunks, err := readLargeFile(filePath)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}

	// 让 AI 总结文件内容
	fileSummary, err := summarizeFileContent(client, ctx, chunks)
	if err != nil {
		fmt.Printf("总结文件内容失败: %v\n", err)
		return
	}

	// 初始化消息列表，添加文件总结作为初始上下文
	messages := []*model.ChatCompletionMessage{
		{
			Role: model.ChatMessageRoleSystem,
			Content: &model.ChatCompletionMessageContent{
				StringValue: volcengine.String(fmt.Sprintf("以下是文件内容的总结：\n%s", fileSummary)),
			},
		},
	}

	fmt.Println("欢迎使用文件问答机器人！输入 'exit' 退出对话。")
	for {
		// 获取用户输入
		fmt.Print("你: ")
		var input string
		fmt.Scanln(&input)

		if strings.ToLower(input) == "exit" {
			break
		}

		// 添加用户消息到消息列表
		userMessage := &model.ChatCompletionMessage{
			// 消息的角色为用户
			Role: model.ChatMessageRoleUser,
			Content: &model.ChatCompletionMessageContent{
				StringValue: volcengine.String(input),
			},
		}
		messages = append(messages, userMessage)

		// 构建聊天完成请求，设置请求的模型和消息内容
		req := model.CreateChatCompletionRequest{
			// 将推理接入点 <Model>替换为 Model ID
			Model:    "doubao-seed-1-6-250615",
			Messages: messages,
		}

		// 发送聊天完成请求，并将结果存储在 resp 中，将可能出现的错误存储在 err 中
		resp, err := client.CreateChatCompletion(ctx, req)
		if err != nil {
			// 若出现错误，打印错误信息并继续下一轮对话
			fmt.Printf("standard chat error: %v\n", err)
			continue
		}

		// 添加机器人回复到消息列表
		botMessage := &model.ChatCompletionMessage{
			Role: model.ChatMessageRoleAssistant,
			Content: resp.Choices[0].Message.Content,
		}
		messages = append(messages, botMessage)

		// 打印聊天完成请求的响应结果
		fmt.Printf("机器人: %s\n", *resp.Choices[0].Message.Content.StringValue)
	}
}
