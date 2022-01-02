/* Constant block */

package main

import "fmt"

const (
	PRODUCT  = "Mobile"
	QUANTITY = 50
	PRICE    = 50.50
	IN_STOCK = true
)

func main() {
	fmt.Println(QUANTITY)
	fmt.Println(PRICE)
	fmt.Println(PRODUCT)
	fmt.Println(IN_STOCK)
}
