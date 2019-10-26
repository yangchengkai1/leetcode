package main

import (
	"fmt"
	"sort"
)

func main() {
	children := []int{1, 3, 5, 7}
	biscuit := []int{9, 9, 1, 8}
	fmt.Println(findContentChildren(children, biscuit))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var i, j, res int
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			res++
			i++
		}
		j++
	}

	return res
}
