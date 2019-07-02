package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validateArgs(args []string) (bottom, top int, err error) {

	if len(args) != 2 {
		err = fmt.Errorf("\nTwo integer numbers required, %d entered.\n", len(args))
		return 0, 0, err
	}

	bottom, err = strconv.Atoi(args[0])
	if err != nil {
		err = fmt.Errorf("\nFirst argument \"%v\" is not an integer number. \n%s\n", args[0], err)
		return
	}

	top, err = strconv.Atoi(args[1])
	if err != nil {
		err = fmt.Errorf("\nSecond argument \"%v\" is not an integer number. \n%s\n", args[1], err)
		return
	}

	if bottom < 0 || top < 0 {
		err = fmt.Errorf("\nArguments have to be positive.\n")
	} else if top < bottom {
		err = fmt.Errorf("\nFirst argument %d have to be lower the second one %d.\n", bottom, top)
	}

	return
}

func getInput(reader *bufio.Reader) (args []string, err error) {
	fmt.Println("App returns slice of fibonacci sequence inbetween two boundary-numbers. \n" +
		"Enter two not equal integer numbers: smaller first, then - larger one.")

	str, err := reader.ReadString('\n')
	if err == nil {
		args = strings.Fields(str)
		for i := range args {
			strings.TrimSpace(args[i])
		}
	}
	return
}

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}

	return fibonacci(num-2) + fibonacci(num-1)
}

func fibBorders(x int) (int, int) {
	var i int

	for fibonacci(i) < x {
		i++
		if fibonacci(i) == x {
			return x, x
		}
	}

	return fibonacci(i - 1), fibonacci(i)
}

func fibSlice(bottom, top int) (fibSl []string) {
	for i := 0; fibonacci(i) <= top; i++ {
		if fibonacci(i) >= bottom {
			fibSl = append(fibSl, strconv.Itoa(int(fibonacci(i))))
		}
	}
	return
}

func main() {
	var bottom, top int
	var err = fmt.Errorf("No inputs received.")

	if len(os.Args) > 2 {
		bottom, top, err = validateArgs(os.Args[1:])
	}

	for err != nil {
		fmt.Println(err)
		var inputs []string
		inputs, err = getInput(bufio.NewReader(os.Stdin))
		if err == nil {
			bottom, top, err = validateArgs(inputs)
		}
	}

	_, p := fibBorders(bottom)
	q, _ := fibBorders(top)

	fibSl := fibSlice(p, q)

	if len(fibSl) > 0 {
		fmt.Println(strings.Join(fibSl, ", "))
	} else {
		fmt.Println("No fibonacci numbers between given boundaries.")
	}

}
