package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	content, _ := os.ReadFile("./input.txt")

	part1(string(content))
	part2(string(content))

}

func part1(content string) {
	sum := 0
	for _, row := range strings.Split(content, "\n") {
		sequence := parseSequence(row)
		sum += extrapolateSequence(sequence)
	}
	fmt.Println("Part 1", sum)
}

func part2(content string) {
	sum := 0
	for _, row := range strings.Split(content, "\n") {
		sequence := parseSequence(row)
		sum += extrapolateSequenceBackwards(sequence)
	}
	fmt.Println("Part 2", sum)
}

func parseSequence(content string) []int {
	elements := strings.Split(content, " ")
	numbers := []int{}
	for _, e := range elements {
		n, _ := strconv.Atoi(e)
		numbers = append(numbers, n)
	}
	return numbers
}

func extrapolateSequenceBackwards(sequence []int) int {
	slices.Reverse(sequence)
	return extrapolateSequence(sequence)
}

func extrapolateSequence(sequence []int) int {
	differences := calculateDifferences(sequence)
	if isZeros(differences) {
		return sequence[len(sequence)-1]
	}
	next := extrapolateSequence(differences)
	return sequence[len(sequence)-1] + next
}

func calculateDifferences(sequence []int) []int {
	differences := []int{}
	for i := 1; i < len(sequence); i++ {
		a := sequence[i-1]
		b := sequence[i]
		differences = append(differences, b-a)
	}
	return differences
}

func isZeros(ns []int) bool {
	for _, n := range ns {
		if n != 0 {
			return false
		}
	}
	return true
}
