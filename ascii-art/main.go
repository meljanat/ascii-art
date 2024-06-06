package main

import (
	"log"
	"os"

	"asciiart/functions"
)

func main() {
	/************************************/
	/***** Read And Chekin Arguments*****/
	/************************************/
	arguments := os.Args
	if len(arguments) != 2 {
		log.Fatalln("Please Enter Arguments Correctly")
	}
	if arguments[1] == "" {
		return
	}

	/***********************************/
	/***** Verified Input the User *****/
	/***********************************/
	verifiedInput := functions.VerifiedInput(arguments)

	/**********************************************/
	/******* Read Content the File Standard *******/
	/**********************************************/
	contentFileStandard := functions.ReadFileStandard("./sources/standard.txt")

	/**********************/
	/******* Output *******/
	/**********************/
	functions.Output(verifiedInput, contentFileStandard)
}
