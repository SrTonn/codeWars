// https://www.codewars.com/kata/544675c6f971f7399a000e79/train/go

package main

import "strconv"

func StringToNumber(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}

func main() {
	println(StringToNumber("123"))
}

// Expect(StringToNumber("1234")).To(Equal(1234))
// Expect(StringToNumber("605")).To(Equal(605))
// Expect(StringToNumber("1405")).To(Equal(1405))
// Expect(StringToNumber("-7")).To(Equal(-7))