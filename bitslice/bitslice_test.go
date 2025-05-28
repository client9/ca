package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	bitslen := 64 * 3
	buf := make([]byte, bitslen, bitslen)
	for i := 0; i < len(buf); i++ {
		buf[i] = ' '
	}
	buf[0] = '*'
	buf[bitslen/2] = '*'
	buf[bitslen-1] = '*'

	want := string(buf)
	c := FromString(want, bitslen)
	got := ToString(c)

	fmt.Println(want)
	fmt.Println(got)
	if got != want {
		t.Errorf("To/From round trip did not work")
	}
}

func TestSlice(t *testing.T) {

	bitslen := 64 * 3
	buf := make([]byte, bitslen, bitslen)
	for i := 0; i < len(buf); i++ {
		buf[i] = ' '
	}
	buf[0] = '*'
	buf[bitslen/2] = '*'
	buf[bitslen-1] = '*'

	want := string(buf)

	c := FromString(want, bitslen)

	c = Slice(c)
	c = Unslice(c)

	got := ToString(c)

	if got != want {
		t.Errorf("To/From Slice round trip did not work")
	}
}

func BenchmarkRule(b *testing.B) {

	bitslen := 64 * 3
	buf := make([]byte, bitslen, bitslen)
	for i := 0; i < len(buf); i++ {
		buf[i] = ' '
	}
	buf[bitslen/2] = '*'

	c := FromString(string(buf), bitslen)
	c = Slice(c)
	for b.Loop() {
		c = Rule(c)
	}

}
func BenchmarkRuleIter(b *testing.B) {

	bitslen := 64 * 3
	buf := make([]byte, bitslen, bitslen)
	for i := 0; i < len(buf); i++ {
		buf[i] = ' '
	}
	buf[bitslen/2] = '*'

	c := FromString(string(buf), bitslen)
	c = Slice(c)
	next := New(bitslen)

	for b.Loop() {
		RuleIter(c, next)
		next, c = c, next
	}
}
