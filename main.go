package main

import (
	"fmt"
	"mini/data_mine/llm"
)

func main() {
	response, err := llm.CallLLM()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Response from OpenAI:", response)
}
