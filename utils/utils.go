package utils

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"mini/data_mine/llm"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Samples struct {
	Context  string `json:"input"`
	Question string `json:"ouput"`
}

func AggregateSamples(saveToFolderPath, folderPath string) error {
	files, err := GetFileFromFolder(folderPath, ".json")
	if err != nil {
		return err
	}

	samples := []Samples{}

	for _, file := range files {
		fmt.Println("Workign on file: ", file)
		res, err := ReadJSONFile[llm.Result](file)
		if err != nil {
			return err
		}
		for _, sample := range res.Samples {
			tmp := Samples{Context: sample.Context, Question: sample.Question}
			samples = append(samples, tmp)
		}
	}
	fmt.Println("Total sample size: ", len(samples))

	gen := getLogFileName() + ".json"

	filePath := filepath.Join(saveToFolderPath, gen)
	err = SaveResultToJSONFile(filePath, samples)
	if err != nil {
		return err
	}

	return nil
}

// ReadJSONFile reads the JSON file and unmarshals it into a Topics struct.
func ReadJSONFile[T any](filePath string) (*T, error) {

	// Read the file content
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Unmarshal the JSON content into the Topics struct
	var decodedData T
	err = json.Unmarshal(bytes, &decodedData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return &decodedData, nil
}

// saveResultToFile function takes a file path and a Result variable, and saves the JSON data to the file.
func SaveResultToJSONFile[T any](filePath string, result T) error {

	// Convert the Result struct to JSON
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Write the JSON data to the file
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON to file: %v", err)
	}

	return nil
}

func SplitStringByWordsLimit(input string, maxWords int) []string {
	words := strings.Fields(input) // Split the input string into words
	var result []string
	for i := 0; i < len(words); i += maxWords {
		end := i + maxWords
		if end > len(words) {
			end = len(words)
		}
		chunk := strings.Join(words[i:end], " ")
		result = append(result, chunk)
	}
	return result
}

// ReadFileToString reads the content of a text file and returns it as a string.
func ReadFileToString(filePath string) (string, error) {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert the content to a string and return
	return string(content), nil
}

func CreateLogFolder(basePath string) (string, error) {
	// Combine the base path with the new folder name
	folderName := getLogFileName()
	fullPath := filepath.Join(basePath, folderName)
	// Create the directory
	err := os.Mkdir(fullPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return fullPath, nil
}

func ExtractFileName(filePath string) string {
	// Extract the base name (e.g., "jesc105.txt" from "chapter_text_files/jesc105.txt")
	baseName := filepath.Base(filePath)
	// Remove the extension (e.g., "jesc105" from "jesc105.txt")
	fileName := strings.TrimSuffix(baseName, filepath.Ext(baseName))
	return fileName
}

func GetFileFromFolder(folderPath, fileType string) ([]string, error) {
	// Walk through the folder
	textFilePaths := []string{}
	err := filepath.Walk(folderPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the file has a .pdf extension
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), fileType) {
			textFilePaths = append(textFilePaths, path)
		}
		return nil
	})

	return textFilePaths, err
}

func getLogFileName() string {
	// Get the current time and format it as "03-04-05-PM_30-08-24"
	currentTime := time.Now().Format("03-04-05-PM_02-01-06")
	folderName := fmt.Sprintf("gen_%s", currentTime)
	return folderName
}
