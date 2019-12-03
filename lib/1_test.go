package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuelFromMass(t *testing.T) {
	assert.Equal(t, 2, FuelFromMass(12), "they should be equal")
	assert.Equal(t, 2, FuelFromMass(14), "they should be equal")
	assert.Equal(t, 654, FuelFromMass(1969), "they should be equal")
	assert.Equal(t, 33583, FuelFromMass(100756), "they should be equal")
}

func TestRecursiveFuelFromMass(t *testing.T) {
	assert.Equal(t, 2, RecursiveFuelFromMass(14), "they should be equal")
	assert.Equal(t, 966, RecursiveFuelFromMass(1969), "they should be equal")
	assert.Equal(t, 50346, RecursiveFuelFromMass(100756), "they should be equal")
}
