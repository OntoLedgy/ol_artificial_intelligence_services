package large_language_model_services

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

func ConnectToChatGPT(authToken string) *openai.Client {
	client := openai.NewClient(
		authToken)

	return client

}

func CreateChatCompletion(
	client *openai.Client,
	requestMessage string) string {

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0301,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: requestMessage,
				},
			},
		},
	)

	if err != nil {
		errorMessage := fmt.Sprintf("\"ChatCompletion error: %v\n", err)
		return errorMessage
	}

	response := resp.Choices[0].Message.Content

	return response
}
