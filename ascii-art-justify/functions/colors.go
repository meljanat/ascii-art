package functions

import (
	"fmt"
	"os"
	"strings"
)

func Colors(color string) (string, string) {
	// ANSI color codes
	const (
		Black      = "\x1b[0;30m"
		Red        = "\x1b[1;31m"
		Green      = "\x1b[0;32m"
		Blue       = "\x1b[0;34m"
		Yellow     = "\x1b[1;33m"
		White      = "\x1b[1;37m"
		Orange     = "\033[38;5;208m"
		Brown      = "\033[38;5;130m"
		Purple     = "\x1b[0;35m"
		Gray       = "\x1b[1;30m"
		Lime       = "\x1b[1;32m"
		DodgerBlue = "\x1b[1;34m"
		Magenta    = "\x1b[1;35m"
		Cyan       = "\x1b[1;36m"
		Silver     = "\x1b[0;37m"
		colorReset = "\033[0m"
	)

	switch strings.ToLower(color) {
	case "gray":
		color = Gray
	case "red":
		color = Red
	case "lime":
		color = Lime
	case "yellow":
		color = Yellow
	case "dodgerblue":
		color = DodgerBlue
	case "magenta":
		color = Magenta
	case "cyan":
		color = Cyan
	case "white":
		color = White
	case "black":
		color = Black
	case "brown":
		color = Brown
	case "green":
		color = Green
	case "orange":
		color = Orange
	case "blue":
		color = Blue
	case "purple":
		color = Purple
	case "silver":
		color = Silver

	default:
		fmt.Println("Invalid color: `", color, "`!!")
		os.Exit(0)
	}

	return color, colorReset
}
