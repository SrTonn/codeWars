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
		{10, 3,[]int{8,118,334}},
		{27, 3,[]int{1, 999, 999}},
		{84, 4,[]int{}},
		{35, 6,[]int{123, 116999, 566666}},
	}

	for _, scenario := range scenarios {
		result := FindAll(scenario.input1, scenario.input2)
		if !reflect.DeepEqual(result, scenario.expectedOutput) {
		// if result != scenario.expectedOutput {
			t.Errorf(
				"Param(s)(%d, %d) failed!\nExpected: %d\nReceived: %d",
				scenario.input1, scenario.input2, scenario.expectedOutput, result,
			)
		}
	}
}