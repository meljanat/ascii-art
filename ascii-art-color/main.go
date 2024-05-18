package main

import (
	"fmt"
	"os"

	"asciiart-justify/functions"
)

func main() {
	/******************************************/
	/***** Reading And Checking Arguments *****/
	/******************************************/
	args := os.Args
	if  len(args) != 2 && len(args) != 3 && len(args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}

	Banner, FileName, Option := (functions.Treatargs(args))

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
	if Option != "" {
		Output, TSize := functions.Output(verifiedInput, FileContent)
		functions.Justify(verifiedInput, FileContent, Option, Output, TSize)
	}
	if FileName != "" {
		Output, _ := functions.Output(verifiedInput, FileContent)
		functions.WriteContentintoFile(FileName, Output)
	}
}
