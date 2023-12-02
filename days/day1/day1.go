package day1

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

func Day1(filePath string) int {
	scanner, file := createScanner(filePath)
	defer file.Close()

	total := fileReader(scanner)

	return total

}

func createScanner(filePath string) (*bufio.Scanner, *os.File) {
	content, error := os.Open(filePath)

	if error != nil {
		log.Fatal(error)
	}

	scanner := bufio.NewScanner(content)

	return scanner, content
}

func fileReader(scanner *bufio.Scanner) int {
	var total int

	for scanner.Scan() {
		line := scanner.Text()
		totalLine := lineReader(line)
		total += totalLine
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}

func lineReader(line string) int {
	var firstNumber, lastNumber int

	for _, word := range line {
		if unicode.IsNumber(word) {
			number := int(word - '0')
			if firstNumber == 0 {
				firstNumber = number
			}
			lastNumber = number
		}
	}

	totalLine := firstNumber*10 + lastNumber

	return totalLine
}
