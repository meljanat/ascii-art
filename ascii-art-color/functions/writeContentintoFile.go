package functions

import (
	"log"
	"os"
	"strings"
)

func WriteContentIntoFile(FileName, Content string) {
	/**************************************************/
	/**** Creating File and Writing Output into it ****/
	/**************************************************/
	FileName = strings.ReplaceAll(FileName, "../", "")
	FileName = strings.ReplaceAll(FileName, "./", "")
	Splitted := strings.Split(FileName, "/")
	Dir := "./output/"
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
