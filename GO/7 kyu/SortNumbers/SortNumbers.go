// https://www.codewars.com/kata/5174a4c0f2769dd8b1000003/train/go

package main

func SortNumbers(numbers []int) (result []int) {
	result = numbers

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i - 1] {
			numbers[i-1], numbers[i] = numbers[i], numbers[i-1]
			SortNumbers(numbers)
		}
	}

	return
}

func main() {
	result := SortNumbers([]int{1,5,7,3})
	println(result)
	for _, val := range result {
		println(val)
	}
}

/** Description **/

// Tests
// Expect(Expect(SortNumbers([]int{1, 2, 10, 50, 5})).To(Equal([]int{1,2,5,10,50})))
// Expect(Expect(SortNumbers([]int{})).To(Equal([]int{})))