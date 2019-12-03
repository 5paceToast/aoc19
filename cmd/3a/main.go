package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"toast.cafe/x/aoc19/lib"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	if err := reader.Err(); err != nil {
		os.Exit(1)
	}
	l1 := reader.Text()

	reader.Scan()
	if err := reader.Err(); err != nil {
		os.Exit(1)
	}
	l2 := reader.Text()

	s1 := strings.Split(l1, ",")
	s2 := strings.Split(l2, ",")

	var i1, i2 []byte
	for _, s := range s1 {
		v, err := lib.PathPartToInstructions(s)
		if err != nil {
			panic("y u gib bad input")
		}
		i1 = append(i1, v...)
	}
	for _, s := range s2 {
		v, err := lib.PathPartToInstructions(s)
		if err != nil {
			panic("y u gib bad input")
		}
		i2 = append(i2, v...)
	}

	p1 := lib.PathToPoints(i1)
	p2 := lib.PathToPoints(i2)

	u := lib.PointUnion(p1, p2)

	res, _ := lib.ShortestManhattan(lib.Point{0, 0, 0}, u)
	fmt.Println(res)
}
