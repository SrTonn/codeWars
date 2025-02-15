// https://www.codewars.com/kata/5877e7d568909e5ff90017e6/train/go
package main

import (
	"fmt"
	"strings"
)

func DirReduc(arr []string) []string {
	opposites := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}

	var stack []string
	for _, dir := range arr {
		dir = strings.ToUpper(dir)

		if len(stack) > 0 && stack[len(stack)-1] == opposites[dir] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, dir)
		}
	}

	return stack
}


func main() {
	result := DirReduc([]string{"EAST","NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST", "WEST", "NORTH", "EAST"})
	fmt.Printf("result --> %v\n", result)
}


// a = []string{"NORTH", "SOUTH", "EAST", "WEST"}dotest(a, []string{})
// a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}dotest(a, []string{"WEST"})
// a = []string{"NORTH", "WEST", "SOUTH", "EAST"}dotest(a, []string{"NORTH", "WEST", "SOUTH", "EAST"})
// a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}dotest(a, []string{"NORTH"})