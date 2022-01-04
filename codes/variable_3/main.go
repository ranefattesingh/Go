/* Variable declaration with type inference */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i = 7
	var s = "India"

	fmt.Println(i, "of type", reflect.TypeOf(i))
	fmt.Println(s, "of type", reflect.TypeOf(s))
}
