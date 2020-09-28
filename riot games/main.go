package main

import (
	"fmt"
	"math/rand"
)

// implement integer division without divide by operator
// 10 / 3 = 3
// 1 / 3 = 0

func main() {
	divident := 10132312312
	divisor := 12

	actual := divide(divident, divisor)
	expected := divident / divisor

	if actual != expected {
		fmt.Printf("divide FAILED: Actual \"%d\" != Expected \"%d\"\n", actual, expected)
	} else {
		fmt.Println("divide PASSED")
	}

	actual = fasterDivide(divident, divisor)
	expected = divident / divisor

	if actual != expected {
		fmt.Printf("fasterDivide FAILED: Actual \"%d\" != Expected \"%d\"\n", actual, expected)
	} else {
		fmt.Println("fasterDivide PASSED")
	}

}

// fasterDivide with complexity of O(log(n))
func fasterDivide(divident, divisor int) (quotient int) {
	// NB: IGNORE THIS IF FOR MEASUREMENTS
	complexity := 1
	defer func() {
		fmt.Printf("fasterDivide Complexity: %d\n", complexity)
	}()

	if divisor > divident {
		return 0
	} else if divisor == divident {
		return 1
	} else if divisor == 1 {
		return divident
	}

	remainder := divident % divisor
	cleanDivide := divident - remainder

	// Initialize Start
	high := divident
	low := divisor
	quotient = rand.Intn(high-low) + low

	for quotient*divisor != cleanDivide {
		complexity += 1

		check := quotient * divisor

		if check > cleanDivide {
			high = quotient - 1
			quotient = rand.Intn(high-low+1) + low
		} else if check < cleanDivide {
			low = quotient + 1
			quotient = rand.Intn(high-low+1) + low
		}
	}

	return quotient
}

// divide with complexity of O(n)
func divide(divident, divisor int) int {
	// NB: IGNORE THIS IF FOR MEASUREMENTS
	complexity := 1
	defer func() {
		fmt.Printf("divide Complexity: %d\n", complexity)
	}()

	if divisor > divident {
		return 0
	} else if divisor == divident {
		return 1
	} else if divisor == 1 {
		return divident
	}

	remainder := divident % divisor
	cleanDivide := divident - remainder

	for i := 0; i < cleanDivide; i++ {
		complexity += 1
		if i*divisor == cleanDivide {
			return i
		}

	}

	return 0
}
