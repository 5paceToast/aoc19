package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"toast.cafe/x/aoc19/lib"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	if err := reader.Err(); err != nil {
		os.Exit(1)
	}
	progstr := reader.Text()
	progarr := strings.Split(progstr, ",")
	var prog []int

	for _, i := range progarr {
		v, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		prog = append(prog, v)
	}

	res := lib.RunIntcode(prog)
	fmt.Printf("%v\n", res)
}
