package main

import "fmt"

func main() {
	num := uint32(4324325)
	fmt.Printf("num %032b 有 %d 个 1\n", num, hammingWeight(num))
}

func hammingWeight(num uint32) int {
	m := 0

	for num > 0 {
		num &= (num - 1)
		m++
	}

	return m
}

/*   another version
func hammingWeight(num uint32) int {
	var count int

	for num > 0 {
		if num&1 == 1 {
			count++
		}

		num >>= 1
	}

	return count
}*/
