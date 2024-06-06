package functions

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CheckAlignOption(args []string, flagName string) (string, []string) {
	Option := ""
	flagOne := args[1]

	// Get the option
	if len(flagOne) > len(flagName) {
		Option = strings.ToLower(flagOne[len(flagName):])
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")
		os.Exit(0)
	}
	if Option != "right" && Option != "left" && Option != "center" && Option != "justify" {
		log.Fatalln("Invalid Option: " + Option + ", use: Right, Left, Center or Justify!")
	}

	slInput := VerifiedInput(args[2])

	return Option, slInput
}
