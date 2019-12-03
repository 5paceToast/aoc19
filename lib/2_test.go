package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunOpcode(t *testing.T) {
	assert.Equal(t, []int{2, 0, 0, 0, 99},
		RunOpcode(0, []int{1, 0, 0, 0, 99}), "they should be equal")
	assert.Equal(t, []int{2, 3, 0, 6, 99},
		RunOpcode(0, []int{2, 3, 0, 3, 99}), "they should be equal")
	assert.Equal(t, []int{2, 4, 4, 5, 99, 9801},
		RunOpcode(0, []int{2, 4, 4, 5, 99, 0}), "they should be equal")
}

func TestRunIntcode(t *testing.T) {
	assert.Equal(t, []int{2, 0, 0, 0, 99},
		RunIntcode([]int{1, 0, 0, 0, 99}), "they should be equal")
	assert.Equal(t, []int{2, 3, 0, 6, 99},
		RunIntcode([]int{2, 3, 0, 3, 99}), "they should be equal")
	assert.Equal(t, []int{2, 4, 4, 5, 99, 9801},
		RunIntcode([]int{2, 4, 4, 5, 99, 0}), "they should be equal")
	assert.Equal(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		RunIntcode([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}), "they should be equal")
}
