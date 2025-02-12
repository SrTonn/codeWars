// https://www.codewars.com/kata/5513795bd3fafb56c200049e/train/go

package main


func CountBy(x, n int) []int {
	var list []int = make([]int, n)
	
	for i := range list {
		list[i] = (i+1)*x
	}
	
	return list
}

func main() {
	result := CountBy(50, 5)
	for _, el := range result {
		println(el)
	}
}


	// dotest(1, 5, []int{1, 2, 3, 4, 5})
	// dotest(2, 5, []int{2, 4, 6, 8, 10})
	// dotest(3, 5, []int{3, 6, 9, 12, 15})
	// dotest(50, 5, []int{50, 100, 150, 200, 250})
	// dotest(100, 5, []int{100, 200, 300, 400, 500})