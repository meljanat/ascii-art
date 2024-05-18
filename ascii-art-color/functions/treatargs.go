package functions

import (
	"fmt"
	"os"
	"strings"
)

func Treatargs(args []string) (string, string, string) {
	// Getting Banner file from args
	Banner := "./sources/standard.txt"
	if len(args) == 4 {
		if len(args[3]) > 3 && !strings.HasSuffix(args[3], ".txt") {
			Banner = "./sources/" + args[3] + ".txt"
		} else {
			Banner = "./sources/" + args[3]
		}
	}

	// Getting output filename from args
	FileName := ""
	if strings.HasPrefix(args[1], "--output=") {
		if len(args[1]) > 9 {
			runes := []rune(args[1])
			FileName = string(runes[9:])
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
			os.Exit(0)
		}
	}

	// Getting desired Option from args
	Option := ""
	if strings.HasPrefix(args[1], "--align=") {
		if len(args[1]) > 8 {
			runes := []rune(args[1])
			Option = strings.ToLower(string(runes[8:]))
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
			os.Exit(0)
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
			content, _ := Output(VerifiedInput(args[1]), ReadFile(Banner))
			fmt.Print(content)
			os.Exit(0)
		}
	}
	if len(args) == 2 && args[1] != "" {
		content, _ := Output(VerifiedInput(args[1]), ReadFile(Banner))
		fmt.Print(content)
		os.Exit(0)

	} else if len(args) == 2 && args[1] == "" {
		os.Exit(0)
	}
	/**************************************************************************/

	return Banner, FileName, Option
}
