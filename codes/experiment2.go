/* experimental file*/

package main

import (
	"fmt"
	"reflect"
)

func f(args ...interface{}) []interface{} {
	return args
}

func main() {
	val := f(1, 2, 3, "a", "b", 1.2, 2.3, true, false)
	fmt.Println(val...)
	fmt.Println(val)

	for _, v := range val {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}

	nums := []int{1, 2, 3}
	fmt.Println(nums)
}
