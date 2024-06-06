package functions

import (
	"fmt"
	"os"
	"strings"
)

func FunctionFs(args []string, FileContent map[rune][]string) {
	if !strings.HasPrefix(args[1], "--color=") && !strings.HasPrefix(args[1], "--output=") && !strings.HasPrefix(args[1], "--align=") {

		var verifiedInput []string

		if len(args) == 2 || len(args) == 3 {
			if args[1] == "" {
				os.Exit(0)
			}
			verifiedInput = VerifiedInput(args[1])
		}
		output := Output(verifiedInput, FileContent, "", "", "")
		fmt.Print(output)
	}
}
