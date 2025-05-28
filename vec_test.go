package ca

import (
	"testing"
)

func BenchmarkVec(b *testing.B) {
	var rule CA3
	r := rule.FromInt(30)

	n := 64*3
	current := NewSingle(n)
	next := New(n)
	for b.Loop() {
		r.Step(current, next)
		current, next = next, current
	}
}
