package main

import (
	"reflect"
	"testing"
)

type scenarios struct {
	input []int
	expectedOutput []int
}

func TestKatas(t *testing.T) {
	scenarios := []scenarios{
		{[]int{1, 2, 10, 50, 5}, []int{1, 2, 5, 10, 50}},
		{[]int{}, []int{}},
	}

	for _, scenario := range scenarios {
		result := SortNumbers(scenario.input)
		if !reflect.DeepEqual(result, scenario.expectedOutput) {
			t.Error("Error")
		}
	}
}