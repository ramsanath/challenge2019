package data

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

func ReadCSV(relPath string) (contents [][]string, err error) {
	contents = make([][]string, 0)

	pwd, _ := os.Getwd()
	path := pwd + relPath

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		contents = append(contents, record)
	}

	return
}
