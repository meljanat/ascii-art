package functions

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(file string) map[rune][]string {
	/**********************************/
	/******* Read File Content ********/
	/**********************************/
	open, err := os.Open(file)
	if err != nil {
		fmt.Println("Invalid banner file! banner must be standard, shadow, thinkertoy or example.")
		os.Exit(0)
	}
	defer open.Close()

	scanner := bufio.NewScanner(open)

	FileContent := make(map[rune][]string)
	i := 0
	count := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		if count == 8 {
			i++
			count = 0
		}
		FileContent[rune(i+32)] = append(FileContent[rune(i+32)], scanner.Text())
		count++
		if i == 95 {
			break
		}
	}
	return FileContent
}
