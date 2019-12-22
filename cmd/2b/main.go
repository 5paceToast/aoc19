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

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			prog[1] = i
			prog[2] = j
			cpu := lib.NewIOIntCode(prog)
			cpu.Run()
			s := cpu.State()
			//fmt.Println(s[0])
			if s[0] == 19690720 {
				fmt.Printf("%d\n", 100*i+j)
				os.Exit(0)
			}
		}
	}
}
