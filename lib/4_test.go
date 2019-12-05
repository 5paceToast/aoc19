package lib

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterInt(t *testing.T) {
	assert := assert.New(t)

	truthy := func(num int) bool { return true }
	falsy := func(num int) bool { return false }

	full := FilterInt(0, 10, truthy)
	empty := FilterInt(0, 10, falsy)
	sort.Ints(full)

	assert.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, full)
	assert.Equal([]int(nil), empty)
}

func TestGenFilterRange(t *testing.T) {
	assert := assert.New(t)
	f := GenFilterRange(0, 2)

	assert.False(f(0))
	assert.True(f(1))
	assert.False(f(2))
}

func TestNumToStuff(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]string{"1", "2", "3", "4"}, NumToStringDigits(1234))
	assert.Equal([]int{1, 2, 3, 4}, NumToDigits(1234))
}

func TestFilterTwoAdjacent(t *testing.T) {
	assert := assert.New(t)

	a := 112233
	b := 123444
	c := 111122
	d := 123456

	assert.True(FilterTwoAdjacent(a))
	assert.True(FilterTwoAdjacent(b))
	assert.True(FilterTwoAdjacent(c))
	assert.False(FilterTwoAdjacent(d))

	assert.True(FilterTwoAdjacentExclusive(a))
	assert.False(FilterTwoAdjacentExclusive(b))
	assert.True(FilterTwoAdjacentExclusive(c))
	assert.False(FilterTwoAdjacentExclusive(d))
}

func TestFilterNonDecreasing(t *testing.T) {
	assert := assert.New(t)

	assert.True(FilterNonDecreasing(123456))
	assert.False(FilterNonDecreasing(654321))
}
