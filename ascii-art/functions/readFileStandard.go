package functions

import (
	"bufio"
	"log"
	"os"
)

func ReadFileStandard(file string) map[rune][]string {
	/***********************************/
	/******* Read File Standard ********/
	/***********************************/
	//fileContent, err := os.ReadFile(file)
	open, err := os.Open(file)
	if err != nil {
		log.Fatalln("Error in File Standard")
	}
	defer open.Close()

	scanner := bufio.NewScanner(open)

	contentFileStandard := make(map[rune][]string)
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
		contentFileStandard[rune(i+32)] = append(contentFileStandard[rune(i+32)], scanner.Text())
		count++
		if i == 95 {
			break
		}
	}

	return contentFileStandard
}
