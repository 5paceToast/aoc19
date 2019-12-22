package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// I can't really test stdin/stdout, and this thing is pretty stupid
// but let's at least test immediate mode

func TestImmediate(t *testing.T) {
	var (
		assert = assert.New(t)

		a = NewStubIntCode([]int{1002, 4, 3, 4, 33})
		b = NewStubIntCode([]int{1101, 100, -1, 4, 0})
	)
	a.Run()
	b.Run()

	assert.Equal([]int{1002, 4, 3, 4, 99}, a.State())
	assert.Equal([]int{1101, 100, -1, 4, 99}, b.State())
}
