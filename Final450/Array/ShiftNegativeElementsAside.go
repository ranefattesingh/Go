/* Move all the negative elements to one side of the array  */

package main

import "fmt"

func shiftNegativeElements(arr []int) {
	next := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			temp := arr[next]
			arr[next] = arr[i]
			arr[i] = temp
			next++
		}
	}
}

func main() {
	var (
		n int
		arr []int
	)

	fmt.Print("Enter size of array: ")
	fmt.Scanln(&n)

	if n <= 0 {
		fmt.Println("Size of the array should be greater than zero!")
		return
	}

	arr = make([]int, n)
	fmt.Printf("Enter %d integers elements: ", n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	shiftNegativeElements(arr)

	fmt.Printf("Array after moving negative elements aside : ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}