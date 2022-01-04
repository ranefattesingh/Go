package main

import (
	"fmt"
)

func bitwiseComplement(n int) int {
	mask := 0x01
	for n&mask < n {
		mask = mask<<1 | 0x01
	}

	return n ^ mask
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)
	d := bitwiseComplement(n)
	fmt.Println(fmt.Sprintf("Compliment of %d is %d", n, d))
}
