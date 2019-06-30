package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Envelope [2]float64

func (e Envelope) fitIn(e2 Envelope) bool {
	p := e[0]
	q := e[1]
	a := e2[0]
	b := e2[1]

	if a >= p && b >= q {
		return true
	} else if a < p &&
		b >= (2*p*q*a+(p*p-q*q)*math.Sqrt(p*p+q*q-a*a))/(p*p+q*q) {
		return true
	} else {
		return false
	}
}

func getFloatSide(side, env int) float64 {

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("\nEnter size of side_%d for envelope_%d : ", side+1, env)
		scanner.Scan()
		str := scanner.Text()
		if arg, err := strconv.ParseFloat(str, 64); err == nil {
			return arg
		} else {
			fmt.Println("Not correct input. Required one float number.")
		}
	}
}

func doRepeat() bool {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nDo you want to continue? - y/yes")
	scanner.Scan()
	str := strings.ToLower(scanner.Text())
	return str == "y" || str == "yes"
}

func main() {

	fmt.Println("App discovers does one of two envelopes of given sizes fit into another or reverse.")

	for ok := true; ok; ok = doRepeat() {

		var env1 Envelope
		var env2 Envelope

		for i := range env1 {
			env1[i] = getFloatSide(i, 1)
		}
		for i := range env2 {
			env2[i] = getFloatSide(i, 2)
		}

		sort.Sort(sort.Reverse(sort.Float64Slice(env1[:])))
		sort.Sort(sort.Reverse(sort.Float64Slice(env2[:])))

		if env1.fitIn(env2) {
			fmt.Printf("\nEnvelope_1 fits into envelope_2\n")
		} else if env2.fitIn(env1) {
			fmt.Printf("\nEnvelope_2 fits into envelope_1\n")
		} else {
			fmt.Printf("\nNone of envelopes fit into another\n")
		}

	}
}
