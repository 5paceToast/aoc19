package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"toast.cafe/x/aoc19/lib"
)

func main() {
	var (
		set []*lib.OrbitNode
		sum int
	)
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		text := reader.Text()
		arr := strings.Split(text, ")")
		set = lib.ConnectMagic(arr[0], arr[1], set)
	}
	for v := range set[0].Root().Walk() {
		sum += v.Orbits()
	}
	fmt.Println(sum)
}
