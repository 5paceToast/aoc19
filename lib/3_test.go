package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathPartToInstructions(t *testing.T) {
	assert := assert.New(t)

	u, erru := PathPartToInstructions("U7")
	r, errr := PathPartToInstructions("R6")
	d, errd := PathPartToInstructions("D5")
	l, errl := PathPartToInstructions("L4")
	o, erro := PathPartToInstructions("U1")

	in, errin := PathPartToInstructions("UU")
	id, errid := PathPartToInstructions("X1")

	ur := []path{'U', 'U', 'U', 'U', 'U', 'U', 'U'}
	rr := []path{'R', 'R', 'R', 'R', 'R', 'R'}
	dr := []path{'D', 'D', 'D', 'D', 'D'}
	lr := []path{'L', 'L', 'L', 'L'}
	or := []path{'U'}

	assert.Equal(ur, u)
	assert.Equal(rr, r)
	assert.Equal(dr, d)
	assert.Equal(lr, l)
	assert.Equal(or, o)

	assert.Nil(erru)
	assert.Nil(errr)
	assert.Nil(errd)
	assert.Nil(errl)
	assert.Nil(erro)

	assert.Nil(in)
	assert.Nil(id)
	assert.NotNil(errin)
	assert.NotNil(errid)
}

func TestPathToPoints(t *testing.T) {
	res := []Point{
		Point{1, 0, 1},
		Point{1, 1, 2},
		Point{1, 0, 3},
		Point{1, -1, 4},
		Point{0, -1, 5},
		Point{-1, -1, 6},
	}

	out := PathToPoints([]path{'R', 'U', 'D', 'D', 'L', 'L'})

	assert.Equal(t, res, out)
}

func TestPointUnion(t *testing.T) {
	a := []Point{
		Point{0, 0, 0},
		Point{42, 69, 0},
		Point{5, 5, 0},
		Point{10, 10, 0},
	}
	b := []Point{
		Point{0, 0, 0},
		Point{42, 69, 0},
		Point{12, 5, 0},
	}
	res := []Point{Point{42, 69, 0}}

	assert.Equal(t, res, PointUnion(a, b))
}

func TestManhattanDistance(t *testing.T) {
	assert := assert.New(t)

	a := Point{-1, -1, 0}
	b := Point{1, -1, 0}
	c := Point{-1, 1, 0}
	d := Point{1, 1, 0}

	assert.Equal(0, ManhattanDistance(a, a))
	assert.Equal(2, ManhattanDistance(a, b))
	assert.Equal(2, ManhattanDistance(a, c))
	assert.Equal(4, ManhattanDistance(a, d))

	assert.Equal(0, ManhattanDistance(b, b))
	assert.Equal(4, ManhattanDistance(b, c))
	assert.Equal(2, ManhattanDistance(b, d))

	assert.Equal(0, ManhattanDistance(c, c))
	assert.Equal(2, ManhattanDistance(c, d))

	assert.Equal(0, ManhattanDistance(d, d))

}

func pathappend(s []path, v string) []path {
	res, err := PathPartToInstructions(v)
	if err != nil {
		panic("fix your tests")
	}
	s = append(s, res...)
	return s
}

func TestShortestEverything(t *testing.T) {
	assert := assert.New(t)
	var (
		paa, pab, pba, pbb []path
		taa, tab, tba, tbb []Point
		ta, tb             []Point
		da, db             int
		pa, pb             Point
		target             = Point{0, 0, 0}
	)

	paa = pathappend(paa, "R75")
	paa = pathappend(paa, "D30")
	paa = pathappend(paa, "R83")
	paa = pathappend(paa, "U83")
	paa = pathappend(paa, "L12")
	paa = pathappend(paa, "D49")
	paa = pathappend(paa, "R71")
	paa = pathappend(paa, "U7")
	paa = pathappend(paa, "L72")

	pab = pathappend(pab, "U62")
	pab = pathappend(pab, "R66")
	pab = pathappend(pab, "U55")
	pab = pathappend(pab, "R34")
	pab = pathappend(pab, "D71")
	pab = pathappend(pab, "R55")
	pab = pathappend(pab, "D58")
	pab = pathappend(pab, "R83")

	pba = pathappend(pba, "R98")
	pba = pathappend(pba, "U47")
	pba = pathappend(pba, "R26")
	pba = pathappend(pba, "D63")
	pba = pathappend(pba, "R33")
	pba = pathappend(pba, "U87")
	pba = pathappend(pba, "L62")
	pba = pathappend(pba, "D20")
	pba = pathappend(pba, "R33")
	pba = pathappend(pba, "U53")
	pba = pathappend(pba, "R51")

	pbb = pathappend(pbb, "U98")
	pbb = pathappend(pbb, "R91")
	pbb = pathappend(pbb, "D20")
	pbb = pathappend(pbb, "R16")
	pbb = pathappend(pbb, "D67")
	pbb = pathappend(pbb, "R40")
	pbb = pathappend(pbb, "U7")
	pbb = pathappend(pbb, "R15")
	pbb = pathappend(pbb, "U6")
	pbb = pathappend(pbb, "R7")

	taa = PathToPoints(paa)
	tab = PathToPoints(pab)
	tba = PathToPoints(pba)
	tbb = PathToPoints(pbb)

	ta = PointUnion(taa, tab)
	tb = PointUnion(tba, tbb)

	da, pa = ShortestManhattan(target, ta)
	db, pb = ShortestManhattan(target, tb)
	sa := ShortestSteps(ta)
	sb := ShortestSteps(tb)

	assert.Equal(159, da)
	assert.Equal(135, db)
	assert.Equal(610, sa)
	assert.Equal(410, sb)

	assert.NotEqual(target, pa)
	assert.NotEqual(target, pb)
}
