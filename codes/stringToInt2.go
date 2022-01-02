/* Example to demonstrate ParseInt */

package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	strVal := "100"

	// 8 bit integer i.e int8
	intVar, err := strconv.ParseInt(strVal, 0, 8)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(intVar, reflect.TypeOf(intVar))

	// 16 bit integer i.e int16
	intVar, err = strconv.ParseInt(strVal, 0, 16)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(intVar, reflect.TypeOf(intVar))

	// 32 bit integer i.e int32
	intVar, err = strconv.ParseInt(strVal, 0, 32)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(intVar, reflect.TypeOf(intVar))

	// 64 bit integer i.e int64
	intVar, err = strconv.ParseInt(strVal, 0, 64)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(intVar, reflect.TypeOf(intVar))

}
