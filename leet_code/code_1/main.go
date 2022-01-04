package main

import "fmt"

func twoSum(nums []int, target int) []int {
	s := make(map[int]int)
	for i, v := range nums {
		if j, ok := s[target-v]; ok {
			return []int{j, i}
		}
		s[v] = i
	}
	return nil
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
