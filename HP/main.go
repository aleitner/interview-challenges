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

func Add(num1 []int, num2 []int) string {
	len1 := len(num1)
	len2 := len(num2)

	if len1 == 0 && len2 == 0 {
		return "0"
	}

	var resultSize int
	if len1 > len2 {
		resultSize = len1 + 1
	} else {
		resultSize = len2 + 1
	}

	var result string
	var excess int
	var addByDigitResult int

	for i := 1; i <= resultSize; i++ {
		addByDigitResult, excess = addByDigit(i, num1, num2, excess)
		if addByDigitResult == 0 && excess == 0 {
			break
		}
		result = fmt.Sprintf("%d%s", addByDigitResult, result)
	}

	return result
}

func addByDigit(digit int, num1, num2 []int, previousExcess int) (result, excess int) {
	num1Int := 0
	num2Int := 0

	// if num1 has values and the current digit isn't higher than the array contains
	if len(num1) > 0 && len(num1) >= digit {
		num1Int = num1[len(num1) - digit]
	}

	if len(num2) > 0 && len(num2) >= digit {
		num2Int = num2[digit - 1]
	}

	result = num1Int + num2Int + previousExcess

	if result > 9 {
		excess = result / 10
		result = result - 10
	}

	return result, excess
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
