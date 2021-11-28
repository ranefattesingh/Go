/* Find the maximum and minimum element in an array (Linear Search) */

package main

import "fmt"

type MinMax struct {
	Min int
	Max int
}

func findMinMax(arr []int) MinMax {
	minmax := MinMax{ Min: arr[0], Max: arr[0], }
	for i := 1; i < len(arr); i++ {

		if arr[i] < minmax.Min {
			minmax.Min = arr[i]
		}
		
		if arr[i] > minmax.Max {
			minmax.Max = arr[i]
		}
	}

	return minmax
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

	var minmax = findMinMax(arr)

	fmt.Printf("Minimum element in array is %d\r\n", minmax.Min)
	fmt.Printf("Maximum element in array is %d\r\n", minmax.Max)
}