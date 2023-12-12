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

	fmt.Println("Part 2", "???")
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
