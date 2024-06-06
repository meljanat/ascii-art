package functions

import (
	"fmt"
	"strings"
)

func FunctionFs(args []string, FileContent map[rune][]string) {
	if !strings.HasPrefix(args[1], "--color=") && !strings.HasPrefix(args[1], "--output=") {

		var verifiedInput []string

		if len(args) == 2 || len(args) == 3 {
			verifiedInput = VerifiedInput(args[1])
		}
		output := Output(verifiedInput, FileContent, "", "", "")
		fmt.Print(output)
	}
}
