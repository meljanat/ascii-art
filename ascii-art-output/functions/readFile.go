package functions

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(file string) map[rune][]string {
	/**********************************/
	/******* Read File Content ********/
	/**********************************/
	open, err := os.Open(file)
	if err != nil {
		log.Fatalln("The Banner must be one of these: standard, shadow or thinkertoy!!")
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
