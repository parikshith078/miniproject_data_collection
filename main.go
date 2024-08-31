package main

import (
	"fmt"
	"log"
	"mini/data_mine/llm"
	"mini/data_mine/utils"
)

func main() {
	// generateTopicsDB()
	// generateQuestionSamples()
	utils.AggregateSamples("./aggregated-samples", "./samples-db/gen_04-24-58-PM_31-08-24")
}

func generateQuestionSamples() {
	files, err := utils.GetFileFromFolder("./topics-db/gen_04-22-15-PM_31-08-24", ".json")
	if err != nil {
		panic(err)
	}
	folderPath, err := utils.CreateLogFolder("./samples-db")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println("Working on: ", file)
		res, err := utils.ReadJSONFile[llm.Topics](file)
		if err != nil {
			panic(err)
		}
		fileName := utils.ExtractFileName(file)
		for j, topic := range res.Topic {
			contextString := topic.SubTopic + "\n" + topic.Content
			questionSamples := llm.GenerateQuestionSamples(contextString)
			filePath := fmt.Sprintf("%s/%s_%d.json", folderPath, fileName, j)
			err := utils.SaveResultToJSONFile(filePath, questionSamples)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

func generateTopicsDB() {

	res, err := utils.GetFileFromFolder("./chapter_text_files", ".txt")
	if err != nil {
		panic(err)
	}
	path, err := utils.CreateLogFolder("./topics-db")
	if err != nil {
		panic(err)
	}
	for _, file := range res {
		fileName := utils.ExtractFileName(file)
		content, err := utils.ReadFileToString(file)
		if err != nil {
			panic(err)
		}
		chunk := utils.SplitStringByWordsLimit(content, 2000)
		for i, contextString := range chunk {
			filePath := fmt.Sprintf("%s/%s_%d.json", path, fileName, i)
			res := llm.GenerateTopics(contextString)
			err := utils.SaveResultToJSONFile(filePath, res)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Created files: ", filePath)
		}
	}

}
