package day2

import (
	"fmt"
	"strconv"
	"strings"
)

var validGame = map[string]int{
	"blue":  14,
	"red":   12,
	"green": 13,
}

type game struct {
	id       int
	subgames []subgame
}

type subgame struct {
	red   int
	blue  int
	green int
}

func (g *game) validateGames() bool {
	for _, sg := range g.subgames {
		if !(sg.green <= validGame["green"] && sg.red <= validGame["red"] && sg.blue <= validGame["blue"]) {
			return false
		}
	}
	return true
}

func (g *game) parseGame(line string) {
	t := strings.Split(line, ":")
	g.id, _ = strconv.Atoi(strings.ReplaceAll(t[0], "Game ", ""))
	g.subgames = parseSubGames(t[1])
}

func parseSubGames(g string) []subgame {
	var games []subgame
	subgames := strings.Split(g, ";")
	//subgames will look like ["3 blue, 4 red" "1 red , 2 green, 6 blue" "2 green"]
	for _, game := range subgames {
		var s subgame
		sg := strings.Split(game, ",")
		//sg will look like ["3 blue" "4 red"]
		for _, g := range sg {
			i := strings.Split(g, " ")
			switch i[2] {
			case "blue":
				s.blue, _ = strconv.Atoi(i[1])
			case "red":
				s.red, _ = strconv.Atoi(i[1])
			case "green":
				s.green, _ = strconv.Atoi(i[1])

			}
		}

		games = append(games, s)
	}
	return games
}
func Part1(puzzle []string) {
	tally := 0
	for _, line := range puzzle {
		var g game
		g.parseGame(line)
		if g.validateGames() {
			tally += g.id
		}
	}
	fmt.Println(tally)
}

func Part2(puzzle []string) {
	tally := 0
	for _, line := range puzzle {
		var g game
		g.parseGame(line)
		maxSubGames := findMax(g.subgames)
		power := maxSubGames.blue * maxSubGames.red * maxSubGames.green
		tally += power
	}
	fmt.Println(tally)

}

func findMax(s []subgame) subgame {
	var redMax, blueMax, greenMax int
	for _, game := range s {
		if game.blue > blueMax {
			blueMax = game.blue
		}
		if game.red > redMax {
			redMax = game.red
		}
		if game.green > greenMax {
			greenMax = game.green
		}
	}
	return subgame{red: redMax, blue: blueMax, green: greenMax}
}
