package functions

func Output(verifiedInput []string, FileContent map[rune][]string, letters string, mainColor string, resetColor string) string {
	/*************************************/
	/****** Making Output Printable ******/
	/*************************************/
	var Content string
	newLine := false
	for _, word := range verifiedInput {
		line := ""
		if word == "\\n" {
			Content += "\n"
		} else {
			newLine = true
			i := 0
			for i < 8 {
				for in, char := range word {
					if CheckIndex(in, word, letters) || len(letters) == 0 {
						line = mainColor + FileContent[char][i] + resetColor
						Content += line
					} else if letters == "" && mainColor == "" && resetColor == "" {
						line = FileContent[char][i]
						Content += line
					} else {
						line = FileContent[char][i]
						Content += line
					}
				}
				if i < 7 {
					Content += "\n"
				}
				i++
			}
		}
	}
	if newLine {
		Content += "\n"
	}
	return Content
}

func CheckIndex(index int, word string, letters string) bool {
	slIndexs := GetIndexsForColor(word, letters)
	for _, num := range slIndexs {
		if num == index {
			return true
		}
	}
	return false
}
