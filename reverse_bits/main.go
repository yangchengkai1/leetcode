package main

import "fmt"

func main() {
	num := uint32(439329)
	fmt.Printf("%032b\n", num)
	fmt.Printf("%b\n", reverseBits(num))
	fmt.Printf("%b\n", reverseBitsFirst(num))
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

// Another version
func reverseBitsFirst(num uint32) uint32 {
	bits, count := uint32(0), uint32(0)
	mask := uint32(2147483648)

	for count < 32 {
		mm := num & mask
		if mm == 0 {
			count++
			mask >>= 1
			continue
		}

		bits |= 1 << count
		count++
		mask >>= 1
	}

	return bits
}
