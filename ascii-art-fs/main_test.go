package main

import (
	"os/exec"
	"strings"
	"testing"
)

func RunCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func Test(t *testing.T) {
	command := "go run . \"DEFAULT\" standard"
	output, err := RunCommand(command)
	if err != nil {
		t.Errorf("Error running command: %v", err)
	}
	arr := []string{
		" _____    ______   ______              _    _   _        _______  ",
		"|  __ \\  |  ____| |  ____|     /\\     | |  | | | |      |__   __| ",
		"| |  | | | |__    | |__       /  \\    | |  | | | |         | |    ",
		"| |  | | |  __|   |  __|     / /\\ \\   | |  | | | |         | |    ",
		"| |__| | | |____  | |       / ____ \\  | |__| | | |____     | |    ",
		"|_____/  |______| |_|      /_/    \\_\\  \\____/  |______|    |_|    ",
		"                                                                  ",
		"                                                                  ",
		"",
	}
	expected := strings.Join(arr, "\n")

	if output != expected {
		t.Errorf("Error: returned and expected output did not match")
	}
}
