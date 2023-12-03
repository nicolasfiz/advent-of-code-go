package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nicolasfiz/advent-of-code-go/days/utils"
)

func Day2Part2(path string) int {
	var result int
	lines := utils.LinesReader(path)

	for _, line := range lines {
		_, gameSets := getDataLine(line)
		power := gamePower(gameSets)
		result += power
	}

	return result
}

func gamePower(game string) int {
	var maxMap = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	var result int = 1
	sets := strings.Split(game, ";")
	for i, set := range sets {
		plays := strings.Split(set, ",")
		fmt.Printf("For set %d ->\n", i)
		for _, play := range plays {
			separatorNumberColor := strings.Fields(play)
			numberPlay, _ := strconv.Atoi(separatorNumberColor[0])
			colorPlay := separatorNumberColor[1]
			fmt.Printf("\t%d %s\n", numberPlay, colorPlay)
			if numberPlay > maxMap[colorPlay] {
				maxMap[colorPlay] = numberPlay
			}
		}
	}
	for _, v := range maxMap {
		if v != 0 {
			result *= v
		}
	}
	return result
}
