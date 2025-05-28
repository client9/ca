package ca

import (
	"strings"
)

type cavec []int

func New(n int) cavec {
	return make(cavec, n, n)
}

func NewSingle(n int) cavec {
	c := New(n)
	c[n/2] = 1
	return c
}

func FromString(s string) cavec {
	v := New(len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			v[i] = 1
		}
	}
	return v
}

func CenterCell(c cavec) int {
	return c[len(c)/2]
}

func ToString(c cavec) string {
	var sb strings.Builder
	for i := 0; i < len(c); i++ {
		if c[i] == 1 {
			sb.WriteRune('\u2588')
		} else {
			sb.WriteRune(' ')
			//sb.WriteRune('\u2591')
		}
	}
	/*
		buf := make([]byte, len(c), len(c))
		for i := 0; i < len(c); i++ {
			if c[i] == 1 {
				buf[i] = '*'
			} else {
				buf[i] = ' '
			}
		}
		return string(buf)
	*/
	return sb.String()
}
func Clear(c cavec) cavec {
	for i := 0; i < len(c); i++ {
		c[i] = 0
	}
	return c
}
