package main

import (
	"testing"
)

type scenarios struct {
	input string
	expectedOutput string
}

func TestKatas(t *testing.T) {
	scenarios := []scenarios{
		{"Sam Harris","S.H"},
		{"Patrick Feenan","P.F"},
		{"Evan Cole","E.C"},
		{"P Favuzzi","P.F"},
		{"David Mendieta","D.M"},
	}

	for _, scenario := range scenarios {
		result := AbbrevName(scenario.input)
		if result != scenario.expectedOutput {
			t.Errorf(
				"Param(s)(%v) failed!\nExpected: %v\nReceived: %v",
				scenario.input, scenario.expectedOutput, result,
			)
		}
	}
}