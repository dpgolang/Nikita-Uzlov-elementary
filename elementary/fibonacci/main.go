package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validateArgs(args []string) (uint, uint, bool) {

	if len(args) == 2 {
		if n, err := strconv.Atoi(args[0]); err == nil {
			if m, err := strconv.Atoi(args[1]); err == nil {
				if n > 0 && m > 0 && m > n {
					return uint(n), uint(m), true
				}
			}
		}
	}
	return 0, 0, false
}

func getInput() (uint, uint, bool) {
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("App returns slice of fibonacci sequence inbetween two boundary-numbers. ")
		fmt.Println("Enter two not equal integer numbers: smaller first, then - larger one.")

		if str, err := reader.ReadString('\n'); err == nil {
			args := strings.Fields(str)
			if n, m, isIt := validateArgs(args); isIt == true {
				return n, m, true
			}
		}
	}
}

func fib(num int) uint {
	if num <= 1 {
		return uint(num)
	}

	return fib(num-2) + fib(num-1)
}

func fiBorders(x uint) (uint, uint) {
	i := 0
	j := 0
	for fib(i) < x {
		i++
	}
	for fib(j) <= x {
		j++
	}

	return fib(i), fib(j - 1)
}

func main() {
	var p, q uint
	var fibSl []string

	if n, m, ok := validateArgs(os.Args[1:]); ok == true {
		p, _ = fiBorders(n)
		_, q = fiBorders(m)
	} else if n, m, ok := getInput(); ok == true {
		p, _ = fiBorders(n)
		_, q = fiBorders(m)
	}

	for i := 0; fib(i) <= q; i++ {
		if fib(i) >= p {
			fibSl = append(fibSl, strconv.Itoa(int(fib(i))))
		}
	}
	if len(fibSl) > 0 {
		fmt.Println(strings.Join(fibSl, ", "))
	} else {
		fmt.Println("No fibonacci numbers inbetween given boundaries.")
	}
}
