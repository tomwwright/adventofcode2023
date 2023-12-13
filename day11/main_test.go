package main

import (
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestFindHashes(t *testing.T) {
	input := strings.TrimSpace(`
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`)

	points := findHashes(input)
	expected := []Point{
		{0, 3},
		{1, 7},
		{2, 0},
		{4, 6},
		{5, 1},
		{6, 9},
		{8, 7},
		{9, 0},
		{9, 4},
	}

	if !slices.Equal(points, expected) {
		t.Error("testfailed", "\n", points, "\n", expected)
	}
}

func TestMarkSpace(t *testing.T) {
	input := strings.TrimSpace(`
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`)

	expected := strings.TrimSpace(`
..x#.x..x.
..x..x.#x.
#.x..x..x.
xxxxxxxxxx
..x..x#.x.
.#x..x..x.
..x..x..x#
xxxxxxxxxx
..x..x.#x.
#.x.#x..x.
`)

	marked := markSpaces(input)

	if marked != expected {
		t.Error("testfailed", "\n", marked, "\n", expected)
	}
}

func TestExpandSpace(t *testing.T) {
	input := strings.TrimSpace(`
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`)

	expected := strings.TrimSpace(`
....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......
`)

	expanded := expandSpace(input)

	if expanded != expected {
		t.Error("testfailed", "\n", expanded, "\n", expected)
	}
}

func TestTranspose(t *testing.T) {
	input := strings.TrimSpace(`
123
456
789
`)

	expected := strings.TrimSpace(`
147
258
369
`)

	transposed := strings.Join(transpose(strings.Split(input, "\n")), "\n")

	if transposed != expected {
		t.Error("test failed", "\n", transposed, "\n", expected)
	}
}

func TestCombinations(t *testing.T) {
	input := []Point{
		{0, 3},
		{1, 7},
		{2, 0},
	}

	combinations := combinations(input)

	expected := [][2]Point{
		{
			{0, 3},
			{1, 7},
		},
		{
			{0, 3},
			{2, 0},
		},
		{
			{1, 7},
			{2, 0},
		},
	}

	if !slices.Equal(combinations, expected) {
		t.Error("testfailed", "\n", combinations, "\n", expected)
	}
}
