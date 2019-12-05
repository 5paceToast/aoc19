package main

import (
	"flag"
	"fmt"

	"toast.cafe/x/aoc19/lib"
)

func main() {
	start := flag.Int("start", 0, "number to start at")
	end := flag.Int("end", 0, "number to end at")
	flag.Parse()

	filter := func(num int) bool {
		return lib.FilterSixDigits(num) &&
			lib.FilterTwoAdjacent(num) &&
			lib.FilterNonDecreasing(num)
	}

	res := lib.FilterInt(*start, *end, filter)

	fmt.Println(len(res))
}
