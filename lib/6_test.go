package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrbits(t *testing.T) {
	assert := assert.New(t)

	parent := NewOrbitNode("parent", nil)
	child := NewOrbitNode("child", nil) // test AddChild too
	parent.AddChild(child)

	// test adding orbits
	assert.Equal([]*OrbitNode{child}, parent.Children)
	assert.Equal(child.Parent, parent)

	// test counting orbits
	assert.Equal(1, child.Orbits())
	assert.Equal(0, parent.Orbits())

	// test finding root
	assert.Equal(parent, child.Root())
	assert.Equal(parent, parent.Root())
}

func TestOrbitWalk(t *testing.T) {
	assert := assert.New(t)

	var (
		a = NewOrbitNode("a", nil)
		b = NewOrbitNode("b", nil)
		c = NewOrbitNode("c", nil)
		d = NewOrbitNode("d", nil)
	)
	a.AddChild(b)
	a.AddChild(c)
	c.AddChild(d)
	// A -> B | C -> D

	genfilter := func(tag string) func(*OrbitNode) bool {
		return func(n *OrbitNode) bool {
			if n.Tag == tag {
				return true
			}
			return false
		}
	}

	ares := d.Filter(genfilter("a"))
	dres := a.Filter(genfilter("d"))

	assert.Equal([]*OrbitNode{a}, ares)
	assert.Equal([]*OrbitNode{d}, dres)
}

func TestMagicConnect(t *testing.T) {
	var set []*OrbitNode

	set = ConnectMagic("a", "b", set)
	set = ConnectMagic("a", "c", set)
	set = ConnectMagic("c", "d", set)
	set = ConnectMagic("a", "d", set) // technically not needed for the exercise but I can

	root := set[0].Root()
	a := root.Filter(func(n *OrbitNode) bool { return n.Tag == "a" })[0]

	assert.Equal(t, a, root)
}

func TestOrbitData(t *testing.T) {
	assert := assert.New(t)
	var set []*OrbitNode

	set = ConnectMagic("COM", "B", set)
	set = ConnectMagic("B", "C", set)
	set = ConnectMagic("C", "D", set)
	set = ConnectMagic("D", "E", set)
	set = ConnectMagic("E", "F", set)
	set = ConnectMagic("B", "G", set)
	set = ConnectMagic("G", "H", set)
	set = ConnectMagic("D", "I", set)
	set = ConnectMagic("E", "J", set)
	set = ConnectMagic("J", "K", set)
	set = ConnectMagic("K", "L", set)

	com := set[0].Root()
	d := com.Filter(func(n *OrbitNode) bool { return n.Tag == "D" })[0]
	l := com.Filter(func(n *OrbitNode) bool { return n.Tag == "L" })[0]

	assert.Equal(0, com.Orbits())
	assert.Equal(3, d.Orbits())
	assert.Equal(7, l.Orbits())

	var sum int
	for v := range com.Walk() { // aka Root
		sum += v.Orbits()
	}
	assert.Equal(42, sum)
}
