package functions

import "fmt"

func FunctionTheColor(args []string, FileContent map[rune][]string) {
	var verifiedInput []string

	flagName := CheckFlag(args[1])
	flag := args[1]

	mainColor, resetColor := Colors(string(flag[len(flagName):]))
	/*********************************/
	/******** Verifying Input ********/
	/*********************************/

	var letters string

	if len(args) == 4 {
		verifiedInput = VerifiedInput(args[3])
		letters = args[2]

	} else if len(args) == 3 {
		verifiedInput = VerifiedInput(args[2])
		letters = ""
	}
	output := Output(verifiedInput, FileContent, letters, mainColor, resetColor)
	fmt.Print(output)
}
