package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, _ := os.ReadFile("./input.txt")

	part1(string(content))
	part2(string(content))

}

func part1(content string) {
	grid := parseGrid(content)
	path := parseLoop(grid)
	fmt.Println("Part 1", len(path)/2)
}

func part2(content string) {
	fmt.Println("Part 2", "???")
}

func parseLoop(grid [][]rune) [][]int {
	startRow, startColumn := findStart(grid)

	row, column, last := startRow, startColumn, "none"
	path := [][]int{{startRow, startColumn}}

	for last == "none" || row != startRow || column != startColumn {
		row, column, last = findNext(grid, row, column, last)
		path = append(path, []int{row, column})
	}

	return path
}

func findNext(grid [][]rune, row int, column int, last string) (int, int, string) {
	current := grid[row][column]

	north := func() (bool, int, int) {
		if row > 0 && last != "north" {
			ch := grid[row-1][column]
			if ch == '|' || ch == 'F' || ch == '7' || ch == 'S' {
				return true, row - 1, column
			}
		}
		return false, -1, -1
	}

	south := func() (bool, int, int) {
		if row < len(grid)-1 && last != "south" {
			ch := grid[row+1][column]
			if ch == '|' || ch == 'J' || ch == 'L' || ch == 'S' {
				return true, row + 1, column
			}
		}
		return false, -1, -1
	}

	west := func() (bool, int, int) {
		if column > 0 && last != "west" {
			ch := grid[row][column-1]
			if ch == '-' || ch == 'L' || ch == 'F' || ch == 'S' {
				return true, row, column - 1
			}
		}
		return false, -1, -1
	}

	east := func() (bool, int, int) {
		if column < len(grid[0])-1 && last != "east" {
			ch := grid[row][column+1]
			if ch == '-' || ch == 'J' || ch == '7' || ch == 'S' {
				return true, row, column + 1
			}
		}
		return false, -1, -1
	}

	found := false
	nextRow := -1
	nextColumn := -1

	// determine list of directions to check based on the current cell
	directionChecks := []func() (bool, int, int){}
	switch current {
	case 'S':
		directionChecks = []func() (bool, int, int){north, south, west, east}
	case '-':
		directionChecks = []func() (bool, int, int){west, east}
	case '|':
		directionChecks = []func() (bool, int, int){north, south}
	case 'J':
		directionChecks = []func() (bool, int, int){west, north}
	case '7':
		directionChecks = []func() (bool, int, int){west, south}
	case 'L':
		directionChecks = []func() (bool, int, int){east, north}
	case 'F':
		directionChecks = []func() (bool, int, int){east, south}
	}

	// check each of the valid directions and return the first match
	for _, direction := range directionChecks {
		found, nextRow, nextColumn = direction()
		if found {
			last := findLastDirection(row, column, nextRow, nextColumn)
			return nextRow, nextColumn, last
		}
	}

	return -1, -1, "unknown"
}

func findLastDirection(row int, column int, nextRow int, nextColumn int) string {
	if row == nextRow && column == nextColumn-1 {
		return "west"
	}

	if row == nextRow && column == nextColumn+1 {
		return "east"
	}

	if row == nextRow-1 && column == nextColumn {
		return "north"
	}

	if row == nextRow+1 && column == nextColumn {
		return "south"
	}

	return "unknown"
}

func findStart(grid [][]rune) (int, int) {
	for i, row := range grid {
		for j, ch := range row {
			if ch == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

func parseGrid(content string) [][]rune {
	rows := strings.Split(strings.TrimSpace(content), "\n")
	grid := make([][]rune, len(rows))
	for i, row := range rows {
		grid[i] = []rune(row)
	}
	return grid
}
