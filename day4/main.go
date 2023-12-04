package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(content), "\n")

	fmt.Println("Sum", part1(lines))
	fmt.Println("Copies", part2(lines))
}

func part1(lines []string) int {
	sum := 0
	for i, line := range lines {
		winners, have := parseCard(line)

		score := scoreCard(winners, have)
		if score > 1 {
			score = int(math.Pow(float64(2), float64(score-1)))
		}
		fmt.Println(i+1, "Score", score)

		sum += score
	}
	return sum
}

func part2(lines []string) int {

	winsOnEachCard := []int{}
	copiesOfEachCard := []int{}

	for _, line := range lines {
		winners, have := parseCard(line)
		wins := scoreCard(winners, have)
		winsOnEachCard = append(winsOnEachCard, wins)
		copiesOfEachCard = append(copiesOfEachCard, 1) // start with one copy of each
	}

	max := len(lines)
	for i, wins := range winsOnEachCard {
		for j := 1; j <= wins; j++ {
			if i+j >= max {
				break
			}
			copiesOfEachCard[i+j] += copiesOfEachCard[i]
		}
	}

	sum := 0
	for _, copies := range copiesOfEachCard {
		sum += copies
	}
	return sum
}

func scoreCard(winners []int, have []int) int {
	mask := createMask(winners)
	matches := 0
	for _, n := range have {
		if mask[n] {
			matches++
		}
	}
	return matches
}

func createMask(numbers []int) [100]bool {
	mask := [100]bool{}
	for _, n := range numbers {
		mask[n] = true
	}
	return mask
}

func parseCard(card string) ([]int, []int) {
	r := regexp.MustCompile(`Card\s+\d+:((?:\s+\d+)+) \|((?:\s+\d+)+)`)
	match := r.FindStringSubmatch(card)

	winners := parseNumbers(match[1])
	have := parseNumbers(match[2])

	return winners, have
}

func parseNumbers(str string) []int {
	numbers := []int{}
	split := strings.Split(str, " ")
	for _, number := range split {
		number = strings.TrimSpace(number)
		if number != "" {
			parsed, _ := strconv.Atoi(number)
			numbers = append(numbers, parsed)
		}
	}
	return numbers
}
