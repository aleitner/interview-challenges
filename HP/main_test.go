package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	testcases := []struct {
		num1     []int
		num2     []int
		expected string
	}{
		{[]int{}, []int{}, "0"},
		{[]int{9}, []int{9}, "18"},
		{[]int{1, 2, 3}, []int{}, "123"},
		{[]int{1, 2, 3}, []int{3}, "126"},
		{[]int{1, 0}, []int{1, 0}, "11"},
		{[]int{1}, []int{1, 2, 3}, "322"},
		{[]int{}, []int{1, 2, 3}, "321"},
		{[]int{0, 0, 1}, []int{1, 0, 0}, "2"},
		{[]int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1}, "246912"},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}, "777777"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}, "24691357975308642"},
	}

	for _, tc := range testcases {
		actual := Add(tc.num1, tc.num2)
		assert.Equal(t, tc.expected, actual)
	}

}

func BenchmarkMathAdd0(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = 0 + 0
	}
}

func BenchmarkAdd0(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Add([]int{}, []int{})
	}
}

func BenchmarkMathAdd3Digits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = 123 + 123
	}
}

func BenchmarkAdd3Digits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Add([]int{1, 2, 3}, []int{1, 2, 3})
	}
}

func BenchmarkMathAddLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = 12345678987654321 + 12345678987654321
	}
}

func BenchmarkAddLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Add([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	}
}
