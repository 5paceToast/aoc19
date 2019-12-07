package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"toast.cafe/x/aoc19/lib"
)

func main() {
	var set []*lib.OrbitNode
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		text := reader.Text()
		arr := strings.Split(text, ")")
		set = lib.ConnectMagic(arr[0], arr[1], set)
	}

	root := set[0].Root()
	you := root.Filter(func(n *lib.OrbitNode) bool { return n.Tag == "YOU" })[0]
	san := root.Filter(func(n *lib.OrbitNode) bool { return n.Tag == "SAN" })[0]
	youp := you.Parent
	sanp := san.Parent

	d := lib.OrbitDistance(youp, sanp)
	fmt.Println(d)
}
