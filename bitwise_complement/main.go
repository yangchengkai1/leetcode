package main

import "fmt"

func main() {
	num := 324
	fmt.Printf("十进制数 %d = 二进制数 %b ,它的反码为 十进制 %d , 二进制 %b\n", num, num, bitwiseComplement(num), bitwiseComplement(num))
}

func bitwiseComplement(N int) int {
	num := 1
	for num < N {
		num = (num << 1) + 1
	}

	return N ^ num
}

/* another version
func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}

	var mask = 1
	n := N
	for N > 0 {
		N = N / 2
		mask = mask * 2
	}

	return (mask - 1) ^ n
}
*/
