package main

import (
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

var grid = readGrid(strings.TrimSpace(`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......7551
...$.*....
.664.598..`))

func TestParsePartNumber(t *testing.T) {
	testCases := []struct {
		coords   [2]int
		expected int
	}{
		{
			[2]int{0, 2},
			467,
		},
		{
			[2]int{1, 2},
			0,
		},
		{
			[2]int{1, 3},
			0,
		},
		{
			[2]int{0, 5},
			114,
		},
		{
			[2]int{2, 7},
			633,
		},
		{
			[2]int{7, 8},
			7551,
		},
	}
	for _, tc := range testCases {
		parsed, _ := parsePartNumber(tc.coords[0], tc.coords[1], grid)
		if parsed != tc.expected {
			t.Error("test failed", parsed, tc)
		}
	}
}

func TestFindAdjacentPartNumbers(t *testing.T) {
	testCases := []struct {
		coords   [2]int
		expected []int
	}{
		{
			[2]int{1, 3},
			[]int{467, 35},
		},
		{
			[2]int{8, 3},
			[]int{664},
		},
		{
			[2]int{5, 5},
			[]int{592},
		},
	}
	for _, tc := range testCases {
		parsed := findAdjacentPartNumbers(tc.coords[0], tc.coords[1], grid)
		if !slices.Equal(parsed, tc.expected) {
			t.Error("test failed", parsed, tc)
		}
	}
}

func TestFindSymbols(t *testing.T) {
	symbols := findSymbols(grid)
	if len(symbols) != 6 {
		t.Error("test failed", symbols)
	}
}
