// https://www.codewars.com/kata/5513795bd3fafb56c200049e/train/go

package main

import (
	"strings"
)


func AbbrevName(name string) string {
	return strings.ToUpper(string(name[0]) + "." + string(name[strings.Index(name," ")+1]))
}

func main() {
	result := AbbrevName("Sam Harris")
	println(result)
}


// Expect(AbbrevName("Sam Harris")).To(Equal("S.H"))
// Expect(AbbrevName("Patrick Feenan")).To(Equal("P.F"))
// Expect(AbbrevName("Evan Cole")).To(Equal("E.C"))
// Expect(AbbrevName("P Favuzzi")).To(Equal("P.F"))
// Expect(AbbrevName("David Mendieta")).To(Equal("D.M"))