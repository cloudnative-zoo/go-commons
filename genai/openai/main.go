package main

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func main() {
	text := "What is the capital of the United States?"
	response, err := invokeChatGPTAPI("", "", text)
	if err != nil {
		panic(err)
	}
	println(response)
}

func invokeChatGPTAPI(url, token, text string) (string, error) {
	config := openai.DefaultConfig(token)
	config.BaseURL = url

	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gemini-1.5-flash",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, err
}
