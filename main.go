package main

import (
	"fmt"
	"log"
	"mini/data_mine/cmd"

	"mini/data_mine/utils"
	"time"
)

func main() {
	start := time.Now()

	// err := cmd.ConcurrentlyRunFullCycle("./data/chapter_text_files")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	sampleFolder, err := cmd.ConcurrentlyGenerateQuestionSamples2to1("./data/topics-db/gen_08-29-35-PM_01-09-24")
	if err != nil {
		log.Fatal(err)
	}
	err = utils.AggregateSamples("./data/aggregated-samples", sampleFolder)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Time taken %v to run\n", time.Since(start))
}
