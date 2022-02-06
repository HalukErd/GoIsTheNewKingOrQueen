package main

import (
	"fmt"
	"strings"
)

func bomberMan(n int32, grid []string) []string {
	if n == int32(1) {
		return grid
	} else if n%2 == 0 {
		grid = fill(grid, 'O', 'O')
		return grid
	} else {
		firstExp := bomb(grid)
		secondExp := bomb(firstExp)
		if n/2%2 == 1 {
			return firstExp
		} else {
			return secondExp
		}
	}
}

func bomb(grid []string) []string {
	dot := '.'
	o := 'O'
	return fill(grid, dot, o)
}

func fill(grid []string, dot rune, o rune) []string {
	var sb strings.Builder
	var ss []string
	for rowI, row := range grid {
		for colI, ch := range row {
			if ch == o || checkNearBomb(rowI, colI, grid) {
				sb.WriteRune(dot)
			} else {
				sb.WriteRune(o)
			}
		}
		ss = append(ss, sb.String())
		sb.Reset()
	}
	return ss
}

func checkNearBomb(rowI int, colI int, grid []string) bool {
	o := uint8('O')
	if (colI-1 >= 0 && grid[rowI][colI-1] == o) ||
		(colI+1 < len(grid[0]) && grid[rowI][colI+1] == o) ||
		(rowI-1 >= 0 && grid[rowI-1][colI] == o) ||
		(rowI+1 < len(grid) && grid[rowI+1][colI] == o) {
		return true
	}
	return false
}

func main() {
	//grid := []string{".......", "...O...", "....O..", ".......", "OO.....", "OO....."}
	grid := []string{".......",
		"...O.O.",
		"....O..",
		"..O....",
		"OO...OO",
		"OO.O..."}
	result := bomberMan(1, grid)
	for _, ln := range result {
		fmt.Println(ln)
	}
	fmt.Println()
	result = bomberMan(3, grid)
	for _, ln := range result {
		fmt.Println(ln)
	}

	fmt.Println()
	result = bomberMan(5, grid)
	for _, ln := range result {
		fmt.Println(ln)
	}
	fmt.Println()
	result = bomberMan(7, grid)
	for _, ln := range result {
		fmt.Println(ln)
	}
	fmt.Println()
	result = bomberMan(9, grid)
	for _, ln := range result {
		fmt.Println(ln)
	}
	fmt.Println()
}
