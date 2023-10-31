package reader

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type FileReader struct {
	file *os.File
	data []string
}

func NewFileReader(fileName string) (*FileReader, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	fr := &FileReader{
		file: file,
	}

	fr.setData()

	return fr, nil
}

func (r *FileReader) ReadOne() (string, error) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := random.Intn(len(r.data))
	randomString := r.data[randomIndex]

	return randomString, nil
}

func (r *FileReader) setData() {
	scanner := bufio.NewScanner(r.file)

	r.data = make([]string, 0)
	for scanner.Scan() {
		r.data = append(r.data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
