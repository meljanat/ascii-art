package functions

import (
	"strings"
)

func CheckBanner(args []string) string {
	Banner := "./sources/standard.txt"

	if (len(args) == 3 && !strings.HasPrefix(args[1], "--output=") && !strings.HasPrefix(args[1], "--color=") && !strings.HasPrefix(args[1], "--align=")) ||
		(len(args) == 4 && !strings.HasPrefix(args[1], "--color=")) {
		if strings.HasSuffix(args[len(args)-1], ".txt") {
			Banner = "./sources/" + args[len(args)-1]
		} else {
			Banner = "./sources/" + args[len(args)-1] + ".txt"
		}
	}
	return Banner
}
