package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tables := []struct {
		string1 string
		string2 string
		int1    int
		int2    int
		limit   int
		result  []string
	}{
		{"fizz", "buzz", 3, 5, 16,
			[]string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"}},
		{"deux", "trois", 2, 3, 6,
			[]string{"1", "deux", "trois", "deux", "5", "deuxtrois"}},
		{"", "", 1, 1, 10,
			[]string{"", "", "", "", "", "", "", "", "", ""}},
	}

	for _, table := range tables {
		fizzbuzz := FizzBuzz(table.string1, table.string2, table.int1, table.int2, table.limit)
		if !reflect.DeepEqual(table.result, fizzbuzz) {
			t.Errorf("FizzBuzz result %v is not matching expectations %v.", fizzbuzz, table.result)
		}
	}
}
