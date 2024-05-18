package functions

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Output(verifiedInput []string, FileContent map[rune][]string) (string, int) {
	/*******************************************/
	/********* Making output printable *********/
	/*******************************************/

	// Getting Terminal size *******************/
	Size := RunCommand("stty size")
	SSize := strings.Split(Size, " ")
	TSize, _ := strconv.Atoi(strings.Trim(SSize[1], "\n"))
	/*******************************************/

	output := ""
	newLine := false
	for _, word := range verifiedInput {
		line := ""
		if word == "\\n" {
			output += "\n"
		} else {
			newLine = true
			i := 0
			for i < 8 {
				for _, char := range word {
					line = FileContent[char][i]
					output += line
				}
				if i < 7 {
					output += "\n"
				}
				i++
			}
		}

	}
	if newLine {
		output += "\n"
	}
	return output, TSize
}

func Justify(verifiedInput []string, FileContent map[rune][]string, Option, Output string, TSize int) {
	/*******************************************************/
	/********* Printing output with desired option *********/
	/*******************************************************/
	if Option != "right" && Option != "left" && Option != "center" && Option != "justify" {
		log.Fatalln("Invalid Option: " + Option + ", use: Right, Left, Center or Justify!")
	}
	
	SOutput := strings.Split(Output, "\n")
	LineLenth := len(SOutput[0])

	newLine := false
	for _, word := range verifiedInput {
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

func RunCommand(command string) string {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdin = os.Stdin
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(output)
}
