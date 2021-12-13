package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var indices []int = make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if sum := nums[i] + nums[j]; sum == target {
				indices[0] = i
				indices[1] = j
				i = len(nums)
				break
			}
		}
	}
	return indices
}

func main() {
	var (
		n, target int
		arr       []int
	)
	fmt.Print("Enter size of array: ")
	fmt.Scanln(&n)

	arr = make([]int, n)
	fmt.Printf("Enter %d integer numbers: ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scanln()

	fmt.Print("Enter target: ")
	fmt.Scanln(&target)

	indices := twoSum(arr, target)
	fmt.Println("Sum is ", indices)
}
