package main

import (
	"context"
	"fmt"
	"os"

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
	// 构建聊天完成请求，设置请求的模型和消息内容
	req := model.CreateChatCompletionRequest{
		// 将推理接入点 <Model>替换为 Model ID
		Model: "doubao-seed-1-6-250615",
		Messages: []*model.ChatCompletionMessage{
			{
				// 消息的角色为用户
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String("你好，请分析下明太祖朱元璋的生平，从当时人民的生活水平到后世对他的评价，完整的评价一下他的一生，并且用中文回复"),
				},
			},
		},
	}

	// 发送聊天完成请求，并将结果存储在 resp 中，将可能出现的错误存储在 err 中
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		// 若出现错误，打印错误信息并终止程序
		fmt.Printf("standard chat error: %v\n", err)
		return
	}
	// 打印聊天完成请求的响应结果
	fmt.Println(*resp.Choices[0].Message.Content.StringValue)
}
