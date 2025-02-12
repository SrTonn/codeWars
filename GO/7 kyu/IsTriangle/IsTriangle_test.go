package main

import (
	"testing"
)

type scenarios struct {
	input1 int
	input2 int
	input3 int
	expectedOutput bool
}

func TestKatas(t *testing.T) {
	scenarios := []scenarios{
		{5, 1, 2, false},
		{1, 2, 5, false},
		{2, 5, 1, false},
		{4, 2, 3, true},
		{5, 1, 5, true},
		{2, 2, 2, true},
		{-1, 2, 3, false},
	}

	for _, scenario := range scenarios {
		result := IsTriangle(scenario.input1, scenario.input2, scenario.input3)
		if result != scenario.expectedOutput {
			t.Errorf(
				"Param(s)(%d, %d, %d) failed!\nExpected: %t\nReceived: %t",
				scenario.input1, scenario.input2, scenario.input3, scenario.expectedOutput, result,
			)
		}
	}
}