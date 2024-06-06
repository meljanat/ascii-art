package functions

import "strings"

func AsciiOutputLenth(line string, FileContent map[rune][]string) int {
	Content := ""
	Line := strings.Split(line, "")
	for _, word := range Line {
		line := ""
		i := 0
		for i < 8 {
			for _, char := range word {
				line = FileContent[char][i]
				Content += line
			}
			i++
		}
	}
	return len(Content) / 8
}
