// Longest Substring Without Repeating Characters

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var old_s, new_s string
	max := 0
	for _, v := range s {
		loc := get_location(new_s, v)
		if loc != -1 {
			if len(new_s) > len(old_s) {
				old_s = new_s
				max = len(old_s)
			}
			new_s = new_s[loc+1:]
		}
		new_s += string(v)
	}

	if max > len(new_s) {
		return max
	} else {
		return len(new_s)
	}
}

func get_location(s string, r rune) int {
	for i, v := range s {
		if v == r {
			return i
		}
	}
	return -1
}

func main() {
	var s string
	fmt.Print("Enter string: ")
	//fmt.Scanln(&s)
	s = "pwwwkew"
	l := lengthOfLongestSubstring(s)

	fmt.Println("Length is = ", l)
}
