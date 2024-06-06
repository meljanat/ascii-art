package functions

import (
	"fmt"
)

func Justify(verifiedInput []string, FileContent map[rune][]string, Option string, TSize int) {
	/*******************************************************/
	/********* Printing output with desired option *********/
	/*******************************************************/
	newLine := false
	for _, word := range verifiedInput {
		LineLenth := AsciiOutputLenth(word, FileContent)
		Spaces := CountSpaces(word)
		line := ""
		if word == "\\n" {
			fmt.Println()
		} else {
			newLine = true
			i := 0
			for i < 8 {
				if Option == "right" {
					for j := 1; j <= TSize-LineLenth; j++ {
						fmt.Print(" ")
					}
				} else if Option == "center" {
					for j := 1; j <= (TSize-LineLenth)/2; j++ {
						fmt.Print(" ")
					}
				}

				for _, char := range word {
					if Option == "justify" && char == ' ' {
						for j := 1; j <= ((TSize-LineLenth)/Spaces); j++ {
							fmt.Print(" ")
						}
					}
					line = FileContent[char][i]
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
