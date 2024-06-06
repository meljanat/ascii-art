package main

import (
	"log"
	"os"

	"asciiart-fs/functions"
)

func main() {
	/******************************************/
	/***** Reading And Checking Arguments *****/
	/******************************************/
	args := os.Args
	if len(args) != 2 && len(args) != 3 {
		log.Fatalln("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}
	if args[1] == "" {
		return
	}
	Banner := "./sources/standard.txt"
	if len(args) == 3 {
		runes := []rune(args[2])
		if len(args[2]) > 3 && string(runes[len(args[2])-4:]) != ".txt" {
			Banner = "./sources/" + args[2] + ".txt"
		} else {
			Banner = "./sources/" + args[2]
		}
	}

	/*********************************/
	/******** Verifying Input ********/
	/*********************************/
	verifiedInput := functions.VerifiedInput(args)

	/************************************/
	/******* Reading File Content *******/
	/************************************/
	FileContent := functions.ReadFile(Banner)

	/**********************/
	/******* Output *******/
	/**********************/
	functions.Output(verifiedInput, FileContent)
}
