package js_file_chunker

import (
	"crypto/sha256"
	"io"
	//"io/ioutil"
	//"log"
	"os"
)

func FindBoundaries(fileName string) ([][]int, error) {
	var boundaries [][]int
	buf := make([]byte, WindowSize*1024*1024) //read 64MB at a time
	windowBuf := make([]byte, WindowSize)
	st := 0
	fileSize := 0

	file, err := os.Open(fileName)
	if err != nil {
		return [][]int{}, err
	}
	defer file.Close()

	for {
		n, err := file.Read(buf)
		if n > 0 {
			fileSize += n
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return [][]int{}, nil
			}
		}

		if n < WindowSize {
			boundaries = append(boundaries, []int{st, fileSize - 1}) // it is okay for the last chunked file to be less the size of MinChunkedFileSize
			return boundaries, nil
		}

		for i := 0; i+WindowSize-1 <= n-1; i++ {
			windowBuf = buf[i : i+WindowSize]
			rollingHash := sha256.Sum256(windowBuf)

			if CheckIfIsBoundary(rollingHash) {
				boundaries = append(boundaries, []int{st, fileSize - n + i + WindowSize})
				st = fileSize - n + i + WindowSize + 1
			}
		}
		file.Seek(int64(fileSize-WindowSize+1), 0) // if already reach end of file, then seek will not work
	}

	boundaries = append(boundaries, []int{st, fileSize - WindowSize})

	return boundaries, nil
}

func CheckIfIsBoundary(data [32]byte) bool {
	for i := 29; i <= 31; i++ {
		if data[i] != '0' {
			return false
		}
	}

	return true
}
