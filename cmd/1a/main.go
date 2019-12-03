package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"toast.cafe/x/aoc19/lib"
)

func main() {
	var sum int64
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		b, err := strconv.ParseInt(reader.Text(), 0, 0)
		if err != nil {
			break
		}
		r := lib.FuelFromMass(int(b))
		fmt.Println(r)
		sum += int64(r)
	}

	fmt.Println("--------")
	fmt.Println(sum)

	if err := reader.Err(); err != nil {
		fmt.Println(err)
	}
}
