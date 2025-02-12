// https://www.codewars.com/kata/56606694ec01347ce800001b/train/go

package main

func IsTriangle(a, b, c int) bool {
	return a + b > c && a + c > b && b + c > a
}

func main() {
	println(IsTriangle(5, 1, 2))
}


// Expect(IsTriangle(5, 1, 2)).To(Equal(false))
// Expect(IsTriangle(1, 2, 5)).To(Equal(false))
// Expect(IsTriangle(2, 5, 1)).To(Equal(false))
// Expect(IsTriangle(4, 2, 3)).To(Equal(true))
// Expect(IsTriangle(5, 1, 5)).To(Equal(true))
// Expect(IsTriangle(2, 2, 2)).To(Equal(true))
// Expect(IsTriangle(-1, 2, 3)).To(Equal(false))