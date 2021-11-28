/* Find the maximum and minimum element in an array (Divide and Conquer - Tournament Method)*/

package main

import "fmt"

type MinMax struct {
	Min int
	Max int
}

func findMinMax(arr []int, low, high int) MinMax {
	if low == high {
		return MinMax{ Min: arr[low], Max: arr[low]}
	} else if high == low + 1 {
		if arr[low] > arr[high] {
			return MinMax{ Min: arr[high], Max: arr[low]}
		} else {
			return MinMax{ Min: arr[low], Max: arr[high]}
		}
	} 
	
	var mid int = (high + low) / 2;
	minmax1 := findMinMax(arr, low, mid)
	minmax2 := findMinMax(arr, mid + 1, high)

	min := minmax1.Min
	if min > minmax2.Min {
		min = minmax2.Min
	}

	max := minmax1.Max
	if max < minmax2.Max {
		max = minmax2.Max
	}

	return MinMax{ Min: min, Max: max, }
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

	var minmax = findMinMax(arr, 0, len(arr) - 1)

	fmt.Printf("Minimum element in array is %d\r\n", minmax.Min)
	fmt.Printf("Maximum element in array is %d\r\n", minmax.Max)
}