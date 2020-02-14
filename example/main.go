package main

import (
	fileChunker "github.com/jscode017/go_content_defined_file_chunker"
	"log"
	"os"
)

func main() {
	if _, err := os.Stat("./input"); os.IsNotExist(err) {
		os.Mkdir("./input", 0766)
	}

	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		os.Mkdir("./output", 0766)
	}

	fileNums, err := fileChunker.ChunkFile("randfile", "./input") //first, create a randfile using dd in your terminal
	if err != nil {
		log.Fatal(err)
	}

	err = fileChunker.MergeFile("./input", "./output", "randfile", fileNums)
	if err != nil {
		log.Fatal(err)
	}
}
