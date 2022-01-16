package main

import "fmt"

func maxDistToClosest(seats []int) int {
	j := 0
	var start, end, max int

	for i := 0; i < len(seats); i++ {
		fmt.Println("j", j, "i", i, "start", start, "end", end, "max", max, len(seats[j:i]))
		if seats[i] == 1 {
			// k = i
			if len(seats[j:i]) > max {
				start = j
				end = i - 1
				max = len(seats[j:i])
			}
			j = i + 1
		}
	}

	return j
}

func main() {
	res := maxDistToClosest([]int{1, 0, 0, 0})
	fmt.Println(res)
}
