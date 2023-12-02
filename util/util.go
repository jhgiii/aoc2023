package util

import (
	"bufio"
	"os"
)

func ParsePuzzleInput(filename string) ([]string, error) {
	var result []string
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}
	return result, nil
}
