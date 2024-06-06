package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func TerminalSize() int {
	// cmd := exec.Command("zsh", "-c", "tput cols")
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	Size, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	TSize, _ := strconv.Atoi(strings.Trim(string(Size), "\n"))
	return TSize
}
