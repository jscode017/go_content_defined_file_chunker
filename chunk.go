package js_file_chunker

import (
	"errors"
	//"io/ioutil"
	"log"
	"os"
	"strconv"
)

func ChunkFile(inputFile, dir string) (int, error) {
	boundaries, err := FindBoundaries(inputFile)
	if err != nil {
		return 0, err
	}

	f, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	tmpFileIndex := 0
	processedBytes := 0
	for i := range boundaries {
		if boundaries[i][1] < boundaries[i][0] {
			return 0, errors.New("invalid boundary")
		}

		tmpFileSize := boundaries[i][1] - boundaries[i][0] + 1
		chunkedBytes := make([]byte, tmpFileSize)
		_, err := f.Read(chunkedBytes)
		if err != nil {
			return 0, err
		}

		outputFile, err := os.OpenFile(dir+"/"+inputFile+"_"+strconv.Itoa(tmpFileIndex), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return 0, err
		}
		tmpFileIndex++
		n2, err := outputFile.Write(chunkedBytes)
		if err != nil {
			return 0, err
		}
		outputFile.Close() //ignore err

		processedBytes += n2
	}

	log.Printf("processed %d bytes, chunked it to %d files\n", processedBytes, len(boundaries))

	return len(boundaries), nil
}
