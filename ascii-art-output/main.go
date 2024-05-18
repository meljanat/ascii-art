package main

import (
	"fmt"
	"os"

	"asciiart-output/functions"
)

func main() {
	/******************************************/
	/***** Reading And Checking Arguments *****/
	/******************************************/
	args := os.Args
	if len(args) != 2 && len(args) != 3 && len(args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	Banner, FileName := (functions.Treatargs(args))

	/*********************************/
	/******** Verifying Input ********/
	/*********************************/
	verifiedInput := functions.VerifiedInput(args[2])

	/************************************/
	/******* Reading File Content *******/
	/************************************/
	FileContent := functions.ReadFile(Banner)

	/**********************/
	/******* Output *******/
	/**********************/

	functions.WriteContentintoFile(FileName, functions.Output(verifiedInput, FileContent))
}
