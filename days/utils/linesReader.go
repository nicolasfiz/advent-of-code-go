package utils

import (
	"bufio"
	"log"
	"os"
)

func LinesReader(filePath string) []string {
	scanner, file := createScanner(filePath)
	defer file.Close()

	lines := fileReader(scanner)

	return lines
}

func createScanner(filePath string) (*bufio.Scanner, *os.File) {
	content, error := os.Open(filePath)

	if error != nil {
		log.Fatal(error)
	}

	scanner := bufio.NewScanner(content)

	return scanner, content
}

func fileReader(scanner *bufio.Scanner) []string {
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
