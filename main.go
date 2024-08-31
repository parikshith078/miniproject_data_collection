package main

import (
	"fmt"
	"mini/data_mine/utils"
)

type Result struct {
	Samples []struct {
		Context  string `json:"context"`
		Question string `json:"question"`
	} `json:"samples"`
}

type Topics struct {
	Topic []struct {
		SubTopic string `json:"sub-topic"`
		Content  string `json:"Content"`
	} `json:"topic"`
}

func main() {
	res, err := utils.GetTextFilePaths("./chapter_text_files")
	if err != nil {
		panic(err)
	}
	for _, file := range res {
		content, err := utils.ReadFileToString(file)
		if err != nil {
			panic(err)
		}
		chunk := utils.SplitStringByWords(content, 2000)
		fmt.Println("chunk len: ", len(chunk))
	}

}

// contextFilePath := "./context-output.txt"
// contextContent, err := utils.ReadFileToString(contextFilePath)
// if err != nil {
// 	log.Fatalf("Error while ReadFileToString result: %v", err)
// }
// // systemPrompt := "Your are a helpful science tutor for highschool students. Given some content (part of a chapter), create multiple context and question (word range of question should be 5-12, word range of context should be 20-70). Context is section of the content from where you got the question."
// // schemaName := "question_generation"
//
// systemPromptForTopics := "Your are a helpful science tutor for highschool students. Given some content as context sort them based on topic and subtopics. Each subtopics should have content in them which has to be taken from the given context (Retain most of the content in the output). Your just sorting the content here so do not reduce the content in the output just sort them in the told order."
// schemaNameTopics := "topic_wise_sorting"
//
// result := llm.GenerateResponseWithSchema[Topics](systemPromptForTopics, contextContent, schemaNameTopics)
//
// outputFile := "output_topics.json"
// err = utils.SaveResultToFile(outputFile, result)
// if err != nil {
// 	log.Fatalf("Error while saving result: %v", err)
// }
