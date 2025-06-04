package main

import (
	"fmt"
	"math/bits"
	"time"
)

const m = 3
const width = 64

func Slice(v []uint64) []uint64 {
	m := len(v)
	width := 64
	nv := make([]uint64, m, m)

	for i, val := range v {
		for j := 0; j < width; j++ {
			/*


				a0
				a3
				a
							pos = i * width + j --> NORMAL
					        pos = j * m + i
					(3 x 64)
					a0 a1 a2(0,2) a3 ... a63
					b0 b1 b2(1,2) b3 ... b63

					( 64 x 3 )
					a0 b0 c0
					a1 b1 c1
					a2 b2 c2

					a0 a3 a6 a9 a12
					a1 a4 a7 a10 a13
					a2 a5


					bucket = pos % 3
					bit = pos - bucket
			*/
			if val&0x01 == 1 {
				// normal position is
				//  i * width + j
				pos := i*64 + j
				bucket := pos % 3
				bit := pos / 3
				nv[bucket] = nv[bucket] | (1 << bit)
			}
			val = val >> 1
		}
	}
	return nv
}

func Equals(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Unslice(v []uint64) []uint64 {
	m := len(v)
	width := 64
	nv := make([]uint64, m, m)

	for i, val := range v {
		for j := 0; j < width; j++ {
			if val&0x01 == 1 {
				pos := j*m + i //i*64 + j //j*m + i
				bucket := pos / width
				bit := pos % width
				nv[bucket] = nv[bucket] | (1 << bit)
			}
			val = val >> 1
		}
	}
	return nv
}
func ToString(v []uint64) string {
	m := len(v)
	width := 64
	buf := make([]byte, m*width, m*width)
	i := 0
	for _, val := range v {
		for j := 0; j < width; j++ {
			if val&0x01 == 1 {
				buf[i] = '*'
			} else {
				buf[i] = ' '
			}
			val = val >> 1
			i++
		}
	}
	return string(buf)
}

func Rule(v []uint64) []uint64 {
	m := len(v)
	nv := make([]uint64, m, m)
	return RuleIter(v, nv)
	/*
	   nv[0] = bits.RotateLeft64(v[m-1], 1) ^ (v[0] | v[1])

	   	for i := 1; i < m-1; i++ {
	   		nv[i] = v[i-1] ^ (v[i] | v[i+1])
	   	}

	   nv[m-1] = v[m-2] ^ (v[m-1] | bits.RotateLeft64(v[0], width-1))
	   return nv
	*/
}

func RuleIter(in, out []uint64) []uint64 {
	m := len(in)
	width := 64

	out[0] = bits.RotateLeft64(in[m-1], 1) ^ (in[0] | in[1])
	for i := 1; i < m-1; i++ {
		out[i] = in[i-1] ^ (in[i] | in[i+1])
	}
	out[m-1] = in[m-2] ^ (in[m-1] | bits.RotateLeft64(in[0], width-1))
	return out
}
func New(numbits int) []uint64 {
	width := 64
	v := make([]uint64, numbits/width, numbits/width)
	return v
}

func FromString(s string, numbits int) []uint64 {
	width := 64
	v := make([]uint64, numbits/width, numbits/width)

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			bucket := i / width
			bit := i % width
			v[bucket] = v[bucket] | (1 << bit)
		}
	}
	return v
}

func main() {
	bitslen := 64 * 3
	buf := make([]byte, bitslen, bitslen)
	for i := 0; i < len(buf); i++ {
		buf[i] = ' '
	}
	buf[bitslen/2] = '*'

	c := FromString(string(buf), bitslen)
	next := New(bitslen)

	start := time.Now()
	c = Slice(c)
	//fmt.Println(ToString(Unslice(c)))
	for i := 0; i < 20000000; i++ {
		RuleIter(c, next)
		c, next = next, c
		//fmt.Println(ToString(Unslice(c)))
	}
	fmt.Println(ToString(Unslice(c)))
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)

}
