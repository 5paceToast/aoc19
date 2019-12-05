package lib

import (
	"strconv"
	"strings"
	"sync"
)

// ayy we map filter reduce up in here

type filter = func(num int) bool

// FilterInt filters out numbers from start to end that don't pass f
func FilterInt(start int, end int, f filter) []int {
	var (
		res      []int
		resmutex = &sync.Mutex{}
		wg       sync.WaitGroup
	)

	ff := func(t int) {
		if f(t) {
			resmutex.Lock()
			res = append(res, t)
			resmutex.Unlock()
		}
		wg.Done()
	}

	for i := start; i <= end; i++ {
		wg.Add(1)
		go ff(i)
	}
	wg.Wait()

	return res
}

// GenFilterRange generates a filter that checks against a range
func GenFilterRange(min, max int) filter {
	return func(num int) bool {
		if num > min && num < max {
			return true
		}
		return false
	}
}

// FilterSixDigits filters numbers that are not 6 digits long
var FilterSixDigits = GenFilterRange(99999, 1000000)

// NumToStringDigits converts a number to a list of strings of all its digits
func NumToStringDigits(num int) []string {
	s := strconv.Itoa(num)
	return strings.Split(s, "")
}

// NumToDigits converts a number to a list of its individual digits
func NumToDigits(num int) []int {
	c := NumToStringDigits(num)
	res := make([]int, len(c))
	for i := range c {
		res[i], _ = strconv.Atoi(c[i]) // it can't fail in this exercise, trust me
	}
	return res
}

// FilterTwoAdjacent filters numbers that do not have two identical digits
func FilterTwoAdjacent(num int) bool {
	c := NumToStringDigits(num)
	for i := range c {
		if i == 0 {
			continue
		} // skip 0, because we do look-behind to avoid bounds-checking
		if c[i] == c[i-1] {
			return true
		}
	}
	return false
}

// FilterTwoAdjacentExclusive filters numbers that do not have exactly two identical digits
func FilterTwoAdjacentExclusive(num int) bool {
	c := NumToStringDigits(num)
	deepmagic := func(index int) string {
		if index < 0 || index >= len(c) {
			return "" // impossible value, I don't even care at this point
		}
		return c[index]
	}
	for i := range c {
		var (
			before = deepmagic(i - 1)
			a      = deepmagic(i)
			b      = deepmagic(i + 1)
			after  = deepmagic(i + 2)
		)
		if a == b && a != before && a != after {
			return true
		}
	}
	return false
}

// FilterNonDecreasing filters numbers that have any digit which is lower than the prior one
func FilterNonDecreasing(num int) bool {
	d := NumToDigits(num)
	for i := range d {
		if i == 0 {
			continue
		} // skip 0, because we do look-behind to avoid bounds-checking
		if d[i] < d[i-1] {
			return false
		}
	}
	return true
}
