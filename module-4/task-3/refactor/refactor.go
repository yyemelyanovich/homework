package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(path string) ([]byte, error) {
	fd, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func countLines(raw []byte) int {
	count := 0
	for i := 0; i < len(raw); i++ {
		if raw[i] == '\n' {
			count++
		}
	}

	return count
}

func main() {
	path := "./README.md"

	raw, err := readFile(path)
	if err != nil {
		fmt.Println("Error occured: %s", err)
		return
	}

	fmt.Println(countLines(raw))
}
