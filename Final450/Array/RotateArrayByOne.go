/* Write a program to cyclically rotate an array by one. */

package main

import "fmt"

func rotate(arr []int) {
	lastElement := arr[len(arr) - 1]
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i - 1]
	}
	arr[0] = lastElement
}

func main() {
	var (
		n int
		arr[] int
	)

	fmt.Print("Enter the size of array: ")
	fmt.Scanln(&n)

	if n <= 0 {
		fmt.Println("Size of the array should be greater than zero!")
		return
	}

	arr = make([]int, n)
	fmt.Printf("Enter %d elements: ", n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	rotate(arr)

	fmt.Print("Array is rotated by one elment: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}