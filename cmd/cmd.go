package cmd

import (
	"fmt"
	"mini/data_mine/llm"
	"mini/data_mine/utils"
)

// Takes chapter txt files as input
func RunFullCycle(txtFilesFolder string) error {
	fmt.Println("Starting full Cycle")
	fmt.Println("Generating Topics...")
	topicsFolderPath, err := GenerateTopicsDB(txtFilesFolder)
	if err != nil {
		return err
	}
	fmt.Println("Generating Samples...")
	sampleFolderPath, err := GenerateQuestionSamples(topicsFolderPath)
	if err != nil {
		return err
	}
	fmt.Println("Aggregating samples...")
	utils.AggregateSamples("./data/aggregated-samples", sampleFolderPath)
	return nil
}

func GenerateQuestionSamples(topicsFolder string) (string, error) { // folder containg topics db json files
	files, err := utils.GetFileFromFolder(topicsFolder, ".json")
	if err != nil {
		return "", err
	}
	folderPath, err := utils.CreateLogFolder("./data/samples-db")
	if err != nil {
		return "", err
	}
	for _, file := range files {
		fmt.Println("Working on: ", file)
		res, err := utils.ReadJSONFile[llm.Topics](file)
		if err != nil {
			return "", err
		}
		fileName := utils.ExtractFileName(file)
		for j, topic := range res.Topics {
			contextString := topic.SubTopic + "\n" + topic.Content
			questionSamples := llm.GenerateQuestionSamples(contextString)
			filePath := fmt.Sprintf("%s/%s_%d.json", folderPath, fileName, j)
			err := utils.SaveResultToJSONFile(filePath, questionSamples)
			if err != nil {
				return "", err
			}
		}
	}
	return folderPath, nil
}

func GenerateTopicsDB(txtFilesFolder string) (string, error) {

	res, err := utils.GetFileFromFolder(txtFilesFolder, ".txt")
	if err != nil {
		return "", nil
	}
	folderPath, err := utils.CreateLogFolder("./data/topics-db")
	if err != nil {
		return "", nil
	}
	for _, file := range res {
		fileName := utils.ExtractFileName(file)
		content, err := utils.ReadFileToString(file)
		if err != nil {
			return "", nil
		}
		chunk := utils.SplitStringByWordsLimit(content, 2000)
		for i, contextString := range chunk {
			filePath := fmt.Sprintf("%s/%s_%d.json", folderPath, fileName, i)
			res := llm.GenerateTopics(contextString)
			err := utils.SaveResultToJSONFile(filePath, res)
			if err != nil {
				return "", nil
			}
			fmt.Println("Created files: ", filePath)
		}
	}
	return folderPath, nil
}
