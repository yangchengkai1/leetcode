package main

import "fmt"

func main() {
	num := uint32(439329)
	fmt.Printf("num = %032b,reverse = %b\n", num, reverseBits(num))
}

func reverseBits(num uint32) uint32 {
	bits, count := uint32(0), uint32(32)

	for count > 0 {
		bits |= (num & 1) << (count - 1)
		num >>= 1
		count--
	}
	return bits
}
