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
	path := parsePath(grid)
	fmt.Println("Part 1", len(path)/2)
}

func part2(content string) {

	grid := parseGrid(content)
	largeGrid := constructHighResolutionGrid(grid)

	floodfillGrid(largeGrid)

	count := countEnclosed(largeGrid)
	fmt.Println("Part 2", count, "tiles")
}

func countEnclosed(grid [][]rune) int {
	count := 0
	for i := 0; i < len(grid); i += 3 {
		for j := 0; j < len(grid[i]); j += 3 {
			// combine the cells in this local 3x3
			cells := string([]rune{
				grid[i][j],
				grid[i][j+1],
				grid[i][j+2],
				grid[i+1][j],
				grid[i+1][j+1],
				grid[i+1][j+2],
				grid[i+2][j],
				grid[i+2][j+1],
				grid[i+2][j+2],
			})
			if cells == "........." {
				count++
			}
		}
	}
	return count
}

func floodfillGrid(grid [][]rune) {
	blocked := 'x'
	filled := 'O'
	startRow := 0
	startColumn := 0
	maxRow := len(grid) - 1
	maxColumn := len(grid[0]) - 1

	stack := [][2]int{{startRow, startColumn}}
	pop := func() (int, int) {
		elem := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop from the end because memory
		return elem[0], elem[1]
	}
	push := func(coords [2]int) {
		stack = append(stack, coords)
	}

	for len(stack) > 0 {
		row, column := pop()
		grid[row][column] = filled

		directions := [4][2]int{
			{row - 1, column},
			{row + 1, column},
			{row, column - 1},
			{row, column + 1},
		}

		for _, next := range directions {
			nextRow := next[0]
			nextColumn := next[1]
			if nextRow < 0 || nextRow > maxRow || nextColumn < 0 || nextColumn > maxColumn {
				continue
			}
			nextCh := grid[nextRow][nextColumn]
			if nextCh == blocked || nextCh == filled {
				continue
			}
			push(next)
		}
	}

}

func constructHighResolutionGrid(grid [][]rune) [][]rune {
	largeGrid := make([][]rune, len(grid)*3)
	for i := range largeGrid {
		largeGrid[i] = make([]rune, len(grid[0])*3)
		for j := range largeGrid[i] {
			largeGrid[i][j] = '.'
		}
	}

	applyPatch := func(patch [][]rune, row int, column int) {
		for i, patchRow := range patch {
			for j, ch := range patchRow {
				largeGrid[row*3+i][column*3+j] = ch
			}
		}
	}

	path := parsePath(grid)

	for _, coords := range path {
		row := coords[0]
		column := coords[1]
		ch := grid[row][column]

		if ch == '-' {
			applyPatch([][]rune{
				{'.', '.', '.'},
				{'x', 'x', 'x'},
				{'.', '.', '.'},
			}, row, column)
		}

		if ch == '|' {
			applyPatch([][]rune{
				{'.', 'x', '.'},
				{'.', 'x', '.'},
				{'.', 'x', '.'},
			}, row, column)
		}

		if ch == 'F' {
			applyPatch([][]rune{
				{'.', '.', '.'},
				{'.', 'x', 'x'},
				{'.', 'x', '.'},
			}, row, column)
		}

		if ch == '7' {
			applyPatch([][]rune{
				{'.', '.', '.'},
				{'x', 'x', '.'},
				{'.', 'x', '.'},
			}, row, column)
		}

		if ch == 'J' {
			applyPatch([][]rune{
				{'.', 'x', '.'},
				{'x', 'x', '.'},
				{'.', '.', '.'},
			}, row, column)
		}

		if ch == 'L' {
			applyPatch([][]rune{
				{'.', 'x', '.'},
				{'.', 'x', 'x'},
				{'.', '.', '.'},
			}, row, column)
		}

		if ch == 'S' {
			/*
			 * Cheating here because we can just look at the input
			 * to determine what "shape" this should be in 3x3
			 *
			 * To do this dynamically:
			 * - The adjacent points to S are second and second-last elements of path
			 * - use findCardinalDirection to get direction from S
			 * - apply a patch based on the two cardinal directions
			 */
			applyPatch([][]rune{
				{'.', 'x', '.'},
				{'x', 'x', '.'},
				{'.', '.', '.'},
			}, row, column)
		}
	}

	return largeGrid
}

func parsePath(grid [][]rune) [][]int {
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
			last := findCardinalDirection(row, column, nextRow, nextColumn)
			return nextRow, nextColumn, last
		}
	}

	return -1, -1, "unknown"
}

func findCardinalDirection(toRow int, toColumn int, fromRow int, fromColumn int) string {
	if toRow == fromRow && toColumn < fromColumn {
		return "west"
	}

	if toRow == fromRow && toColumn > fromColumn {
		return "east"
	}

	if toRow < fromRow && toColumn == fromColumn {
		return "north"
	}

	if toRow > fromRow && toColumn == fromColumn {
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
