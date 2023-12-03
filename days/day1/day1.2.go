package day1

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

var numbersMap = map[string]int{
	"zero":  0,
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

func Day1Part2(filePath string) int {
	scanner, file := createScanner(filePath)
	defer file.Close()

	total := fileReader(scanner, lineReaderDay2)

	return total
}

func lineReaderDay2(line string) int {
	var numbers []int

	for i, letter := range line {
		for word, digit := range numbersMap {
			if strings.HasPrefix(line[i:], word) {
				numbers = append(numbers, digit)
				break
			}
		}
		if unicode.IsNumber(letter) {
			number, error := strconv.Atoi(string(letter))
			if error != nil {
				log.Fatal("Could not convert into")
			}
			numbers = append(numbers, number)
		}
	}

	firstNumber := numbers[0]
	lastNumber := numbers[len(numbers)-1]

	return firstNumber*10 + lastNumber
}
