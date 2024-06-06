package functions

import (
	"fmt"
	"os"
	"strings"
)

func FileOutput(args []string, flagName string) (string, []string) {
	fileName := ""

	flagOne := args[1]
	// Get the File Name
	if len(flagOne) > len(flagName) {
		fileName = flagOne[len(flagName):]
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	// Check if it ends with txt
	if !strings.HasSuffix(fileName, ".txt") {
		fileName = fileName + ".txt"
	}

	slInput := VerifiedInput(args[2])
	return fileName, slInput
}
