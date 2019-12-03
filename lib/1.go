package lib

// FuelFromMass calculates the required amount of fuel for a given module to launch, as in scenario 1
func FuelFromMass(mass int) int {
	res := mass / 3
	return res - 2
}

// RecursiveFuelFromMass calculates the require amount of fuel, fuel for the fuel, etc
func RecursiveFuelFromMass(mass int) int {
	var (
		v   = FuelFromMass(mass)
		sum int
	)

	for v > 0 {
		sum += v
		v = FuelFromMass(v)
	}

	return sum
}
