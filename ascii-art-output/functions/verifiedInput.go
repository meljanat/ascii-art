package functions

import (
	"log"
)

func VerifiedInput(input string) []string {
	/********************************/
	/******* Verifying Input ********/
	/********************************/

	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 {
			log.Fatalln("Invalid Character: '" + string(input[i]) + "', Please Enter Valid Characters!")
		}
	}

	slInput := []string{}
	start := 0
	for i := 0; i < len(input); i++ {
		if i+1 < len(input) && input[i] == '\\' && input[i+1] == 'n' {
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
