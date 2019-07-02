package main

import (
	"strconv"
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

func TestFibSlice(t *testing.T) {
	var tests = []struct {
		inp1 int
		inp2 int
		want []string
	}{
		{5, 34, []string{"5", "8", "13", "21", "34"}},
		{13, 89, []string{"13", "21", "34", "55", "89"}},
		{144, 987, []string{"144", "233", "377", "610", "987"}},
		{8, 987, []string{"8", "13", "21", "34", "55", "89", "144", "233", "377", "610", "987"}},
	}
	for _, test := range tests {
		got := fibSlice(test.inp1, test.inp2)
		if !equals(got, test.want) {
			t.Errorf("fibSlice(%#v) = \"%v\", want \"%v\"", []int{test.inp1, test.inp2}, got, test.want)
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

func TestValidateArgs(t *testing.T) {
	type wanted struct {
		first  int
		second int
		err    error
	}
	var tests = []struct {
		input []string
		want  wanted
	}{
		{[]string{"24", "345"}, wanted{24, 345, nil}},
		{[]string{"12", "818"}, wanted{12, 818, nil}},
	}

	for _, test := range tests {
		got1, got2, gotErr := validateArgs(test.input)
		var got = wanted{got1, got2, gotErr}
		if got != test.want {
			t.Errorf("validateArgs(%#v) = \"%v\", want \"%v\"", test.input, got, test.want)
		}
	}
}

func TestValidateArgs2(t *testing.T) {

	_, _, got := validateArgs([]string{"34", "dsf"})
	_, err := strconv.Atoi("dsf")
	if !strings.Contains(got.Error(), err.Error()) {
		t.Errorf("validateArgs(%#v) = \"%v\", want \"%v\"",[]string{"34", "dsf"}, got.Error(), err.Error() )
	}
}