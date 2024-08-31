package llm

import (
	"context"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

func GenerateResponseWithSchema[T any](systemPrompt, userPrompt, shemaName string) T { // userPrompt is context here

	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)
	ctx := context.Background()

	var result T
	schema, err := jsonschema.GenerateSchemaForType(result)
	if err != nil {
		log.Fatalf("GenerateSchemaForType error: %v", err)
	}
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: userPrompt,
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:   shemaName,
				Schema: schema,
				Strict: true,
			},
		},
	})
	if err != nil {
		log.Fatalf("CreateChatCompletion error: %v", err)
	}
	err = schema.Unmarshal(resp.Choices[0].Message.Content, &result)
	if err != nil {
		log.Fatalf("Unmarshal schema error: %v", err)
	}
	return result
}
