package functions

import (
	"log"
	"os"
	"strings"
)

func Output(verifiedInput []string, FileContent map[rune][]string) string {
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
				for _, char := range word {
					line = FileContent[char][i]
					Content += line
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

func WriteContentintoFile(FileName, Content string) {
	/**************************************************/
	/**** Creating File and Writing Output into it ****/
	/**************************************************/
	FileName = strings.ReplaceAll(FileName, "../", "")
	FileName = strings.ReplaceAll(FileName, "./", "")
	Splitted := strings.Split(FileName, "/")
	Dir := "./Output/"
	for i, word := range Splitted {
		if word == "" {
			continue
		}
		if len(Splitted) == 1 {
			FileName = word
		} else if i < len(Splitted)-1 {
			Dir += word + "/"
		} else if i == len(Splitted)-1 {
			FileName = word
		}
	}

	err := os.MkdirAll(Dir, 0o755)
	if err != nil {
		log.Fatalln("Error creating directory", err)
	}

	file, err := os.Create(Dir + FileName)
	if err != nil {
		log.Fatalln("Error creating file", err)
	}
	defer file.Close()

	_, err = file.WriteString(Content)
	if err != nil {
		log.Fatalln("Error writing to file", err)
	}
}
