package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func main() {
	client := arkruntime.NewClientWithApiKey(
		//通过 os.Getenv 从环境变量中获取 ARK_API_KEY
		os.Getenv("ARK_API_KEY"),
	)
	// 创建一个上下文，通常用于传递请求的上下文信息，如超时、取消等
	ctx := context.Background()
	
	// 初始化消息列表
	messages := []*model.ChatCompletionMessage{}

	fmt.Println("欢迎使用连续对话机器人！输入 'exit' 退出对话。")
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
			Model: "doubao-seed-1-6-250615",
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
