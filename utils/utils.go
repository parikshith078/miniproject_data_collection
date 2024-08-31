package utils

import (
	"encoding/json"
	"fmt"
	"os"
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
