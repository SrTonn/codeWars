// https://www.codewars.com/kata/554e4a2f232cdd87d9000038/train/go

package main

import (
	"strings"
)

func DNAStrand(dna string) string {
	r := strings.NewReplacer("A", "T", "T", "A", "C", "G", "G", "C")
	return r.Replace(dna)
}

func main() {
	result := DNAStrand("GTAT")
	println(result)
}

//   Expect(DNAStrand("AAAA")).To(Equal("TTTT"))
//   Expect(DNAStrand("ATTGC")).To(Equal("TAACG"))
//   Expect(DNAStrand("GTAT")).To(Equal("CATA"))