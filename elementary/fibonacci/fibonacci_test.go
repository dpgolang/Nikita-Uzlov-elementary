package main

import (
	"strings"
	"testing"
)

func equals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestFibonacci(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{7, 13},
		{9, 34},
		{12, 144},
		{13, 233},
	}
	for _, test := range tests {
		got := fibonacci(test.input)
		if got != test.want {
			t.Errorf("fibonacci(%#v) = \"%v\", want \"%v\"", test.input, got, test.want)
		}
	}
}

func TestFibBorders(t *testing.T) {
	type wanted struct {
		want1 int
		want2 int
	}
	var tests = []struct {
		input int
		wanted
	}{
		{10, wanted{8, 13}},
		{43, wanted{34, 55}},
		{100, wanted{89, 144}},
		{2000, wanted{1597, 2584}},
	}
	for _, test := range tests {
		got1, got2 := fibBorders(test.input)
		if got1 != test.want1 || got2 != test.want2 {
			t.Errorf("fibBorders(%#v) = \"%v\", want \"%v\"", test.input, []int{got1, got2}, test.wanted)
		}
	}
}

func TestFibSlice(t *testing.T) {
	var tests = []struct {
		input [2]int
		want  []string
	}{
		{[2]int{5, 34}, []string{"5", "8", "13", "21", "34"}},
		{[2]int{13, 89}, []string{"13", "21", "34", "55", "89"}},
		{[2]int{144, 987}, []string{"144", "233", "377", "610", "987"}},
		{[2]int{8, 987}, []string{"8", "13", "21", "34", "55", "89", "144", "233", "377", "610", "987"}},
	}
	for _, test := range tests {
		got := fibSlice(test.input)
		if !equals(got, test.want) {
			t.Errorf("fibSlice(%#v) = \"%v\", want \"%v\"", test.input, got, test.want)
		}
	}
}

func TestValidArgs(t *testing.T) {
	var errTests = []struct {
		input []string
		want  string
	}{
		{[]string{"5", "8", "13", "21", "34"}, "Two inputs required, 5 entered."},
		{[]string{"34"}, "Two inputs required, 1 entered."},
		{[]string{"sfsf", "34"}, "First argument \"sfsf\" is not an integer number. "},
		{[]string{"3", "sdfdf"}, "Second argument \"sdfdf\" is not an integer number."},
		{[]string{"-2", "3"}, "\nArguments have to be positive.\n"},
		{[]string{"2", "-3"}, "\nArguments have to be positive.\n"},
		{[]string{"-2", "-3"}, "\nArguments have to be positive.\n"},
		{[]string{"2", "2"}, "\nNumbers should not be equal.\n"},
	}
	for _, errTest := range errTests {
		_, err := validArgs(errTest.input)
		if err == nil || !strings.Contains(err.Error(), errTest.want) {
			t.Errorf("validateInputs(%#v) error = \"%v\", want \"%v\"", errTest.input, err, errTest.want)
		}
	}
}
