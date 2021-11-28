/* Given an array which consists of only 0, 1 and 2. Sort the array without using any sorting algorithm */

package main

import (
	"fmt"
)

func sort012(arr []int) {

	low := 0
	high := len(arr) - 1
	for i := 0; i < len(arr) - 1; i++ {

		switch arr[i] {
		case 0: 
			if i > 0 && arr[i - 1] == 0 {
				low = i;
			} else {
				temp := arr[i]
				arr[i] = arr[low]
				arr[low] = temp
				low++
			}
		case 2:
			if i < len(arr) && arr[i + 1] == 2 {
				high = i;
			} else {
				temp := arr[i]
				arr[i] = arr[high]
				arr[high] = temp
				high--
			}
		}
	}
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
	fmt.Printf("Enter %d integers elements: ", n)
	for i := 0; i < n; i++ {
		var e int
		fmt.Scanf("%d", &e)
		
		if e == 0 || e == 1 || e == 2 {
			arr[i] = e
			continue
		}

		fmt.Print("You've entered invalid input!")
		break;
	}

	sort012(arr)

	fmt.Printf("Sorted array is: ");
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}

}