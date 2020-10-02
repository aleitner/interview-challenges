// You are given two arbitrarily large numbers,
// stored one digit at a time in a slice.
// The first must be added to the second,
// and the second must be reversed before addition.
//
// The goal is to calculate the sum of those two sets of values.
//
// IMPORTANT NOTE:
// - The input can be any lengths (i.e: it can be 20+ digits long).
// - num1 and num2 can be different lengths.
//
// Sample Inputs:
// num1 = 123456
// num2 = 123456
//
// Sample Output:
// Result: 777777
//
// We would also like to see a demonstration of appropriate unit tests
// for this functionality.

package main

import "fmt"

// Add two number arrays together and print result as string
func Add(num1 []int, num2 []int) string {
	len1 := len(num1)
	len2 := len(num2)

	// If both nums are empty then no math necessary
	if len1 == 0 && len2 == 0 {
		return "0"
	}

	// The resulting number is up to 1 digit higher than the highest number
	var resultSize int
	if len1 > len2 {
		resultSize = len1 + 1
	} else {
		resultSize = len2 + 1
	}

	var result string // Summation of num1 and num2
	var sumDigitExcess int
	var sumDigitResult int

	for i := 1; i <= resultSize; i++ {
		sumDigitResult, sumDigitExcess = sumDigit(i, num1, num2, sumDigitExcess)
		if sumDigitResult == 0 && sumDigitExcess == 0 {
			break
		}
		result = fmt.Sprintf("%d%s", sumDigitResult, result)
	}

	return result
}

// sumDigit will sum the digit of num1 and num2 in addition to the amount carried over from a previous sumDigit with a lower digit value
func sumDigit(digit int, num1, num2 []int, previousExcess int) (result, carryOver int) {
	// Default to 0
	num1DigitValue := 0
	num2DigitValue := 0

	// check if num1 has a value for digit
	if len(num1) > 0 && len(num1) >= digit {
		num1DigitValue = num1[len(num1)-digit]
	}

	// check if num2 has a value for digit
	if len(num2) > 0 && len(num2) >= digit {
		num2DigitValue = num2[digit-1]
	}

	// Sum the two values in addition to the amount carried over from the sum of the lower digit
	result = num1DigitValue + num2DigitValue + previousExcess

	if result > 9 {
		carryOver = result / 10
		result = result - 10
	}

	return result, carryOver
}

func main() {
	num1 := []int{}
	num2 := []int{}

	num1length := 6
	for i := 1; i <= num1length; i++ {
		num1 = append(num1, i)
	}

	num2length := 6
	for i := 1; i <= num2length; i++ {
		num2 = append(num2, i)
	}

	result := Add(num1, num2)

	fmt.Println("Result:", result)
}
