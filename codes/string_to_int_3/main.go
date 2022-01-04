package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	var intVal int
	strVar := "100 200"

	_, err := fmt.Sscan(strVar, &intVal)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(intVal, reflect.TypeOf(intVal))
}
