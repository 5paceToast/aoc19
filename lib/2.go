package lib

// btw technically this whole thing mutates because slices are not const
// LUL

type op = int

const (
	opAdd op = 1
	opMul op = 2
	opEnd op = 99
)

// RunOpcode runs an Opcode at position pos, mutating prog
func RunOpcode(pos int, prog []int) []int {
	switch prog[pos] {
	case opAdd:
		posa := prog[pos+1]
		posb := prog[pos+2]
		a := prog[posa]
		b := prog[posb]
		npos := prog[pos+3]
		prog[npos] = a + b
	case opMul:
		posa := prog[pos+1]
		posb := prog[pos+2]
		a := prog[posa]
		b := prog[posb]
		npos := prog[pos+3]
		prog[npos] = a * b
	case opEnd:
	}
	return prog
}

// RunIntcode will run the Intcode program until termination
func RunIntcode(prog []int) []int {
	pos := 0

	for {
		prog = RunOpcode(pos, prog)
		pos += 4
		if prog[pos] == opEnd {
			break
		} // look if your program is broken it's not my fault
	}

	return prog
}
