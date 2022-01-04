/* Reverse the array */

package main

import "fmt"

func reverseArray(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		temp := arr[i]
		arr[i] = arr[len(arr) - i - 1]
		arr[len(arr) - i - 1] = temp
	}
}

func main() {
	var (
		n int
		arr []int
	)
	
	fmt.Print("Enter size of the array: ")
	fmt.Scanln(&n)

	if (n <= 0) {
		fmt.Println("Size of the array should be greater than zero!")
		return
	} 

	arr = make([]int, n)

	fmt.Printf("Enter %d integers elements: ", n);
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	reverseArray(arr)

	fmt.Printf("Reverse of array is: ");
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}