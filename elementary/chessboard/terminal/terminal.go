package terminal

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func size() (string, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	return string(out), err
}

func parse(input string) (uint, uint, error) {
	parts := strings.Split(input, " ")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.Atoi(strings.Replace(parts[1], "\n", "", 1))
	if err != nil {
		return 0, 0, err
	}
	return uint(x), uint(y), nil
}

// Width return the width of the terminal.
func Width() (uint, error) {
	output, err := size()
	if err != nil {
		return 0, err
	}
	_, width, err := parse(output)
	return width, err
}

// Height returns the height of the terminal.
func Height() (uint, error) {
	output, err := size()
	if err != nil {
		return 0, err
	}
	height, _, err := parse(output)
	return height, err
}
