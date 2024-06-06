package functions

import (
	"fmt"
)

func Output(verifiedInput []string, contentFileStandard map[rune][]string) {
	/**********************************/
	/**** Print Output in Terminal ****/
	/**********************************/
	newLine := false
	for _, word := range verifiedInput {
		line := ""
		if word == "\\n" {
			fmt.Println()
		} else {
			newLine = true
			i := 0
			for i < 8 {
				for _, char := range word {
					line = contentFileStandard[char][i]
					fmt.Print(line)
				}
				if i < 7 {
					fmt.Println()
				}
				i++
			}
		}

	}
	if newLine {
		fmt.Println()
	}
}
