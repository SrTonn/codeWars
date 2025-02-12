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
		{"AAAA", "TTTT"},
		{"ATTGC", "TAACG"},
		{"GTAT", "CATA"},
	}

	for _, scenario := range scenarios {
		result := DNAStrand(scenario.input)
		if result != scenario.expectedOutput {
			t.Errorf(
				"Param(s)(%s) failed!\nExpected: %s\nReceived: %s",
				scenario.input, scenario.expectedOutput, result,
			)
		}
	}
}