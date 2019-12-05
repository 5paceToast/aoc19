package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// btw technically this whole thing mutates because slices are not const
// LUL

type op = int

const (
	opAdd op = 1
	opMul op = 2
	opSav op = 3
	opOut op = 4
	opJIT op = 5
	opJIF op = 6
	opLT  op = 7
	opEQ  op = 8
	opEnd op = 99
)

func modeget(prog []int, param int, mode bool) int {
	if mode {
		return param
	} // 1, aka immediate mode
	return prog[param]
}

// RunOpcode runs an Opcode at position pos, mutating prog
func RunOpcode(pos int, prog []int) ([]int, int) {
	in := prog[pos]
	op := in % 100

	ins := 4 // advance by how many? 4 is default for compat

	mod := [3]bool{}
	if in > 100 {
		mod[2] = in/10000 > 0
		if mod[2] {
			in -= 10000
		}
		mod[1] = in/1000 > 0
		if mod[1] {
			in -= 1000
		}
		mod[0] = in/100 > 0
	} else {
		mod[0] = false
		mod[1] = false
		mod[2] = false
	}

	switch op {
	case opAdd:
		posa := prog[pos+1]
		posb := prog[pos+2]
		a := modeget(prog, posa, mod[0])
		b := modeget(prog, posb, mod[1])
		npos := prog[pos+3]
		prog[npos] = a + b // never in immediate mode
	case opMul:
		posa := prog[pos+1]
		posb := prog[pos+2]
		a := modeget(prog, posa, mod[0])
		b := modeget(prog, posb, mod[1])
		npos := prog[pos+3]
		prog[npos] = a * b // never in immediate mode
	case opSav:
		fmt.Print("-> ")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		text := scan.Text()
		val, _ := strconv.Atoi(text)
		npos := prog[pos+1]
		prog[npos] = val // never in immediate mode
		ins = 2
	case opOut:
		npos := prog[pos+1]
		out := modeget(prog, npos, mod[0])
		fmt.Println(out) // REEEE STATE
		ins = 2
	case opJIT:
		posa := prog[pos+1]
		posb := prog[pos+2]
		a := modeget(prog, posa, mod[0])
		b := modeget(prog, posb, mod[1])
		if a != 0 {
			ins = b - pos
		} else {
			ins = 3
		}
	case opJIF:
		posa := prog[pos+1]
		posb := prog[pos+2]
		a := modeget(prog, posa, mod[0])
		b := modeget(prog, posb, mod[1])
		if a == 0 {
			ins = b - pos
		} else {
			ins = 3
		}
	case opLT:
		posa := prog[pos+1]
		posb := prog[pos+2]
		posc := prog[pos+3] // never in immediate mode
		a := modeget(prog, posa, mod[0])
		b := modeget(prog, posb, mod[1])
		if a < b {
			prog[posc] = 1
		} else {
			prog[posc] = 0
		}
	case opEQ:
		posa := prog[pos+1]
		posb := prog[pos+2]
		posc := prog[pos+3] // never in immediate mode
		a := modeget(prog, posa, mod[0])
		b := modeget(prog, posb, mod[1])
		if a == b {
			prog[posc] = 1
		} else {
			prog[posc] = 0
		}
	case opEnd:
		ins = 0
	}
	return prog, ins
}

// RunIntcode will run the Intcode program until termination
func RunIntcode(prog []int) []int {
	pos := 0

	for {
		prog, ins := RunOpcode(pos, prog)
		pos += ins
		if prog[pos] == opEnd {
			break
		} // look if your program is broken it's not my fault
	}

	return prog
}
