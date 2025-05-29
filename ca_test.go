package ca

import "testing"

func TestToRule3(t *testing.T) {
	fn := func(a, b, c int) int {
		return a ^ (b | c)
	}
	r := ToRule3(fn)
	if r != uint32(30) {
		t.Fatalf("Failed")
	}
}

// CARule30 is 66847740
func TestToRule5(t *testing.T) {
	fn := func(a, b, c, d, e int) int {
		//return b ^ (c | d)

		// 16843262
		//return (a | b) ^ (c | (d | e))

		// 4127787510 or mirror 2526646686
		//return (a ^ b) ^ (c | (d ^ e))

		// 151587318
		//return (a | b) ^ (c | (d ^ e))

		// FAIL
		//return (a | b) ^ (c ^ (d | e))

		// 518119710
		//return (a ^ b) ^ (c ^ (d | e))

		// 1094795710
		//return (a | b) ^ ((c ^ d) | e)

		// 3191947710
		//return (a ^ b) ^ ((c ^ d) | e)

		return (a | b) ^ (c ^ d ^ e)

	}
	r := ToRule5(fn)
	if r != uint32(30) {
		t.Fatalf("Got %d", r)
	}
}
