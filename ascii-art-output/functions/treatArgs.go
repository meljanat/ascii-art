package functions

import (
	"fmt"
	"os"
	"strings"
)

func Treatargs(args []string) (string, string) {
	// Getting Banner file from args
	Banner := "./sources/standard.txt"
	if len(args) == 4 {
		runes := []rune(args[3])
		if len(args[3]) > 3 && string(runes[len(args[3])-4:]) != ".txt" {
			Banner = "./sources/" + args[3] + ".txt"
		} else {
			Banner = "./sources/" + args[3]
		}
	}

	/****************** Single string or string with Banner cases ******************/
	if len(args) == 3 && !strings.HasPrefix(args[1], "--output=") {
		if args[2] == "standard" || args[2] == "standard.txt" || args[2] == "shadow" || args[2] == "shadow.txt" || 
			args[2] == "thinkertoy" || args[2] == "thinkertoy.txt" || args[2] == "example.txt" || args[2] == "example" {
			if !strings.HasSuffix(args[2], ".txt") {
				Banner = "./sources/" + args[2] + ".txt"
			} else {
				Banner = "./sources/" + args[2]
			}
			content := Output(VerifiedInput(args[1]), ReadFile(Banner))
			fmt.Print(content)
			os.Exit(0)
		} else if args[2] != "" {
			fmt.Println("The Banner Must Contain One Of These: standard, shadow or thinkertoy!!")
			os.Exit(0)
		}
	}
	if len(args) == 2 && args[1] != "" {
		content := Output(VerifiedInput(args[1]), ReadFile(Banner))
		fmt.Print(content)
		os.Exit(0)

	} else if len(args) == 2 && args[1] == "" {
		os.Exit(0)
	}
	/**************************************************************************/

	// Getting output filename from args
	FileName := ""
	runes := []rune(args[1])
	if len(args[1]) > 9 && string(runes[:9]) == "--output=" {
		FileName = string(runes[9:])
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	return Banner, FileName
}
