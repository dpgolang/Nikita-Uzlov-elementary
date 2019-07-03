package main

import (
	"fmt"
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

func TestFibSlice(t *testing.T) {
	var tests = []struct {
		input [2]int
		want  []string
	}{
		{[2]int{5, 34},[]string{"5", "8", "13", "21", "34"}},
		{[2]int{13, 89}, []string{"13", "21", "34", "55", "89"}},
		{[2]int{144, 987},[]string{"144", "233", "377", "610", "987"}},
	}
	for _, test := range tests {
		got := fibSlice(test.input)
		if !equals(got, test.want) {
			t.Errorf("fibSlice(%#v) = \"%v\", want \"%v\"", test.input, got, test.want)
		}
	}
}

func TestValidArgs(t *testing.T) {
	type want struct {
		res [2]int
		err error
	}
	var tests = []struct {
		input []string
		want
	}{{[]string{"7", "15"},
		want{[2]int{7, 15}, nil}},
		{[]string{"5", "8", "13", "21", "34"},
			want{[2]int{0, 0}, fmt.Errorf("\nTwo inputs required, %d entered.\n", 5)}},
		{[]string{"34"},
			want{[2]int{0, 0}, fmt.Errorf("\nTwo inputs required, %d entered.\n", 1)}},
		{[]string{"sfsf", "34"},
			want{[2]int{0, 0}, fmt.Errorf("\nFirst argument \"sfsf\" is not an integer number.\n")}},
		{[]string{"34", "sdfdf"},
			want{[2]int{0, 0}, fmt.Errorf("\nSecond argument \"sdfdf\" is not an integer number.\n")}},
		{[]string{"-2", "3"},
			want{[2]int{0, 0}, fmt.Errorf("\nArguments have to be positive.\n")}},
		{[]string{"2", "-3"},
			want{[2]int{0, 0}, fmt.Errorf("\nArguments have to be positive.\n")}},
		{[]string{"-2", "-3"},
			want{[2]int{0, 0}, fmt.Errorf("\nArguments have to be positive.\n")}},
		{[]string{"2", "2"},
			want{[2]int{0, 0}, fmt.Errorf("\nNumbers should not be equal.\n")}},
	}
	for _, test := range tests {
		val, err := validArgs(test.input)
		if err != nil && err.Error() != test.err.Error() || val != test.res {
			t.Errorf("validArgs(%#v) = \"%v\" \"%v\", want \"%d\" \"%s\"\n", test.input, val, err.Error(), test.res, test.err.Error())
		} else if err == nil && val != test.res {
			t.Errorf("validArgs(%#v) value =  %d, want %d", test.input, val, test.res)
		}
	}
}
