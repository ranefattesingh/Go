/* find Largest sum contiguous Subarray [V. IMP] */

package main

import "fmt"

func findLargestSubarraySum(arr []int) int {
	var maxSum, sum int = arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		sum = sum + arr[i]
		if sum > maxSum {
			maxSum = sum
		}
	} 

	return maxSum
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

	largestSum := findLargestSubarraySum(arr)

	fmt.Printf("Largest sub array sum is: %d", largestSum);
}

