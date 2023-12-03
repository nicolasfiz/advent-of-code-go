package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/nicolasfiz/advent-of-code-go/days/utils"
)

var colorCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Day2(path string) int {
	var result int
	lines := utils.LinesReader(path)

	for _, line := range lines {
		gameNumber, gameSets := getDataLine(line)
		if validGame(gameSets) {
			result += gameNumber
		}
	}

	return result
}

func getDataLine(line string) (int, string) {
	// Game nº
	separatorGamePlays := strings.Split(line, ":")
	gameSeparator := strings.Fields(separatorGamePlays[0])
	gameNumber, err := strconv.Atoi(gameSeparator[1])
	if err != nil {
		log.Fatal("Cannot convert game number")
	}
	fmt.Printf("Game nº: %d\n", gameNumber)
	// Plays
	gameSets := separatorGamePlays[1]
	return gameNumber, gameSets
}

func validGame(game string) bool {
	sets := strings.Split(game, ";")
	for i, set := range sets {
		plays := strings.Split(set, ",")
		fmt.Printf("For set %d ->\n", i)
		for _, play := range plays {
			separatorNumberColor := strings.Fields(play)
			numberPlay, _ := strconv.Atoi(separatorNumberColor[0])
			colorPlay := separatorNumberColor[1]
			fmt.Printf("\t%d %s\n", numberPlay, colorPlay)
			if numberPlay > colorCubes[colorPlay] {
				return false
			}
		}
	}
	return true
}
