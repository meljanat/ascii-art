package functions

import (
	"log"
)

func VerifiedInput(arguments []string) []string {
	/********************************/
	/******* Verifying Input ********/
	/********************************/
	input := arguments[1]
	
	for i := 0; i < len(input); i++ {
		if i+1 < len(input) && input[i] < 32 || input[i] > 126 {
			log.Fatalln("Please Enter Valid Characters")
		}
	}

	slInput := []string{}
	start := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\\' && input[i+1] == 'n' {
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
