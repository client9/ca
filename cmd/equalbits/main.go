package main

import (
	"fmt"
	"math/bits"
)

func swap(n uint32, i, j int) uint32 {
	b1 := (n >> i) & 1
	b2 := (n >> j) & 1
	val := b1 ^ b2
	mask := (val << i) | (val << j)
	return n ^ mask
}
func Mirror5(x uint32) uint32 {
	x = swap(x, 1, 16)
	x = swap(x, 2, 8)
	x = swap(x, 3, 26)
	x = swap(x, 5, 20)
	x = swap(x, 6, 12)
	x = swap(x, 7, 28)
	x = swap(x, 9, 18)
	x = swap(x, 11, 26)
	x = swap(x, 13, 22)
	x = swap(x, 15, 30)
	x = swap(x, 19, 25)
	x = swap(x, 23, 29)
	return x
}

// 601,080,390
// 31! / (16! * 16!)
func main() {
	count := 0
	var i uint32
	for i = 0x0000FFFF; i <= 0xFFFF0000; i++ {
		// high and low bits must be 0
		// to be 0 to have "white background" when starting
		//  with both 1, then "background" is back
		//  with 0 and 1, then alternates between black and white every cycle
		if i&(0x80000001) != 0 {
			continue
		}

		// must have half of bits set to 1 (i.e. 16)
		if bits.OnesCount32(i) != 16 {
			continue
		}

		// we did it already
		if Mirror5(i) < i {
			continue
		}

		/*
			// inverse?
			if i ^ 0xFFFFFFFF < i {
				continue
			}
		*/

		/*
			if (i & 1) == 1 {
				continue
			}
			if (i >> 31 & 1 == 1) {
				continue
			}
		*/
		count++
	}
	fmt.Println(count)
}
