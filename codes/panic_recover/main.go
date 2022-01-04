package main

import (
	"fmt"
	"strings"
)

func printWorkTitle(a string) {
	defer func() {
		errorMsg := recover()
		if errorMsg != nil {
			fmt.Println(errorMsg)
		}
	}()

	if strings.Contains(strings.ToLower(a), "student") {
		fmt.Println("Oh so you're a Student!")
	} else if strings.Contains(strings.ToLower(a), "professional") {
		fmt.Println("Wow! you started working")
	} else {
		panic("Sorry!!! what are you?")
	}

	fmt.Println("Ok bye!")
}

func main() {
	printWorkTitle("I am a Alien")
}
