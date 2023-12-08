package day1

import (
	"fmt"
	"maps"
	"sort"
	"strconv"
)

var atoi = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

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

func Part2(puzzleInput []string) {
	tally := 0
	for _, line := range puzzleInput {

		ints := part2ParseInts(line)
		strs := part2ParseStrings(line)
		tally += part2CombineMaps(ints, strs)

	}
	fmt.Println(tally)
}

func part2ParseInts(line string) map[int]int {
	acc := make(map[int]int)
	for i, char := range line {
		c, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}
		acc[i] = c
	}
	return acc
}

func part2ParseStrings(line string) map[int]int {
	acc := make(map[int]int)
	for strnum, intnum := range atoi {
		idx := indexOfSubstring(line, strnum)
		for _, index := range idx {
			acc[index] = intnum
		}

	}
	return acc
}
func indexOfSubstring(str, subStr string) []int {
	var acc []int
	for i := range str {
		if i+len(subStr) > len(str) {
			continue
		}
		if str[i:i+len(subStr)] == subStr {
			acc = append(acc, i)
		}
	}
	return acc
}
func part2CombineMaps(intNum, strNum map[int]int) int {
	var acc []int
	maps.Copy(intNum, strNum)
	//Sort intNum by Keys
	keys := make([]int, 0, len(intNum))
	for k := range intNum {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		acc = append(acc, intNum[k])
	}
	return (10 * acc[0]) + acc[len(acc)-1]

}
