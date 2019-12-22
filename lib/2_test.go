package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunOpcode(t *testing.T) {
	var (
		assert = assert.New(t)

		a = NewStubIntCode([]int{1, 0, 0, 0, 99})
		b = NewStubIntCode([]int{2, 3, 0, 3, 99})
		c = NewStubIntCode([]int{2, 4, 4, 5, 99, 0})
	)

	a.Run()
	b.Run()
	c.Run()

	assert.Equal([]int{2, 0, 0, 0, 99}, a.State())
	assert.Equal([]int{2, 3, 0, 6, 99}, b.State())
	assert.Equal([]int{2, 4, 4, 5, 99, 9801}, c.State())
}

func TestRunIntcode(t *testing.T) {
	var (
		assert = assert.New(t)

		a = NewStubIntCode([]int{1, 0, 0, 0, 99})
		b = NewStubIntCode([]int{2, 3, 0, 3, 99})
		c = NewStubIntCode([]int{2, 4, 4, 5, 99, 0})
		d = NewStubIntCode([]int{1, 1, 1, 4, 99, 5, 6, 0, 99})
	)
	a.Run()
	b.Run()
	c.Run()
	d.Run()

	assert.Equal([]int{2, 0, 0, 0, 99}, a.State())
	assert.Equal([]int{2, 3, 0, 6, 99}, b.State())
	assert.Equal([]int{2, 4, 4, 5, 99, 9801}, c.State())
	assert.Equal([]int{30, 1, 1, 4, 2, 5, 6, 0, 99}, d.State())
}
