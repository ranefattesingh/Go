/* Multiple variables on same line */

package main

import "fmt"

func main() {
	var fname, lname = "Fattesingh", "Rane"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 15000

	fmt.Println("Name is", fname, lname)
	fmt.Println("Code is", m+n+o)
	fmt.Println("Item is", item, "price is", price)

}
