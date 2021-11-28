/* Minimise the maximum difference between heights [V.IMP] */

package main

import (
	"fmt"
	"sort"
)


func findMinimumDifference(arr []int, k int) int {
	sort.Ints(arr)
	max := arr[len(arr) - 1]
	sum := 0

	for i := 0; i < len(arr); i++ {
		if (sum + arr[i] + k) > max {
			arr[i] -= k
		} else {
			arr[i] += k
		}
		sum += arr[i]
	}

	return arr[len(arr) - 1] - arr[0]
}

func main() {
	var (
		n, k int
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

	fmt.Scanln()
	fmt.Print("Enter value of k: ")
	fmt.Scanln(&k)

	if k <= 0 || k > n {
		fmt.Println("You've entered invalid value for k!")
		return
	}

	maxDifference := findMinimumDifference(arr, k)

	fmt.Printf("Maximum differebce is: %d", maxDifference);
}