package main

import (
	"flag"
	"fmt"

	"github.com/client9/ca"
)

func main() {
	//	var nFlag = flag.Int("n", 5, "Number of neighbors 3 or 5")
	var rFlag = flag.Uint("r", 0, "Rule number")
	var mirrorFlag = flag.Bool("mirror", false, "Use mirror image of rule")
	var reverseFlag = flag.Bool("reverse", false, "Use reverse image of rule")
	var sFlag = flag.Int("s", 20, "Steps to run")
	var cFlag = flag.Int("c", 64*1, "Number of cells")
	flag.Parse()

	var rule ca.CA5
	r := rule.FromInt(uint32(*rFlag))

	if *mirrorFlag == true {
		r = r.Mirror().(ca.CA5)
	}

	if *reverseFlag == true {
		r = r.Reverse().(ca.CA5)
	}
	fmt.Printf("\n\n===== RULE %d [ %s ] ================\n\n", r.Int(), r.String())
	fmt.Println(r.Diagram())
	fmt.Printf("Mirror is %d\n", r.Mirror().Int())

	c := ca.NewSingle(*cFlag)
	next := ca.New(*cFlag)
	fmt.Println(ca.ToString(c))
	for i := 0; i < *sFlag; i++ {
		r.Step(c, next)
		c, next = next, c
		fmt.Println(ca.ToString(c))
	}
}
