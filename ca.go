package ca

import (
	"fmt"
	"math/bits"
	"strings"
)

type CARule interface {
	Len() int
	Mirror() CARule
	Reverse() CARule
	FromInt(uint32) CARule
	Int() uint32
	Diagram() string
	String() string

	Step(cur cavec, next cavec)
}

type CA3 uint32

func (r CA3) FromInt(val uint32) CARule {
	if val > 255 {
		panic("CA3 must be in [0,255]")
	}
	return CA3(val)
}

func (r CA3) Len() int {
	return 8
}

func (r CA3) Int() uint32 {
	return uint32(r)
}

func (r CA3) String() string {
	return ruleToBString(uint32(r), 3)
}

func (r CA3) Reverse() CARule {
	return r.FromInt(uint32(bits.Reverse8(uint8(r))))
}

func (r CA3) Mirror() CARule {
	x := uint32(r)
	return r.FromInt((x & 1) |
		(((x >> 1) & 1) << 4) |
		(((x >> 2) & 1) << 2) |
		(((x >> 3) & 1) << 6) |
		(((x >> 4) & 1) << 1) |
		(((x >> 5) & 1) << 5) |
		(((x >> 6) & 1) << 3) |
		(((x >> 7) & 1) << 7))
}

func (r CA3) cell(a, b, c int) int {

	// determine which bit to check
	// this return 0-7
	val := a<<2 | b<<1 | c

	if r&(1<<val) == 0 {
		return 0
	}
	return 1
}

func (c CA3) Step(cur, next cavec) {
	last := len(cur) - 1
	next[0] = c.cell(cur[last], cur[0], cur[1])
	next[last] = c.cell(cur[last-1], cur[last], cur[0])
	for i := 1; i < last; i++ {
		next[i] = c.cell(cur[i-1], cur[i], cur[i+1])
	}
}

func (r CA3) Diagram() string {
	const n = 3
	const max = (1 << n) - 1
	buf := strings.Builder{}

	for i := max; i >= 0; i-- {
		buf.WriteString(fmt.Sprintf("%03b ", i))
	}
	buf.WriteString("\n")
	for i := max; i >= 0; i-- {
		val := (r >> i) & 1
		buf.WriteString(fmt.Sprintf(" %d  ", val))
	}
	buf.WriteString("\n")
	return buf.String()
}

type CA5 uint32

func (r CA5) FromInt(val uint32) CARule {
	return CA5(val)
}
func (r CA5) Len() int {
	return 32
}

func (r CA5) Int() uint32 {
	return uint32(r)
}

func (r CA5) String() string {
	return ruleToBString(uint32(r), 5)
}

func (r CA5) Diagram() string {
	const n = 5
	const max = (1 << n) - 1
	buf := strings.Builder{}

	for i := max; i >= 0; i-- {
		buf.WriteString(fmt.Sprintf("  %2d  ", i))
	}
	buf.WriteString("\n")

	for i := max; i >= 0; i-- {
		buf.WriteString(fmt.Sprintf("%05b ", i))
	}
	buf.WriteString("\n")

	for i := max; i >= 0; i-- {
		val := (r >> i) & 1
		buf.WriteString(fmt.Sprintf("  %d   ", val))
	}
	buf.WriteString("\n")
	return buf.String()
}

func (r CA5) Mirror() CARule {
	x := uint32(r)
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
	return r.FromInt(x)
}

func (r CA5) Reverse() CARule {
	return r.FromInt(bits.Reverse32(r.Int()))
}

func (r CA5) cell(a, b, c, d, e int) int {

	// determine which bit to check
	// this return 0-7
	val := a<<4 | b<<3 | c<<2 | d<<1 | e
	//fmt.Printf("State for %d %d %d ==> %d\n", a, b, c, val)

	if r&(1<<val) == 0 {
		return 0
	}
	return 1
}

func absmod(i, n int) int {
	if i < 0 {
		return n+i
	}
	return i
}

func (r CA5) Step(c, next cavec) {

	if true {
		last := len(c) - 1
		next[0] = r.cell(c[last-1], c[last], c[0], c[1], c[2])
		next[1] = r.cell(c[last], c[0], c[1], c[2], c[3])
		next[last-1] = r.cell(c[last-3], c[last-2], c[last-1], c[last], c[0])
		next[last] = r.cell(c[last-2], c[last-1], c[last], c[0], c[1])
		for i := 2; i < last-1; i++ {
			next[i] = r.cell(c[i-2], c[i-1], c[i], c[i+1], c[i+2])
		}
	} else {
		n := len(c)
		for i := 0; i < n; i++ {
			next[i] = r.cell(c[absmod(i-11, n)], c[(i+3)%n], c[i], c[absmod(i-5, n)], c[(i+7)%n])
		}
	}
}

func swap(n uint32, i, j int) uint32 {
	b1 := (n >> i) & 1
	b2 := (n >> j) & 1
	val := b1 ^ b2
	mask := (val << i) | (val << j)
	return n ^ mask
}

func ruleToBString(r uint32, n int) string {
	var ch byte

	max := (1 << n)
	buf := make([]byte, max, max)
	j := 0
	for i := max - 1; i >= 0; i-- {
		val := (r >> i) & 1
		switch val {
		case 0:
			ch = '0'
		case 1:
			ch = '1'
		}
		buf[j] = ch
		j++
	}
	return string(buf)
}
