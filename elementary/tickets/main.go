package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getParams() (string, string) {
	var algo, filePath string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter an algorithm type <piter/moscow> and file-path.")

	if input, err := reader.ReadString('\n'); err == nil {
		if inputs := strings.Split(input, " "); len(inputs) > 1 {
			algo = inputs[0]
			filePath = inputs[1]
		}
	}

	return algo, filePath
}

func validInput(str1, str2 string) (string, string, error) {
	var algo, filePath string
	var err = errors.New("Not valid first input.")

	if strings.ToLower(str1) == "moscow" || strings.ToLower(str1) == "piter" {
		algo = strings.ToLower(str1)
		err = nil
	}

	file, err := os.Open(str2)
	defer file.Close()
	if err == nil {
		filePath = str2
	} else {
		err = errors.New("Not valid second input.")
	}

	return algo, filePath, err
}

func getTickets(filePath string) []int {
	var ticketNums = make([]int, 0)

	file, _ := os.Open(filePath)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ticketStr := strings.TrimSpace(scanner.Text())
		if ticketNum, err := strconv.Atoi(ticketStr); err == nil && len(ticketStr) == 6 {
			ticketNums = append(ticketNums, ticketNum)
		}
	}
	return ticketNums
}

func mosсowAlgo(tickets []int) int {
	var count int
	for _, tick := range tickets {
		var part1, part2 int
		part1 = tick/100000 + tick%100000/10000 + tick%10000/1000
		part2 = tick%1000/100 + tick%100/10 + tick%10
		if part1 == part2 {
			count++
		}
	}
	return count
}

func piterAlgo(tickets []int) int {
	var count int
	for _, tick := range tickets {
		var part1, part2 int
		part1 = tick/100000 + tick%100/10 + tick%10000/1000
		part2 = tick%1000/100 + tick%100000/10000 + tick%10
		if part1 == part2 {
			count++
		}
	}
	return count
}

func main() {
	var algo, file string
	var err = errors.New("No valid comman-line arguments.\n")

	if len(os.Args) > 2 {
		algo, file, err = validInput(os.Args[1], os.Args[2])
	}

	for err != nil {
		fmt.Println(err)
		inp1, inp2 := getParams()
		algo, file, err = validInput(inp1, inp2)
	}

	tickets := getTickets(file)

	if strings.ToLower(algo) == "moscow" {
		fmt.Printf("There are %d lucky tickets in the file counting moscow way.\n", mosсowAlgo(tickets))
	} else if strings.ToLower(algo) == "piter" {
		fmt.Printf("There are %d lucky tickets in the file counting piter way.\n", piterAlgo(tickets))
	}
}
