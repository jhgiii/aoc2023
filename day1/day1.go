package day1

import (
	"fmt"
	"strconv"
)

func Part1(puzzleInput []string) {
	tally := 0
	for _, line := range puzzleInput {
		r := parseIntegers(line)
		tally += r
	}
	fmt.Println(tally)
}

func parseIntegers(input string) int {
	var integers []int
	for _, char := range input {
		c, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}
		integers = append(integers, c)
	}
	intAsString := strconv.Itoa(integers[0]) + strconv.Itoa(integers[len(integers)-1])
	result, _ := strconv.Atoi(intAsString)
	return result
}
