package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"toast.cafe/x/aoc19/lib"
)

func checkIteration(noun, verb, res int, prog []int) bool {
	b := make([]int, len(prog))
	copy(b, prog)
	b[1] = noun
	b[2] = verb

	lib.RunIntcode(b)

	if b[0] == res {
		return true
	}
	return false
}

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

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if checkIteration(i, j, 19690720, prog) {
				fmt.Printf("%d", 100*i+j)
				os.Exit(0)
			}
		}
	}
}
