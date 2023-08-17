package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFileToArray(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := make([]string, 10)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}
