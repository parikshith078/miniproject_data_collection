package utils

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func ReadJSONFile(filePath string) ([]string, error) {
	// Read the JSON file
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Define a slice to hold the decoded data
	var content []string

	// Unmarshal the JSON data into the slice
	err = json.Unmarshal(jsonData, &content)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return content, nil
}

// saveResultToFile function takes a file path and a Result variable, and saves the JSON data to the file.
func SaveResultToFile[T any](filePath string, result T) error {

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

func SplitStringByWords(input string, maxWords int) []string {
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

func GetTextFilePaths(folderPath string) ([]string, error) {
	// Walk through the folder
	textFilePaths := []string{}
	err := filepath.Walk(folderPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the file has a .pdf extension
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".txt") {
			textFilePaths = append(textFilePaths, path)
		}
		return nil
	})

	return textFilePaths, err
}
