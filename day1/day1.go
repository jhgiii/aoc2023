package day1

import (
	"fmt"
	"maps"
	"sort"
	"strconv"
	"strings"
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

/*
For Part 2, I am continuing the string manipulation strategy and plan on leveraging the strings.Index() function to find the index of the word
form of each digit by iterating over a map[string]int where each key is the word-form form of each digit and the value is the integer value

	  The index returned, and the corresponding integer that is mapped to the word-form will be inserted into a map[int]int where the key is the index location
	  and the value is the corresponding integer representation of the digit.

	  For parsing Ints we will copy the part1 logic, but rather than use a []int to track, we will use a map[int]int, where the key is the index, in the line, and the
	  value is the integer value parsed.

	  We will then combine these maps using Go 1.21's maps.Copy() function. Sort this merged map by its keys (index where a digit was located), iterate through these
	  indicies, and return a []int where each element is the integer value found in order.

	  ie:

	  Given "xtwone3four", we will see an []int with [2 1 3 4]
	We can then simply accumulate (10 * []int[0] + []int[len(x)-1]
*/
func Part2(puzzleInput []string) {
	tally := 0
	for _, line := range puzzleInput {

		ints := part2ParseInts(line)
		strs := part2ParseStrings(line)
		res := part2CombineMaps(ints, strs)
		tally += (10*res[0] + res[1])

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
		if (strings.Index(line, strnum)) >= 0 {
			acc[strings.Index(line, strnum)] = intnum
		}
	}
	return acc
}

func part2CombineMaps(intNum, strNum map[int]int) []int {
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

	return []int{acc[0], acc[len(acc)-1]}

}

//func part2ConvertInts(ints []int) (int, error) {
// 	result := strconv.Itoa(ints[0]) + strconv.Itoa(ints[1])
// 	return strconv.Atoi(result)
// }
