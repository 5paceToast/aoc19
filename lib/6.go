package lib

// total number of orbits is just distance to the root node

// OrbitNode represents an orbiting node
type OrbitNode struct {
	Tag      string
	Parent   *OrbitNode
	Children []*OrbitNode
}

// NewOrbitNode creates a new orbit node
func NewOrbitNode(tag string, parent *OrbitNode) *OrbitNode {
	return &OrbitNode{
		Tag:      tag,
		Parent:   parent,
		Children: nil,
	}
}

// AddChild is a convenience function because I got tired of typing this out
func (r *OrbitNode) AddChild(child *OrbitNode) {
	if child.Parent == nil {
		child.Parent = r
	}
	r.Children = append(r.Children, child)
}

// Orbits counts the indirect and direct number of orbits of a given node
func (r *OrbitNode) Orbits() int {
	res := 0
	current := r

	for current.Parent != nil {
		res++
		current = current.Parent
	}

	return res
}

// Root finds the root node and returns a pointer to it
func (r *OrbitNode) Root() *OrbitNode {
	p := r
	for p.Parent != nil {
		p = p.Parent
	}
	return p
}

// Walk will walk the tree from the given node, returns a channel
func (r *OrbitNode) Walk() <-chan *OrbitNode {
	c := make(chan *OrbitNode)
	go func() {
		defer close(c)
		c <- r                 // append the node
		if r.Children != nil { // if the node has children
			for _, child := range r.Children {
				for v := range child.Walk() { // recurse into every child
					c <- v
				}
			}
		}
	}()
	return c
}

// Filter filters all nodes in this tree that pass the test f
func (r *OrbitNode) Filter(f func(*OrbitNode) bool) []*OrbitNode {
	root := r.Root()
	var res []*OrbitNode

	for v := range root.Walk() {
		if f(v) {
			res = append(res, v)
		}
	}

	return res
}

// ConnectMagic is an unusual thing purely for the sake of the exercise
// I hope you're happy
func ConnectMagic(parent, child string, set []*OrbitNode) []*OrbitNode {
	var (
		c, p *OrbitNode
	)
	for _, v := range set {
		if v.Tag == child {
			c = v
		}
		if v.Tag == parent {
			p = v
		}
	}

	if p == nil { // the parent does not exist
		p = NewOrbitNode(parent, nil)
		set = append(set, p)
	}
	if c == nil { // the child doesn't exist
		c = NewOrbitNode(child, p)
		set = append(set, c)
	}
	p.AddChild(c)
	return set
}
