/*  */

package main

import "fmt"

func findLargestSubarray(arr []int) []int {
	var startIndex, endIndex, gcount int = 0, 0, 0
	for i := 0; i < len(arr) - 1; {
		j := i
		count := 0
		for (j + 1) < len(arr) && arr[j] == arr[j + 1] - 1 {
			j++
			count++
		}
		
		if count > gcount {
			gcount = count
			startIndex = i
			endIndex = j + 1
		}

		if j > i {
			i = j
		} else {
			i += 1
		}
	} 

	return arr[startIndex: endIndex]
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

	largestSubArray := findLargestSubarray(arr)

	fmt.Printf("Largest sub array is: ");
	for i := 0; i < len(largestSubArray); i++ {
		fmt.Printf("%d ", largestSubArray[i])
	}
}

