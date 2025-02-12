// https://www.codewars.com/kata/56747fd5cb988479af000028/train/go

package main

func GetMiddle(s string) string {
	if len(s)%2 == 0 {
		return s[(len(s)/2)-1 : (len(s)/2)+1]
	}

	return s[(len(s) / 2) : (len(s)/2)+1]
}

func main() {
	println(GetMiddle("testing"))
}

// Expect(GetMiddle("test")).To(Equal("es"))
// Expect(GetMiddle("testing")).To(Equal("t"))
// Expect(GetMiddle("middle")).To(Equal("dd"))
// Expect(GetMiddle("A")).To(Equal("A"))
