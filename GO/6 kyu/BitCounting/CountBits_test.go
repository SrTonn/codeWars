package main

import (
	"testing"
)

type scenarios struct {
	input uint
	expectedOutput int
}

func TestKatas(t *testing.T) {
	scenarios := []scenarios{
		{0, 0},
		{4, 1},
		{7, 3},
		{9, 2},
		{10, 2},
	}

	for _, scenario := range scenarios {
		result := CountBits(scenario.input)
		if result != scenario.expectedOutput {
			t.Errorf(
				"Param(s)(%d) failed!\nExpected: %d\nReceived: %d",
				scenario.input, scenario.expectedOutput, result,
			)
		}
	}
}