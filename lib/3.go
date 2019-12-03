package lib

import (
	"fmt"
	"strconv"
	"strings"
)

type path = byte

// Point represents a Point on a grid
type Point struct {
	X, Y, Steps int
}

const (
	pathUp    path = 'U'
	pathDown  path = 'D'
	pathLeft  path = 'L'
	pathRight path = 'R'
)

// PathPartToInstructions converts a PathPart like U7 into a list of path instructions
func PathPartToInstructions(pathpart string) ([]path, error) {
	pp := strings.SplitN(pathpart, "", 2)
	t := pp[0][0]
	s, err := strconv.Atoi(pp[1])
	if err != nil {
		return nil, err
	}
	switch t { // error out on invalid entry
	case pathUp:
	case pathDown:
	case pathLeft:
	case pathRight:
	default: // uh oh
		return nil, fmt.Errorf("invalid path direction %q", t)
	}

	res := make([]path, s)
	for i := range res {
		res[i] = t
	}

	return res, nil
}

// PathToPoints converts a set of path instructions to a set of visited points
func PathToPoints(p []path) []Point {
	var res []Point
	loc := Point{0, 0, 0}

	for _, v := range p {
		switch v {
		case pathUp:
			loc.Y++
		case pathDown:
			loc.Y--
		case pathLeft:
			loc.X--
		case pathRight:
			loc.X++
		default:
			continue // no default for you this should already be valid
		}
		loc.Steps++
		res = append(res, loc) // repeats are fine, trust me
	}

	return res
}

// PointUnion calculates the union of two lists of Points
func PointUnion(a, b []Point) []Point {
	var res []Point
	for _, aa := range a {
		if aa.X == 0 && aa.Y == 0 { // center doesn't count, the rules said so
			continue
		}
		for _, bb := range b {
			if aa.X == bb.X && aa.Y == bb.Y {
				steps := aa.Steps + bb.Steps
				p := Point{aa.X, bb.Y, steps}
				res = append(res, p)
			}
		}
	}
	return res
}

// ManhattanDistance calculates the manhattan distance between two points
func ManhattanDistance(a, b Point) int {
	x := a.X - b.X
	y := a.Y - b.Y
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

// ShortestManhattan returns the shortest manhattan distance from center, and the associated point
func ShortestManhattan(target Point, points []Point) (int, Point) {
	if len(points) == 0 {
		return -1, Point{0, 0, 0}
	}
	var (
		resp = points[0]
		resd = ManhattanDistance(target, resp)
	)

	for _, p := range points {
		d := ManhattanDistance(target, p)
		if d < resd {
			resd = d
			resp = p
		}
	}

	return resd, resp
}

// ShortestSteps gets the point with the fewest steps taken
func ShortestSteps(points []Point) int {
	res := points[0].Steps
	for _, p := range points {
		if p.Steps < res {
			res = p.Steps
		}
	}
	return res
}
