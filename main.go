package main

import (
	"fmt"
	"math"
)

func main() {

	input1 := [][]int32{
		{5, 3, 4},
		{1, 5, 8},
		{6, 4, 2},
	}
	input2 := [][]int32{
		{4, 9, 2},
		{3, 5, 7},
		{8, 1, 5},
	}
	input3 := [][]int32{
		{4, 8, 2},
		{4, 5, 7},
		{6, 1, 6},
	}

	fmt.Println(formingMagicSquare(input1))
	fmt.Println(formingMagicSquare(input2))
	fmt.Println(formingMagicSquare(input3))
}

func formingMagicSquare(input [][]int32) int32 {
	var result int32
	n := len(input[0])

	diff := diff(input)
	for {
		r, c := findTarget(diff)
		if r == -1 {
			break
		}
		gap := diff[r][c]
		if r == c && r+c == n-1 {
			gap = int32(math.Round(float64(gap) / 4))
		} else if r == c || r+c == n-1 {
			gap = int32(math.Round(float64(gap) / 3))
		} else {
			gap = int32(math.Round(float64(gap) / 2))
		}
		for i := 0; i < n; i++ {
			if r == c {
				diff[i][i] -= gap
			}
			if r+c == n-1 {
				diff[i][n-i-1] -= gap
			}
			diff[r][i] -= gap
			diff[i][c] -= gap

		}
		result += absolute(gap)
	}
	return result
}

func diff(input [][]int32) [][]int32 {
	n := len(input[0])

	diff := make([][]int32, n)
	for i := range diff {
		diff[i] = make([]int32, n)
	}

	total := int32((n*n*n + n) / 2)
	var sumX int32
	var sumY int32
	for r := 0; r < n; r++ {
		var sumR int32
		var sumC int32
		for c := 0; c < n; c++ {
			sumR += input[r][c]
			sumC += input[c][r]
			if r == c {
				sumX += input[r][c]
			}
			if r+c == n-1 {
				sumY += input[r][c]
			}
		}

		diffR := total - sumR
		diffC := total - sumC
		for c := 0; c < n; c++ {
			if diffR != 0 {
				diff[r][c] += diffR
			} else {
				diff[r][c] += diffR
			}
			if diffC != 0 {
				diff[c][r] += diffC
			} else {
				diff[c][r] += diffC
			}
		}
	}
	diffX := total - sumX
	diffY := total - sumY
	for i := 0; i < n; i++ {
		if diffX != 0 {
			diff[i][i] += diffX
		} else {
			diff[i][i] += diffX
		}
		if diffY != 0 {
			diff[n-i-1][i] += diffY
		} else {
			diff[n-i-1][i] += diffY
		}
	}
	return diff
}

func findTarget(input [][]int32) (int, int) {
	n := len(input[0])
	var biggest int32
	var x, y int
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if absolute(input[r][c]) > absolute(biggest) {
				biggest = input[r][c]
				x = r
				y = c
			}
		}
	}
	if biggest == 0 {
		return -1, -1
	}
	return x, y
}

func absolute(x int32) int32 {
	return int32(math.Abs(float64(x)))
}
