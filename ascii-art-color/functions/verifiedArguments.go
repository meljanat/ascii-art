package functions

import (
	"fmt"
	"os"
	"strings"
)

func VerifiedArguments(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}
	if strings.HasPrefix(args[1], "--output=") || strings.HasPrefix(args[1], "--color=") {
		if len(args) < 3 || len(args) > 4 {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(0)
		}
	} else {
		if len(args) != 2 && len(args) != 3 {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(0)
		}
	}
}
