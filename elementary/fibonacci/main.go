package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput(reader *bufio.Reader) (args []string, err error) {
	fmt.Println("App returns slice of fibonacci sequence between two boundary-numbers. \n" +
		"Enter two not equal integer numbers.")

	str, err := reader.ReadString('\n')
	if err == nil {
		args = strings.Fields(str)
		for i := range args {
			strings.TrimSpace(args[i])
		}
	}
	return
}

func validArgs(inputs []string) (args [2]int, err error) {

	if len(inputs) != 2 {
		err = fmt.Errorf("\nTwo inputs required, %d entered.\n", len(inputs))
		return [2]int{0, 0}, err
	}

	args[0], err = strconv.Atoi(inputs[0])
	if err != nil {
		err = fmt.Errorf("\nFirst argument \"%v\" is not an integer number. \n%s\n", inputs[0], err)
		return
	}

	args[1], err = strconv.Atoi(inputs[1])
	if err != nil {
		err = fmt.Errorf("\nSecond argument \"%v\" is not an integer number. \n%s\n", inputs[1], err)
		return
	}

	if args[0] < 0 || args[1] < 0 {
		err = fmt.Errorf("\nArguments have to be positive.\n")
		return
	}

	if args[0] == args[1] {
		err = fmt.Errorf("\nNumbers should not be equal.\n")
		return
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

func fibSlice(fibs [2]int) (fibSl []string) {
	sort.Ints(fibs[:])
	for i := 0; fibonacci(i) <= fibs[1]; i++ {
		if fibonacci(i) >= fibs[0] {
			fibSl = append(fibSl, strconv.Itoa(int(fibonacci(i))))
		}
	}
	return
}

func main() {
	var args [2]int
	var err = errors.New("No inupts given.")

	if len(os.Args) > 2 {
		args, err = validArgs(os.Args[1:])
	}

	for err != nil {
		fmt.Println(err)
		var inputs []string
		inputs, err = getInput(bufio.NewReader(os.Stdin))
		if err == nil {
			args, err = validArgs(inputs)
		}
	}

	fibSl := fibSlice(args)

	if len(fibSl) > 0 {
		fmt.Println(strings.Join(fibSl, ", "))
	} else {
		fmt.Println("No fibonacci numbers between given boundaries.")
	}

}
