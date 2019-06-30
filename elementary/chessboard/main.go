package main

import (
	"./terminal"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() (string, string) {
	var arg1, arg2 string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter two positive integer numbers: height and width of chess-desk.")

	if input, err := reader.ReadString('\n'); err == nil {

		if inputs := strings.Split(input, " "); len(inputs) == 2 {
			arg1 = strings.TrimSpace(inputs[0])
			arg2 = strings.TrimSpace(inputs[1])
		}
	}

	return arg1, arg2
}

func validInputs(arg1, arg2 string) (height, width int, err error) {

	if height, err = strconv.Atoi(arg1); err != nil || height <= 0 {
		err = errors.New("first argument is not a positive integer")
		return
	}
	if width, err = strconv.Atoi(arg2); err != nil || width <= 0 {
		err = errors.New("second argument is not a positive integer")
		return
	}

	maxHeight, _ := terminal.Height()
	maxWidth, _ := terminal.Width()

	if uint(height) > maxHeight {
		err = errors.New(fmt.Sprintf("height have to be less then %d", maxHeight))
	}
	if uint(width) > maxWidth {
		err = errors.New(fmt.Sprintf("width have to be less then %d", maxWidth))
	}

	return height, width, err
}

func drawBoard(height, width int) {
	for i := 1; i <= height; i++ {

		var tempStr string

		if i%2 != 0 {
			tempStr = "* "
		} else {
			tempStr = " *"
		}
		fmt.Println(strings.Repeat(tempStr, width))
	}
}

func main() {
	var height, width int
	var err = errors.New("No valid command-line arguments.\n")

	if len(os.Args) > 2 {
		height, width, err = validInputs(os.Args[1], os.Args[2])
	}

	for err != nil {
		fmt.Println(err)
		inp1, inp2 := getInput()
		height, width, err = validInputs(inp1, inp2)
	}

	drawBoard(height, width)
}
