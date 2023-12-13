package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	content, _ := os.ReadFile("./input.txt")

	part1(string(content))
	part2(string(content))

}

func part1(content string) {
	space := expandSpace(content)
	points := findHashes(space)
	pairs := combinations(points)

	sum := 0
	for _, pair := range pairs {
		a := pair[0]
		b := pair[1]
		distance := int(math.Abs(float64(a.row-b.row)) + math.Abs(float64(a.column-b.column)))
		sum += distance
	}

	fmt.Println("Part 1", sum)
}

func part2(content string) {
	space := markSpaces(content)
	points := findHashes(space)
	pairs := combinations(points)

	sum := 0
	rows := strings.Split(space, "\n")
	for _, pair := range pairs {
		a := pair[0]
		b := pair[1]
		distance := calculateDistance(rows, a, b)
		sum += distance
	}

	fmt.Println("Part 2", sum)
}

type Point struct {
	row    int
	column int
}

func findHashes(content string) []Point {
	points := []Point{}
	rows := strings.Split(content, "\n")
	for i, row := range rows {
		for j, ch := range row {
			if ch == '#' {
				points = append(points, Point{i, j})
			}
		}
	}
	return points
}

func calculateDistance(space []string, a Point, b Point) int {
	path := ""
	minRow := int(math.Min(float64(a.row), float64(b.row)))
	maxRow := int(math.Max(float64(a.row), float64(b.row)))
	minColumn := int(math.Min(float64(a.column), float64(b.column)))
	maxColumn := int(math.Max(float64(a.column), float64(b.column)))

	// traverse vertically, not covering last row
	for i := minRow + 1; i <= maxRow; i++ {
		path += string(space[i][minColumn])
	}

	// traverse horizontally, including the "corner"
	for i := minColumn + 1; i <= maxColumn; i++ {
		path += string(space[maxRow][i])
	}

	crosses := strings.Count(path, "x")
	dots := len(path) - crosses

	return dots + crosses*1000000
}

func expandSpace(content string) string {

	// expand rows
	rows := strings.Split(content, "\n")
	expandedRows := []string{}
	for _, row := range rows {
		expandedRows = append(expandedRows, row)
		if strings.Count(row, "#") == 0 {
			expandedRows = append(expandedRows, row)
		}
	}

	columns := transpose(expandedRows)

	// expand columns
	expandedColumns := []string{}
	for _, column := range columns {
		expandedColumns = append(expandedColumns, column)
		if strings.Count(column, "#") == 0 {
			expandedColumns = append(expandedColumns, column)
		}
	}

	rows = transpose(expandedColumns)
	return strings.Join(rows, "\n")
}

func markSpaces(content string) string {

	// replace empty rows with x
	rows := strings.Split(content, "\n")
	for i, row := range rows {
		if strings.Count(row, "#") == 0 {
			rows[i] = strings.ReplaceAll(row, ".", "x")
		}
	}

	columns := transpose(rows)

	// replace empty columns with x
	for i, column := range columns {
		if strings.Count(column, "#") == 0 {
			columns[i] = strings.ReplaceAll(column, ".", "x")
		}
	}

	rows = transpose(columns)
	return strings.Join(rows, "\n")
}

func transpose(rows []string) []string {
	columns := []string{}
	for i := range rows[0] {
		column := ""
		for _, row := range rows {
			column += string(row[i])
		}

		columns = append(columns, column)
	}
	return columns
}

func combinations(points []Point) [][2]Point {
	combinations := [][2]Point{}
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]
			combinations = append(combinations, [2]Point{a, b})
		}
	}
	return combinations
}
