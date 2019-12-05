package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunOpcode(t *testing.T) {
	assert := assert.New(t)
	var (
		a, _ = RunOpcode(0, []int{1, 0, 0, 0, 99})
		b, _ = RunOpcode(0, []int{2, 3, 0, 3, 99})
		c, _ = RunOpcode(0, []int{2, 4, 4, 5, 99, 0})
	)
	assert.Equal([]int{2, 0, 0, 0, 99}, a)
	assert.Equal([]int{2, 3, 0, 6, 99}, b)
	assert.Equal([]int{2, 4, 4, 5, 99, 9801}, c)
}

func TestRunIntcode(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{2, 0, 0, 0, 99}, RunIntcode([]int{1, 0, 0, 0, 99}))
	assert.Equal([]int{2, 3, 0, 6, 99}, RunIntcode([]int{2, 3, 0, 3, 99}))
	assert.Equal([]int{2, 4, 4, 5, 99, 9801}, RunIntcode([]int{2, 4, 4, 5, 99, 0}))
	assert.Equal([]int{30, 1, 1, 4, 2, 5, 6, 0, 99}, RunIntcode([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}))
}
