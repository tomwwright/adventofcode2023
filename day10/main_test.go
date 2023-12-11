package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

const input = `
7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`

func TestParseGrid(t *testing.T) {
	grid := parseGrid(input)
	expected := [][]rune{
		{'7', '-', 'F', '7', '-'},
		{'.', 'F', 'J', '|', '7'},
		{'S', 'J', 'L', 'L', '7'},
		{'|', 'F', '-', '-', 'J'},
		{'L', 'J', '.', 'L', 'J'},
	}

	for i, row := range expected {
		if !slices.Equal(grid[i], row) {
			t.Error("test failed", grid[i], row)
		}
	}
}

func TestFindStart(t *testing.T) {
	grid := parseGrid(input)
	row, column := findStart(grid)
	expectedRow, expectedColumn := 2, 0

	if row != expectedRow {
		t.Error("wrong row", row, expectedRow)
	}

	if column != expectedColumn {
		t.Error("wrong row", column, expectedColumn)
	}
}

func TestFindNext(t *testing.T) {
	grid := parseGrid(input)
	testCases := []struct {
		row        int
		column     int
		last       string
		nextRow    int
		nextColumn int
		nextLast   string
	}{
		{
			2, 0, "none", 3, 0, "north", // S, south
		},
		{
			3, 0, "north", 4, 0, "north", // |, south
		},
		{
			4, 0, "north", 4, 1, "west", // L, east
		},
		{
			4, 1, "west", 3, 1, "south", // J, north
		},
		{
			3, 1, "south", 3, 2, "west", // F, east
		},
		{
			3, 2, "west", 3, 3, "west", // -, east
		},
		{
			3, 4, "west", 2, 4, "south", // J, north
		},
		{
			2, 4, "south", 2, 3, "east", // 7, west
		},
	}
	for _, tc := range testCases {
		row, column, last := findNext(grid, tc.row, tc.column, tc.last)
		if row != tc.nextRow || column != tc.nextColumn || last != tc.nextLast {
			t.Error("test failed", row, column, last, tc)
		}
	}
}
