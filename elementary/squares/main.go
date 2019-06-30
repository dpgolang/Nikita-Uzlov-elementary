package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() (string, error) {
	var arg string
	var err error

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter one positive integer number.")

	input, err := reader.ReadString('\n')
	if err != nil {
		err = errors.New("Can't read your input.")
	}

	if inputs := strings.Split(input, " "); len(inputs) == 1 {
		arg = strings.TrimSpace(inputs[0])
	} else {
		err = errors.New("Not appropriate quantity of inputs.")
	}

	return arg, err
}

func validInt(str string) (num int, err error) {
	num, err = strconv.Atoi(str)
	if err != nil || num < 0 {
		err = errors.New("Argument is not a positive integer.")
	}
	return
}

func printBasesOfLessSquares(n int) {
	for i := 0; i*i <= n; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func main() {
	var num int
	var err = errors.New("No command-line argument.\n")

	if len(os.Args) > 1 {
		num, err = validInt(os.Args[1])
	}

	for err != nil {
		fmt.Println(err)
		var arg string
		if arg, err = getInput(); err == nil {
			num, err = validInt(arg)
		}

	}

	printBasesOfLessSquares(num)
}
