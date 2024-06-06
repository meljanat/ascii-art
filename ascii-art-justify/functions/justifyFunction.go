package functions

func JustifyFunction(args []string, FileContent map[rune][]string, flagName string) {
	option, verifiedInput := CheckAlignOption(args, flagName)

	terminalSize := TerminalSize()

	Justify(verifiedInput, FileContent, option, terminalSize)
}
