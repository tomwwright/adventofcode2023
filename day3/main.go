package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const size = 140

type Grid = [size][size]rune

func main() {
	content, _ := os.ReadFile("./input.txt")
	grid := readGrid(string(content))

	part1(grid)
	part2(grid)
}

func part1(grid Grid) {
	symbols := findSymbols(grid)
	partNumberSum := 0
	for _, coords := range symbols {
		partNumbers := findAdjacentPartNumbers(coords[0], coords[1], grid)
		for _, partNumber := range partNumbers {
			partNumberSum += partNumber
		}
	}

	fmt.Println("Part Number Sum", partNumberSum)
}

func part2(grid Grid) {
	symbols := findSymbols(grid)
	gearRatioSum := 0
	for _, coords := range symbols {
		ch := grid[coords[0]][coords[1]]
		if ch == '*' {
			partNumbers := findAdjacentPartNumbers(coords[0], coords[1], grid)
			if len(partNumbers) == 2 {
				gearRatioSum += partNumbers[0] * partNumbers[1]
			}
		}
	}

	fmt.Println("Gear Ratio Sum", gearRatioSum)
}

func readGrid(content string) Grid {
	rows := strings.Split(content, "\n")

	grid := Grid{}
	for i, row := range rows {
		for j, cell := range row {
			grid[i][j] = cell
		}
	}
	return grid
}

func findSymbols(grid Grid) [][2]int {
	symbols := [][2]int{}
	for i, row := range grid {
		for j, ch := range row {
			if ch != '.' && ch != 0 && (ch < '0' || ch > '9') {
				symbols = append(symbols, [2]int{i, j})
			}
		}
	}
	return symbols
}

func findAdjacentPartNumbers(i int, j int, grid Grid) []int {
	searches := [][2]int{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j - 1},
		{i, j + 1},
		{i + 1, j - 1},
		{i + 1, j},
		{i + 1, j + 1},
	}

	partNumbers := []int{}
	for _, search := range searches {
		if search[0] >= 0 && search[0] < size && search[1] > 0 && search[1] < size {
			ch := grid[search[0]][search[1]]
			if ch >= '0' && ch <= '9' {
				partNumber, _ := parsePartNumber(search[0], search[1], grid)
				if !slices.Contains(partNumbers, partNumber) {
					partNumbers = append(partNumbers, partNumber)
				}
			}
		}
	}
	return partNumbers
}

func parsePartNumber(i int, j int, grid Grid) (int, error) {
	foundStart := false
	for !foundStart {
		ch := grid[i][j]
		j--
		if ch < '0' || ch > '9' {
			foundStart = true
			j++
		} else if j < 0 {
			foundStart = true
		}
	}

	partNumberString := ""
	foundEnd := false
	for !foundEnd {
		j++
		if j >= size {
			foundEnd = true
		} else {
			ch := grid[i][j]
			if ch >= '0' && ch <= '9' {
				partNumberString += string(ch)
			} else {
				foundEnd = true
			}
		}
	}

	if partNumberString == "" {
		return 0, nil
	}
	return strconv.Atoi(partNumberString)
}

func DisplayGrid(grid [140][140]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}
