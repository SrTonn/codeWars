package main

import (
	"reflect"
	"testing"
)

type scenarios struct {
	input []string
	expectedOutput []string
}

func TestKatas(t *testing.T) {
	scenarios := []scenarios{
		//           
		{[]string{"NORTH", "SOUTH", "EAST", "WEST"}, []string{}},
		{[]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}, []string{"WEST"}},
		{[]string{"NORTH", "WEST", "SOUTH", "EAST"}, []string{"NORTH", "WEST", "SOUTH", "EAST"}},
		{[]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}, []string{"NORTH"}},
	}

	for _, scenario := range scenarios {
		result := DirReduc(scenario.input)
		if !reflect.DeepEqual(result, scenario.expectedOutput) {
		// if result != scenario.expectedOutput {
			t.Errorf(
				"Param(s)(%v) failed!\nExpected: %v\nReceived: %v",
				scenario.input, scenario.expectedOutput, result,
			)
		}
	}
}