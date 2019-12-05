package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuelFromMass(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2, FuelFromMass(12))
	assert.Equal(2, FuelFromMass(14))
	assert.Equal(654, FuelFromMass(1969))
	assert.Equal(33583, FuelFromMass(100756))
}

func TestRecursiveFuelFromMass(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2, RecursiveFuelFromMass(14))
	assert.Equal(966, RecursiveFuelFromMass(1969))
	assert.Equal(50346, RecursiveFuelFromMass(100756))
}
