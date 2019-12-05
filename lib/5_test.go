package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// I can't really test stdin/stdout, and this thing is pretty stupid
// but let's at least test immediate mode

func TestImmediate(t *testing.T) {
	assert := assert.New(t)
	var (
		a, _ = RunOpcode(0, []int{1002, 4, 3, 4, 33})
		b, _ = RunOpcode(0, []int{1101, 100, -1, 4, 0})
	)
	assert.Equal([]int{1002, 4, 3, 4, 99}, a)
	assert.Equal([]int{1101, 100, -1, 4, 99}, b)
}
