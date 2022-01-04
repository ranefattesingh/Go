package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var (
		arr  []int
		n, s int
	)
	fmt.Print("Enter size of array: ")
	fmt.Scanln(&n)

	arr = make([]int, n)
	fmt.Printf("Enter %d elements: ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Enter element to be searched: ")
	fmt.Scanln(&s)

	loc, err := linearSearch(arr, s)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(s, "Found at location", loc)
}

func linearSearch(arr []int, s int) (int, error) {
	for i, e := range arr {
		if e == s {
			return i + 1, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("%d is not present in list", s))
}
