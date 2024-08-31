package main

import (
	"fmt"
	"log"
	"mini/data_mine/llm"
	"mini/data_mine/utils"
)

func main() {
	// generateTopicsDB()
}

func generateTopicsDB() {

	res, err := utils.GetTextFilePaths("./chapter_text_files")
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
