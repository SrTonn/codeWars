package main

import (
	"reflect"
	"testing"
)

type scenarios struct {
	input1 int
	input2 int
	expectedOutput []int
}

func TestKatas(t *testing.T) {
	scenarios := []scenarios{
		{1, 5, []int{1, 2, 3, 4, 5}},
		{2, 5, []int{2, 4, 6, 8, 10}},
		{3, 5, []int{3, 6, 9, 12, 15}},
		{50, 5, []int{50, 100, 150, 200, 250}},
		{100, 5, []int{100, 200, 300, 400, 500}},
	}

	for _, scenario := range scenarios {
		result := CountBy(scenario.input1, scenario.input2)
		if !reflect.DeepEqual(result, scenario.expectedOutput) {
			t.Error("Error")
		}
	}
}