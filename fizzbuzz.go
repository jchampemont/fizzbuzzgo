package main

import (
	"strconv"
)

// FizzBuzz count from 1 to limit, replacing multiples of int1 by string1, int2 by string2, or both by string1string2.
func FizzBuzz(string1 string, string2 string, int1 int, int2 int, limit int) []string {
	var result []string
	for i := 1; i <= limit; i++ {
		var s string

		if i%int1 == 0 {
			s = string1
		}

		if i%int2 == 0 {
			s += string2
		}

		if i%int1 != 0 && i%int2 != 0 {
			s = strconv.Itoa(i)
		}

		result = append(result, s)
	}
	return result
}
