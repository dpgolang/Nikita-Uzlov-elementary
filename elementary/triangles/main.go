package main

import (
	"./triangle"
	"./validate"
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

func want2add() bool {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Do you want to add another triangle? - y/yes")
	scanner.Scan()
	str := strings.ToLower(scanner.Text())
	return str == "y" || str == "yes"
}

func getParams() (name string, sides [3]float64) {
	var errM = errors.New("No input yet.")

	for ; errM != nil; fmt.Println(errM) {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Input Format: <triangle_name>,<side1>,<side2>,<side3> - sides are float numbers.")

		input, err := reader.ReadString('\n')
		if err != nil {
			errM = errors.New("No Input")
			continue
		}

		inputs := strings.Split(input, ",")
		if len(inputs) != 4 {
			errM = errors.New("Not appropriate amount of parameters.")
			continue
		}

		for i, inp := range inputs {
			inputs[i] = strings.TrimSpace(inp)
		}

		name = inputs[0]
		sides, err := validate.ValidateFloats(inputs[1:])
		if err != nil {
			errM = errors.New("Not all sizes are float numbers.")
			continue
		}
		return name, sides
	}
	return
}

func main() {
	triSlice := make([]triangle.Triangle, 0)

	fmt.Println("App takes triangles with sides-siezez as inputs and returns them in range of area descending.\n")

	for ok := true; ok; ok = want2add() {
		for {
			triName, triSides := getParams()
			if validate.ValidateSides(triSides) == true {
				triSlice = append(triSlice, triangle.New(triName, triSides))
				break
			}
		}
	}

	sort.Slice(triSlice, func(i, j int) bool {
		return triSlice[i].Area > triSlice[j].Area
	})

	fmt.Println("============= Triangles list: ===============")
	for i, tri := range triSlice {
		fmt.Printf("%d.[%s]: %0.2f cm\n", i+1, tri.Name, tri.Area)
	}
}
