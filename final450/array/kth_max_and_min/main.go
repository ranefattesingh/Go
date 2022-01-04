/* Find the "Kth" max and min element of an array  */

package main

import (
	"fmt"
	"sort"
	"strconv"
)

type MinMax struct {
	Min int
	Max int
}

func findMinMax(arr []int, k int) MinMax {
	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}

	return MinMax{ Min:arr[k], Max: arr[len(arr) -k - 1]}
}

func findOrdinal(k int) string {
	value := strconv.FormatInt(int64(k), 10)
	if k <= 3 || k > 20 {
		if value[len(value) - 1] == '1' {
			return "st"
		} else if value[len(value) - 1] == '2' {
			return "nd"
		} else {
			return "rd"
		}
	}
	return "th"
}

func main() {
	var (
		n, k int
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
	fmt.Scanln()
	fmt.Print("Enter value of k: ")
	fmt.Scanln(&k)

	if k <= 0 || k > n {
		fmt.Println("You've entered invalid value for k!")
		return
	}

	minmax := findMinMax(arr, k - 1)
	ordinal := findOrdinal(k)
	fmt.Printf("%d%s Minimum element in array is %d\r\n", k, ordinal, minmax.Min)
	fmt.Printf("%d%s Maximum element in array is %d\r\n", k, ordinal, minmax.Max)
}