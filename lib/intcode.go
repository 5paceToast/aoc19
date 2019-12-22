package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

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

func modecalc(in int) [3]bool {
	res := [3]bool{}
	if in > 100 {
		res[2] = in/10000 > 0
		if res[2] {
			in -= 10000
		}
		res[1] = in/1000 > 0
		if res[1] {
			in -= 1000
		}
		res[0] = in/100 > 0
	} else {
		res[0] = false
		res[1] = false
		res[2] = false
	}

	return res
}

// IntCode represents a preprogrammed IntCode CPU
type IntCode struct {
	prog []int
	pos  int

	in  io.Reader
	out io.Writer
}

// In returns the input stream
func (r *IntCode) In() io.Reader { return r.in }

// Out returns the output stream
func (r *IntCode) Out() io.Writer { return r.out }

// Pos returns the current pointer position
func (r *IntCode) Pos() int { return r.pos }

// State returns the current CPU state
func (r *IntCode) State() []int { return r.prog }

// NewIntCode constructs a new IntCode, including those for programmatic usage
func NewIntCode(prog []int, in io.Reader, out io.Writer) *IntCode {
	// please don't mutate my input
	b := make([]int, len(prog))
	copy(b, prog)
	// ok thanks
	res := IntCode{
		prog: b,
		pos:  0,
		in:   os.Stdin,
		out:  os.Stdout,
	}
	return &res
}

// NewIOIntCode constructs an IntCode for use in a terminal
func NewIOIntCode(prog []int) *IntCode {
	return NewIntCode(prog, os.Stdin, os.Stdout)
}

// NewStubIntCode constructs a stub IntCode with a shared in/out buffer (for tests)
func NewStubIntCode(prog []int) *IntCode {
	var buf bytes.Buffer
	return NewIntCode(prog, &buf, &buf)
}

func (r *IntCode) modeget(param int, mode bool) int {
	if mode {
		return param
	}
	return r.prog[param]
}

// RunOp runs the currently selected operation and moves the pointer
func (r *IntCode) RunOp() {
	var (
		in  = r.prog[r.pos]
		op  = in % 100
		mod = modecalc(in)

		ins = 4
	)
	defer func() { r.pos += ins }() // don't bother with it

	var (
		get  = func(offset int) int { return r.prog[r.pos+offset] }
		getm = func(offset int) int { return r.modeget(get(offset), mod[offset-1]) }
	)

	switch op {
	case opAdd:
		var (
			a    = getm(1)
			b    = getm(2)
			npos = get(3)
		)
		r.prog[npos] = a + b // never in immediate mode
	case opMul:
		var (
			a    = getm(1)
			b    = getm(2)
			npos = get(3)
		)
		r.prog[npos] = a * b // never in immediate mode
	case opSav:
		scan := bufio.NewScanner(r.in)
		scan.Scan()
		text := scan.Text()
		val, _ := strconv.Atoi(text)
		npos := get(1)
		r.prog[npos] = val // never in immediate mode
		ins = 2
	case opOut:
		out := getm(1)
		fmt.Fprintln(r.out, out) // REEEE STATE
		ins = 2
	case opJIT:
		var (
			a = getm(1)
			b = getm(2)
		)
		if a != 0 {
			ins = b - r.pos
		} else {
			ins = 3
		}
	case opJIF:
		var (
			a = getm(1)
			b = getm(2)
		)
		if a == 0 {
			ins = b - r.pos
		} else {
			ins = 3
		}
	case opLT:
		var (
			a    = getm(1)
			b    = getm(2)
			posc = get(3)
		)
		if a < b {
			r.prog[posc] = 1
		} else {
			r.prog[posc] = 0
		}
	case opEQ:
		var (
			a    = getm(1)
			b    = getm(2)
			posc = get(3)
		)
		if a == b {
			r.prog[posc] = 1
		} else {
			r.prog[posc] = 0
		}
	case opEnd:
		ins = 0
	}
}

// Run starts the IntCode CPU and runs the program to termination
func (r *IntCode) Run() {
	for {
		r.RunOp()
		if r.prog[r.pos] == opEnd {
			break
		}
	}
}
