package js_file_chunker

import (
	"io/ioutil"
	"os"
	"strconv"
)

func MergeFile(inputDir, outputDir, fileName string, fileNums int) error {
	outPutFile, err := os.OpenFile(outputDir+"/"+fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer outPutFile.Close() //ignore err

	for i := 0; i < fileNums; i++ {
		inputFile, err := os.Open(inputDir + "/" + fileName + "_" + strconv.Itoa(i))
		if err != nil {
			return err
		}

		chunkedDatas, err := ioutil.ReadAll(inputFile)
		if err != nil {
			return err
		}
		inputFile.Close() //ignore err

		_, err = outPutFile.Write(chunkedDatas)
		if err != nil {
			return err
		}

	}

	return nil
}
