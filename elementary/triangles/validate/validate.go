package validate

import (
	"strconv"
)

func ValidateFloats(inputs []string) (args [3]float64, errM error) {

	for i, input := range inputs {
		var arg, err = strconv.ParseFloat(input, 64)
		if err == nil {
			args[i] = arg
		} else {
			errM = err
		}
	}
	return
}

func ValidateSides(s [3]float64) bool {
	return (s[0] < s[1]+s[2]) && (s[1] < s[0]+s[2]) && (s[2] < s[1]+s[0])
}
