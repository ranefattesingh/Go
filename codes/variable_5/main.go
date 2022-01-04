/* Short variable declaration */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Fattesingh Rane"
	fmt.Println("'"+name+"' type of", reflect.TypeOf(name))
}
