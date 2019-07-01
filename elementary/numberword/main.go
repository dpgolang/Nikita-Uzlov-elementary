package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var digits = map[int]string{
	1: "один", 2: "два", 3: "три", 4: "четыре", 5: "пять", 6: "шесть", 7: "семь", 8: "восемь", 9: "девять",
}
var teens = map[int]string{
	10: "десять", 11: "одинадцать", 12: "двенадцать", 13: "тринадцать", 14: "четырнадцать",
	15: "пятнадцать", 16: "шестнадцать", 17: "семнадцать", 18: "восемнадцать", 19: "девятнадцать",
}
var tens = map[int]string{
	2: "двадцать", 3: "тридцать", 4: "сорок", 5: "пятьдесят", 6: "шестьдесят",
	7: "семьдесят", 8: "восемьдесят", 9: "девяносто",
}
var hundreds = map[int]string{
	1: "сто", 2: "двести", 3: "триста", 4: "четыреста", 5: "пятсот",
	6: "шестьсот", 7: "семьсот", 8: "восемьсот", 9: "девятьсот",
}
var thousands = map[int]string{
	0: "тысяч", 1: "одна тысяча", 2: "две тысячи", 3: "три тысячи", 4: "четыре тысячи",
}
var millions = map[int]string{
	0: "миллионов", 1: "миллион", 2: "миллиона", 3: "миллиона", 4: "миллиона",
	5: "миллионов", 6: "миллионов", 7: "миллионов", 8: "миллионов", 9: "миллионов",
}

func uniqPart(num int) string {
	var str string

	if num < 10 {
		str = digits[num]
	} else if num < 20 {
		str = teens[num]
	}
	return str
}

func under1K(num int) []string {
	var str = make([]string, 0)

	if num >= 100 && num < 1000 {
		str = append(str, hundreds[num/100])
	}

	if num%100 < 20 {
		str = append(str, uniqPart(num%100))
	} else {
		str = append(str, tens[num%100/10], digits[num%10])
	}

	return str
}

func over1K(num int) []string {
	var str = make([]string, 0)

	base := under1K(num / 1000)
	str = append(str, base[:len(base)-1]...)

	if num/1000%10 < 5 {
		str = append(str, thousands[num/1000%10])
	} else {
		str = append(str, base[len(base)-1], thousands[0])
	}

	str = append(str, under1K(num%1000)...)

	return str
}

func over1M(num int) []string {
	var str = make([]string, 0)

	base := under1K(num / 1000000)
	str = append(str, base...)

	if num/1000000%100/10 == 1 {
		str = append(str, millions[0])
	} else {
		str = append(str, millions[num/1000000%10])
	}

	str = append(str, over1K(num%1000000)...)

	return str
}

func getNumber() (int, error) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter an integer number to convert it to word-written one.")
		scanner.Scan()
		str := scanner.Text()
		str = strings.Join(strings.Fields(str), "")
		num, err := strconv.Atoi(str);
		if err != nil {
			fmt.Println("Not correct input. Required one integer number.")
		}
		if num >= 1000000000 {
			fmt.Println("Absolute value of number have to be less then billion.")
		}
		return num, err
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
	for ok := true; ok; ok = doRepeat() {

		var words = make([]string, 0)

		num, err := getNumber()
		if err != nil {
			log.Fatal(err)
		}

		if num < 0 {
			num = -num
			words = append(words, "минус")
		}

		if num == 0 {
			words = append(words, "ноль")
		} else if num < 1000 {
			words = append(words, under1K(num)...)
		} else if num < 1000000 {
			words = append(words, over1K(num)...)
		} else if num < 1000000000 {
			words = append(words, over1M(num)...)
		}

		for _, word := range words {
			fmt.Print(word + " ")
		}
		fmt.Println()
	}
}
