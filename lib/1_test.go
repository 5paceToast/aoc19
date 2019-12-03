package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuelFromMass(t *testing.T) {
	assert.Equal(t, FuelFromMass(12), 2, "they should be equal")
	assert.Equal(t, FuelFromMass(14), 2, "they should be equal")
	assert.Equal(t, FuelFromMass(1969), 654, "they should be equal")
	assert.Equal(t, FuelFromMass(100756), 33583, "they should be equal")
}

func TestRecursiveFuelFromMass(t *testing.T) {
	assert.Equal(t, RecursiveFuelFromMass(14), 2, "they should be equal")
	assert.Equal(t, RecursiveFuelFromMass(1969), 966, "they should be equal")
	assert.Equal(t, RecursiveFuelFromMass(100756), 50346, "they should be equal")
}
