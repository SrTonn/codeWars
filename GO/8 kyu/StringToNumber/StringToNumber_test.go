package main

import "testing"

type scenarios struct {
	input string
	expectedOutput int
}

func TestKatas(t *testing.T) {
	scenarios := []scenarios{
		{"1234", 1234},
		{"605", 605},
		{"1405", 1405},
		{"-7", -7},
	}

	for _, scenario := range scenarios {
		result := StringToNumber(scenario.input)
		if result != scenario.expectedOutput {
			t.Errorf(
				"Param(s)(%s) failed!\nExpected: %d\nReceived: %d",
				scenario.input, scenario.expectedOutput, result,
			)
		}
	}
}