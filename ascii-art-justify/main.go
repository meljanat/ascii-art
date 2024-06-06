package main

import (
	"os"

	"asciiart-justify/functions"
)

func main() {
	/******************************************/
	/***** Reading And Checking Arguments *****/
	/******************************************/
	args := os.Args
	functions.VerifiedArguments(args)

	/*************************************/
	/********** Check Flag Name **********/
	/*************************************/
	flagName := functions.CheckFlag(args)

	/*************************************/
	/********** Checking Banner **********/
	/*************************************/
	Banner := functions.CheckBanner(args)

	/************************************/
	/******* Reading File Content *******/
	/************************************/
	FileContent := functions.ReadFile(Banner)

	/******************/
	/******* FS *******/
	/******************/
	functions.FunctionFs(args, FileContent)

	/******************************************/
	/***** Checking Flags : Color-Output ******/
	/******************************************/
	if flagName == "--color=" {
		functions.FunctionTheColor(args, FileContent)
	}

	if flagName == "--output=" {
		fileName, verifiedInput := functions.FileOutput(args, flagName)

		functions.WriteContentIntoFile(fileName, functions.Output(verifiedInput, FileContent, "", "", ""))
	}

	if flagName == "--align=" {
		functions.JustifyFunction(args, FileContent, flagName)
	}
}
