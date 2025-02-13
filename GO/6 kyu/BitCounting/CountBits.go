// https://www.codewars.com/kata/526571aae218b8ee490006f4/train/go

package main

import (
	"strconv"
	"strings"
)

func CountBits(n uint) int {
    bin := strconv.FormatInt(int64(n), 2)
    return strings.Count(bin, "1")
}

func main() {
	println(CountBits(4))
}


// Expect(CountBits(0)).To(Equal(0))
// Expect(CountBits(4)).To(Equal(1))
// Expect(CountBits(7)         ).To(Equal(3))
// Expect(CountBits(9)         ).To(Equal(2))
// Expect(CountBits(10)        ).To(Equal(2))