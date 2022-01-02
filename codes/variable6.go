/* Scope of variable */

package main

import "fmt"

var s = "India"

func main() {
	fmt.Println(s)
	x := true

	if x {
		y := 1

		if x != false {
			fmt.Println(s)
			fmt.Println(x)
			fmt.Println(y)
		}
	}
	fmt.Println(x)
}
