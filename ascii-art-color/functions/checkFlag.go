package functions

import (
	"fmt"
	"os"
	"strings"
)

func CheckFlag(flag string) string {
	flagName := ""
	if strings.HasPrefix(flag, "--color=") {
		flagName = "--color="
	} else if strings.HasPrefix(flag, "--output=") {
		flagName = "--output="
	}
	if flagName == flag && flag != "" {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}
	return flagName
}
