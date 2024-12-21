package main

import (
	"fmt"
	"fuzzing/reverse"
)

func main() {
	input := "skcid tsom eht sekat dna yag si nojuY"
	rev := reverse.Reverse(input)
	doubleRev := reverse.Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
