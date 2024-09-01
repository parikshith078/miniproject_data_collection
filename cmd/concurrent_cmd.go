package cmd

import (
	"fmt"
	"mini/data_mine/llm"
	"mini/data_mine/utils"
	"sync"
)

// Takes chapter txt files as input
func ConcurrentlyRunFullCycle(txtFilesFolder string) error {
	fmt.Println("Starting full Cycle")
	fmt.Println("Generating Topics...")
	topicsFolderPath, err := ConcurrentlyGenerateTopicsDB(txtFilesFolder)
	if err != nil {
		return err
	}
	fmt.Println("Generating Samples...")
	sampleFolderPath, err := ConcurrentlyGenerateQuestionSamples(topicsFolderPath)
	if err != nil {
		return err
	}
	fmt.Println("Aggregating samples...")
	utils.AggregateSamples("./data/aggregated-samples", sampleFolderPath)
	return nil
}

func ConcurrentlyGenerateQuestionSamples(topicsFolder string) (string, error) {
	// Get all .json files from the folder
	files, err := utils.GetFileFromFolder(topicsFolder, ".json")
	if err != nil {
		return "", err
	}

	// Create a folder to store the generated question samples
	folderPath, err := utils.CreateLogFolder("./data/samples-db")
	if err != nil {
		return "", err
	}

	// Initialize a WaitGroup to manage concurrency
	var wg sync.WaitGroup

	// A channel to collect errors from goroutines
	errChan := make(chan error, len(files))

	// Iterate over each file
	for _, file := range files {
		fmt.Println("Working on:", file)

		// Read the JSON file into a Topics struct
		res, err := utils.ReadJSONFile[llm.Topics](file)
		if err != nil {
			return "", err
		}

		// Extract the file name (without extension)
		fileName := utils.ExtractFileName(file)

		// Iterate over each topic in the JSON file
		for j, topic := range res.Topics {
			// Increment the WaitGroup counter
			wg.Add(1)

			// Run the question sample generation in a separate goroutine
			go func(fileName string, j int, topic llm.Topic) {
				defer wg.Done()

				// Generate question samples from the topic
				contextString := topic.SubTopic + "\n" + topic.Content
				questionSamples := llm.GenerateQuestionSamples(contextString)

				// Define the file path where the result will be saved
				filePath := fmt.Sprintf("%s/%s_%d.json", folderPath, fileName, j)

				// Save the generated question samples to a JSON file
				err := utils.SaveResultToJSONFile(filePath, questionSamples)
				if err != nil {
					// Send any error encountered to the error channel
					errChan <- err
					return
				}

				// Print the path of the created file
				fmt.Println("Created file:", filePath)
			}(fileName, j, topic)
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the error channel after all goroutines have completed
	close(errChan)

	// Check if there were any errors
	for err := range errChan {
		if err != nil {
			return "", err
		}
	}

	return folderPath, nil
}

func ConcurrentlyGenerateQuestionSamples2to1(topicsFolder string) (string, error) {
	// Get all .json files from the folder
	files, err := utils.GetFileFromFolder(topicsFolder, ".json")
	if err != nil {
		return "", err
	}

	// Create a folder to store the generated question samples
	folderPath, err := utils.CreateLogFolder("./data/samples-db")
	if err != nil {
		return "", err
	}

	// Initialize a WaitGroup to manage concurrency
	var wg sync.WaitGroup

	// A channel to collect errors from goroutines
	errChan := make(chan error, len(files))

	// Iterate over each file
	for _, file := range files {
		fmt.Println("Working on:", file)

		// Read the JSON file into a Topics struct
		res, err := utils.ReadJSONFile[llm.Topics](file)
    fmt.Println("res: ", len(res.Topics));
    
		if err != nil {
			return "", err
		}

		// Extract the file name (without extension)
		fileName := utils.ExtractFileName(file)

		// Iterate over topics in pairs
		for j := 0; j < len(res.Topics); j += 2 {
			fmt.Println("j: ", j)
			fmt.Println("topic name: ", res.Topics[j].SubTopic)
			// Increment the WaitGroup counter
			wg.Add(1)

			// Run the question sample generation in a separate goroutine
			go func(fileName string, j int, topic1, topic2 llm.Topic) {
				defer wg.Done()

				// Stack content of two topics
				contextString := topic1.SubTopic + "\n" + topic1.Content
				if topic2.SubTopic != "" && topic2.Content != "" {
					contextString += "\n" + topic2.SubTopic + "\n" + topic2.Content
				}

				// Generate question samples from the stacked content
				questionSamples := llm.GenerateQuestionSamples(contextString)

				// Define the file path where the result will be saved
				filePath := fmt.Sprintf("%s/%s_%d.json", folderPath, fileName, j)

				// Save the generated question samples to a JSON file
				err := utils.SaveResultToJSONFile(filePath, questionSamples)
				if err != nil {
					// Send any error encountered to the error channel
					errChan <- err
					return
				}

				// Print the path of the created file
				fmt.Println("Created file:", filePath)
			}(fileName, j, res.Topics[j], getTopicOrNil(res.Topics, j+1))
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the error channel after all goroutines have completed
	close(errChan)

	// Check if there were any errors
	for err := range errChan {
		if err != nil {
			return "", err
		}
	}

	return folderPath, nil
}

// getTopicOrNil returns the topic at index i or an empty Topic if out of bounds
func getTopicOrNil(topics []llm.Topic, i int) llm.Topic {
	if i >= len(topics) {
		return llm.Topic{}
	}
	return topics[i]
}

func ConcurrentlyGenerateTopicsDB(txtFilesFolder string) (string, error) {
	// Get all .txt files from the folder
	files, err := utils.GetFileFromFolder(txtFilesFolder, ".txt")
	if err != nil {
		return "", err
	}

	// Create a folder to store the generated topics
	folderPath, err := utils.CreateLogFolder("./data/topics-db")
	if err != nil {
		return "", err
	}

	// Initialize a WaitGroup to manage concurrency
	var wg sync.WaitGroup

	// A channel to collect errors from goroutines
	errChan := make(chan error, len(files))

	// Iterate over each file
	for _, file := range files {
		fileName := utils.ExtractFileName(file)

		// Read the content of the file into a string
		content, err := utils.ReadFileToString(file)
		if err != nil {
			return "", err
		}

		// Split the content into chunks
		chunk := utils.SplitStringByWordsLimit(content, 2000)

		// Iterate over each chunk
		for i, contextString := range chunk {
			// Increment the WaitGroup counter
			wg.Add(1)

			// Run the topic generation in a separate goroutine
			go func(fileName string, i int, contextString string) {
				defer wg.Done()

				// Generate topics from the chunk using the API call
				res := llm.GenerateTopics(contextString)

				// Define the file path where the result will be saved
				filePath := fmt.Sprintf("%s/%s_%d.json", folderPath, fileName, i)

				// Save the generated topics to a JSON file
				err := utils.SaveResultToJSONFile(filePath, res)
				if err != nil {
					// Send any error encountered to the error channel
					errChan <- err
					return
				}

				// Print the path of the created file
				fmt.Println("Created file:", filePath)
			}(fileName, i, contextString)
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the error channel after all goroutines have completed
	close(errChan)

	// Check if there were any errors
	for err := range errChan {
		if err != nil {
			return "", err
		}
	}

	return folderPath, nil
}
