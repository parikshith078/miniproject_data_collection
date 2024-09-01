package main

import (
	"fmt"
	"log"
	"mini/data_mine/cmd"
	// "mini/data_mine/utils"
	"time"
)

func main() {
	start := time.Now()

	err := cmd.ConcurrentlyRunFullCycle("./data/chapter_text_files")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Time taken %v to run\n", time.Since(start))
}
