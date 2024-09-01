package main

import (
	"log"
	"mini/data_mine/cmd"
)

func main() {

	err := cmd.RunFullCycle("./data/chapter_text_files")
	if err != nil {
		log.Fatal(err)
	}
}
