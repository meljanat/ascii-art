package functions

import (
	"fmt"
	"os"
	"strings"
)

func CheckFlag(args []string) string {
	flagName := ""
	if len(args) > 1 {
		if strings.HasPrefix(args[1], "--color=") {
			flagName = "--color="
		} else if strings.HasPrefix(args[1], "--output=") {
			flagName = "--output="
		} else if strings.HasPrefix(args[1], "--align=") {
			flagName = "--align="
		}
		if flagName != "" && flagName == args[1] {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")
			os.Exit(0)
		}
	}
	return flagName
}
