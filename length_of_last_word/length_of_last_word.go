package main

import "fmt"

func main() {
	s := "hahahha oooo"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	right := len(s) - 1
	for right >= 0 {
		if string(s[right]) == " " {
			right--
		} else {
			break
		}
	}

	left := right
	for left >= 0 && string(s[left]) != " " {
		left--
	}

	return right - left
}
