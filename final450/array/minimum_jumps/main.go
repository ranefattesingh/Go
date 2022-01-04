/* Minimum no. of Jumps to reach end of an array */

package main

import "fmt"

func findMaxJumps(arr []int) int {
	jumps := 1
	for i := arr[0]; i < len(arr) - 1; i += arr[i] {
		jumps++
	} 
	return jumps
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

	jumps := findMaxJumps(arr)

	fmt.Printf("Minimum jumps needed are: %d", jumps);
}
