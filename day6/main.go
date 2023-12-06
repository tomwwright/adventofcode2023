package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
 * distance travelled is remaining time (after charging) by speed
 * 	distance = (max_t - charge_t) * speed
 * each millisecond charged adds one millisecond of speed
 *	speed = charge_t
 * i.e.
 *  distance = (max_t - charge_t) * charge_t
 *
 * we are interested in beating record, that is
 *	distance > max_distance
 * i.e.
 *  (max_t - charge_t) * charge_t > max_distance
 */

func main() {
	content, _ := os.ReadFile("./input.txt")

	times, distances := parseRaces(string(content))

	part1(times, distances)
	part2(times, distances)
}

func parseRaces(input string) ([]int, []int) {
	pattern := strings.Join(
		[]string{
			`Time:((?:\s+\d+)+)`,
			`Distance:((?:\s+\d+)+)`,
		},
		"\n")
	r := regexp.MustCompile(pattern)
	match := r.FindStringSubmatch(input)

	times := parseNumbers(match[1])
	distances := parseNumbers(match[2])

	return times, distances
}

func part1(times []int, distances []int) {
	fmt.Println("Part 1")
	output := 1
	for i, time := range times {
		distance := distances[i]

		solutions := calculateNumSolutions(time, distance)
		output *= solutions

		fmt.Println("t", time, "d", distance, "solutions", solutions)
	}
	fmt.Println("Output", output)
}

func part2(times []int, distances []int) {
	time := combineNumbers(times)
	distance := combineNumbers(distances)

	solutions := calculateNumSolutions(time, distance)

	fmt.Println("Part 2")
	fmt.Println("t", time, "d", distance, "solutions", solutions)
}

func combineNumbers(numbers []int) int {
	str := ""
	for _, n := range numbers {
		str += strconv.Itoa(n)
	}
	num, _ := strconv.Atoi(str)
	return num
}

func calculateNumSolutions(time int, distance int) int {
	solutions := 0
	for t := 1; t < time; t++ {
		// distance = (max_t - charge_t) * charge_t
		d := (time - t) * t
		if d > distance {
			solutions++
		}
	}
	return solutions
}

func parseNumbers(input string) []int {
	splits := strings.Split(strings.TrimSpace(input), " ")
	numbers := []int{}
	for _, split := range splits {
		if split != "" {
			n, _ := strconv.Atoi(split)
			numbers = append(numbers, n)
		}
	}
	return numbers
}
