package main

import (
	"log"
	"mini/data_mine/llm"
	"mini/data_mine/utils"
)

type Result struct {
	Samples []struct {
		Context  string `json:"context"`
		Question string `json:"question"`
	} `json:"samples"`
}

func main() {
	systemPrompt := "Your are a helpful science tutor for highschool students. Given some content (part of a chapter), create multiple context and question (word range of question should be 5-12, word range of context should be 20-70). Context is section of the content from where you got the question."
	schemaName := "question_generation"
	result := llm.GenerateResponseWithSchema[Result](systemPrompt, "", schemaName)
	outputFile := "output.json"
	err := utils.SaveResultToFile(outputFile, result)
	if err != nil {
		log.Fatalf("Error while saving result: %v", err)
	}
}
