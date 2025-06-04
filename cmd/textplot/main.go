package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func ToString(matrix [][]int) string {

	var sb strings.Builder
	for _, row := range matrix {
		for _, col := range row {
			if col == 1 {
				sb.WriteRune('\u2588')
			} else {
				sb.WriteRune(' ')
				//sb.WriteRune('\u2591')
			}
		}
		sb.WriteRune('\n')
	}
	sb.WriteRune('\n')
	return sb.String()
}

func Counts(matrix [][]int) {
	total := 0
	ones := 0
	for _, row := range matrix {
		for _, col := range row {
			total++
			if col == 1 {
				ones++
			}
		}
	}
	fmt.Printf("%f\n", 100.0*float64(ones)/float64(total))
}

func main() {
	flag.Parse()
	args := flag.Args()
	dat, err := os.ReadFile(args[0])

	for i := 0; i < len(dat); i++ {
		switch dat[i] {
		case '{':
			dat[i] = '['
		case '}':
			dat[i] = ']'
		}
	}

	if err != nil {
		log.Fatalf("Unable to read %q", args[0])
	}
	var matrix [][]int

	err = json.Unmarshal(dat, &matrix)
	if err != nil {
		log.Fatalf("Unable to read JSON")
	}

	fmt.Println(ToString(matrix))

	Counts(matrix)
}
