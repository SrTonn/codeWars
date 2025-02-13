// https://www.codewars.com/kata/5877e7d568909e5ff90017e6/train/go
// kata didn't accept 'cause the code has slow performance
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isValidSequence(num int) bool {
	digits := strconv.Itoa(num)
	for i := 1; i < len(digits); i++ {
		if digits[i-1] > digits[i] {
			return false
		}
	}
	return true
}

func getStrNumSum(num int) (int) {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func getHighestValue(digits int) int {
	if digits <= 0 {
		return 0
	}
	return int(math.Pow10(digits)) - 1
}

func getMinimumValue(digits int) int {
	if digits <= 0 {
		return 0
	}
	ones := strings.Repeat("1", digits)
	minValue, _ := strconv.Atoi(ones)
	return minValue
}

func FindAll(sumDig, digs int) []int {
	var found []int
	minValue := getMinimumValue(digs)
	maxValue := getHighestValue(digs)

	for i := minValue; i <= maxValue; i++ {
		if !isValidSequence(i) { continue }
		sum := getStrNumSum(i)
		if sum == sumDig {
			found = append(found, i)
		}
	}

	if len(found) < 1 {
		return []int{}
	}

	return []int{len(found), found[0], found[len(found)-1]}
}

func main() {
	result := FindAll(10, 3)
	fmt.Printf("result --> %v\n", result)
}


// Expect(FindAll(10, 3)).To(Equal([]int{8,118,334}))
// Expect(FindAll(27, 3)).To(Equal([]int{1, 999, 999}))
// Expect(FindAll(84, 4)).To(BeEmpty())
// Expect(FindAll(35, 6)).To(Equal([]int{123, 116999, 566666}))