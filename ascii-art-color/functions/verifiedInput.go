package functions

import "fmt"

func VerifiedInput(input string) []string {
	/********************************/
	/******* Verifying Input ********/
	/********************************/

	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
			break
		}
	}

	slInput := []string{}
	start := 0
	for i := 0; i < len(input); i++ {
		if i < len(input)-1 && input[i] == '\\' && input[i+1] == 'n' {
			if input[start:i] != "" {
				slInput = append(slInput, input[start:i])
			}
			slInput = append(slInput, "\\n")
			start = i + 2
		}
	}
	if input[start:] != "" {
		slInput = append(slInput, input[start:])
	}
	return slInput
}
