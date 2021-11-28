/* Find the Union and Intersection of the two sorted arrays. */
package main

import "fmt"

func contains (arr []int, search int) bool {
	for _, val := range arr {
		if val == search {
			return true
		}
	}
	return false
}

func findUnionAndIntersection(arr1, arr2 []int) ([]int, []int) {
	union := make([]int, 0, len(arr1) + len(arr2))
	intersection := make([]int, 0, len(arr1) + len(arr2))
	// fmt.Println(intersection)
	for i := 0; i < len(arr1); i++ {
		if  !contains(arr2, arr1[i]) { 
			if !contains(intersection, arr1[i]) {
				intersection = append(intersection, arr1[i])
			}
		} else {
			if !contains(union, arr1[i]) {
				union = append(union, arr1[i])
			}
		}
	}
	union = append(union, intersection...)
	return union, intersection
}

func main() {
	var (
		n1, n2 int
		arr1, arr2 []int
	)

	fmt.Print("Enter size of array1: ")
	fmt.Scanln(&n1)

	fmt.Print("Enter size of array2: ")
	fmt.Scanln(&n2)

	if n1 <= 0 || n2 <= 0 {
		fmt.Println("Size of the arrays should be greater than zero!")
		return
	}

	arr1 = make([]int, n1)
	fmt.Printf("Enter %d integers elements(array1): ", n1)
	for i := 0; i < n1; i++ {
		fmt.Scanf("%d", &arr1[i])
	}
	fmt.Scanln()

	arr2 = make([]int, n2)
	fmt.Printf("Enter %d integers elements(array2): ", n2)
	for i := 0; i < n2; i++ {
		fmt.Scanf("%d", &arr2[i])
	}

	union, intersection := findUnionAndIntersection(arr1, arr2)

	fmt.Printf("Union of arrays : ")
	for i := 0; i < len(union); i++ {
		fmt.Printf("%d ", union[i])
	}

	fmt.Println()
	fmt.Printf("Intersection of arrays : ")
	for i := 0; i < len(intersection); i++ {
		fmt.Printf("%d ", intersection[i])
	}

}