/*use of Atoi Ascii to int*/

package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(intVar, "of type", reflect.TypeOf(intVar))
}
