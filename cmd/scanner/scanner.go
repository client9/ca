package main

import (
	"fmt"
	"math/bits"

	"github.com/client9/ca"
	"github.com/client9/randcheck"
)

func Scanner() []uint32 {

	cells := 64 * 3
	cycles1 := 128
	cycles2 := 64 * cycles1
	cycles3 := 64 * cycles2

	/*
		var ruleInst ca.CA3
		onesNeeded := 4
		rulesMin := uint32(0x00)
		rulesMax := uint32(0xFF)
		mask := uint32(0x81)
	*/
	var ruleInst ca.CA5
	onesNeeded := 16
	rulesMin := uint32(0x0100FFFF)
	rulesMax := uint32(0xFFFFFFFF)
	mask := uint32(0x80000001)

	tested := 0
	found := []uint32{}
	var rule uint32
	var skipped uint32

	for rule = rulesMin; rule <= rulesMax; rule++ {
		if rule&mask != 0 {
			skipped++
			continue
		}
		if bits.OnesCount32(rule) != onesNeeded {
			skipped++
			continue
		}

		r := ruleInst.FromInt(rule)

		// did we already do this?
		if r.Mirror().Int() < rule {
			continue
		}
		tested++

		randbit := make([]int, 0, cycles1)
		c := ca.NewSingle(cells)
		next := ca.New(cells)
		randbit = append(randbit, ca.CenterCell(c))

		for j := 1; j < cycles1; j++ {
			r.Step(c, next)
			c, next = next, c

			randbit = append(randbit, ca.CenterCell(c))
		}

		if err := randcheck.RunAll(randbit); err != nil {
			//fmt.Printf("Rule %d - %v\n", rule, err)
			continue
		}
		for j := 1; j < cycles2-cycles1; j++ {
			r.Step(c, next)
			c, next = next, c

			randbit = append(randbit, ca.CenterCell(c))
		}
		if err := randcheck.RunAll(randbit); err != nil {
			//fmt.Printf("[2] Rule %d passed at %d but at %d - %v\n", rule, cycles1, cycles2, err)
			continue
		}
		for j := 1; j < cycles3-cycles2-cycles1; j++ {
			r.Step(c, next)
			c, next = next, c

			randbit = append(randbit, ca.CenterCell(c))
		}
		if err := randcheck.RunAll(randbit); err != nil {
			fmt.Printf("[3] Rule %d passed at %d but at %d - %v\n", rule, cycles1, cycles3, err)
			continue
		}
		found = append(found, rule)
		fmt.Printf("[%d] RULE %d IS INTERESTING\n", len(found), rule)

		printSample(r)
		//fmt.Println(r.Diagram())
		//printRule5(rule)
	}
	fmt.Printf("tested: %d, found %d\n", tested, len(found))
	return found
}

func printSample(r ca.CARule) {
	fmt.Printf("\n\n===== RULE %d [ %s ] ================\n\n", r.Int(), r.String())
	fmt.Println(r.Diagram())
	fmt.Printf("Mirror is %d\n", r.Mirror().Int())

	if true {
		cycles := 100
		c := ca.NewSingle(64 * 2)
		next := ca.New(64 * 2)
		fmt.Println(ca.ToString(c))
		for i := 0; i < cycles; i++ {
			r.Step(c, next)
			c, next = next, c
			fmt.Println(ca.ToString(c))
		}
	}
}

func main() {
	Scanner()
}
