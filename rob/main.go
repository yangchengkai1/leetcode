package main

import (
	"fmt"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{

	{
		[]int{226, 174, 214, 16, 218, 48, 153, 131, 128, 17, 157, 142, 88, 43, 37, 157, 43, 221, 191, 68, 206, 23, 225, 82, 54, 118, 111, 46, 80, 49, 245, 63, 25, 194, 72, 80, 143, 55, 209, 18, 55, 122, 65, 66, 177, 101, 63, 201, 172, 130, 103, 225, 142, 46, 86, 185, 62, 138, 212, 192, 125, 77, 223, 188, 99, 228, 90, 25, 193, 211, 84, 239, 119, 234, 85, 83, 123, 120, 131, 203, 219, 10, 82, 35, 120, 180, 249, 106, 37, 169, 225, 54, 103, 55, 166, 124},
		7102,
	},

	{
		[]int{100, 1, 1, 100, 1, 1, 100, 1, 1, 100},
		400,
	},

	{
		[]int{2, 1, 1, 2},
		4,
	},

	{
		[]int{},
		0,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		20,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		25,
	},

	// 可以有多个 testcase
}

func main() {
	for _, tc := range tcs {
		fmt.Println(rob(tc.nums))
	}
}

func rob(nums []int) int {
	// a 对于偶数位上的最大值的记录
	// b 对于奇数位上的最大值的记录
	var a, b int
	for i, v := range nums {
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(a, b+v)
		}
	}

	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}