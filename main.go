package main

import (
	"day1"
	"fmt"
	"util"
)

func main() {
	fmt.Println("Day 1, Part 1 Test Result")
	day1Part1Test, _ := util.ParsePuzzleInput("day1/testInput")
	day1.Part1(day1Part1Test)
	fmt.Println("Day 1, Part Puzzle Result")
	day1Part1Input, _ := util.ParsePuzzleInput("day1/puzzleInput")
	day1.Part1(day1Part1Input)
	fmt.Println("Day 1, Part 1 Test Result")
	day1Part2Test, _ := util.ParsePuzzleInput("day1/part2TestInput")
	day1.Part2(day1Part2Test)
	fmt.Println("Day 1, Part 2 Test Result")
	day1Part2Input, _ := util.ParsePuzzleInput("day1/puzzleInput")
	day1.Part2(day1Part2Input)

}
