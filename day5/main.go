package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("./input.txt")

	readMapping(string(content))
}

type Mapping struct {
	source      int
	destination int
	length      int
}

type Transform = [][]Mapping

func readMapping(content string) {
	parts := []string{
		`seeds:((?:\s\d+)+)\n\n`,
		`seed-to-soil map:\n((?:\d+\s)+)\n`,
		`soil-to-fertilizer map:\n((?:\d+\s)+)\n`,
		`fertilizer-to-water map:\n((?:\d+\s)+)\n`,
		`water-to-light map:\n((?:\d+\s)+)\n`,
		`light-to-temperature map:\n((?:\d+\s)+)\n`,
		`temperature-to-humidity map:\n((?:\d+\s)+)\n`,
		`humidity-to-location map:\n((?:\d+\s)+)`,
	}
	r := regexp.MustCompile(strings.Join(parts, ""))
	matches := r.FindStringSubmatch(content)
	t := [][]Mapping{}
	for _, str := range matches[2:] {
		mappings := parseMappings(str)
		t = append(t, mappings)
	}

	seeds := parseSeeds(matches[1])
	fmt.Println("Part 1: Seeds", len(seeds))
	fmt.Println("Min: ", findMinimumLocation(seeds, t))

	seedsFromRanges := parseSeedRanges(matches[1])
	fmt.Println("Part 2: Seeds", len(seedsFromRanges))
	fmt.Println("Min: ", findMinimumLocation(seedsFromRanges, t))
}

func findMinimumLocation(seeds []int, t Transform) int {
	min := 100000000000000
	for _, seed := range seeds {
		translated := transform(seed, t)
		if translated < min {
			min = translated
		}
	}
	return min
}

func transform(n int, transform Transform) int {
	for _, mappings := range transform {
		n = translate(n, mappings)
	}
	return n
}

func translate(n int, mappings []Mapping) int {
	translated := n
	for _, mapping := range mappings {
		min := mapping.destination
		max := mapping.destination + mapping.length - 1
		if n >= min && n <= max {
			translated = mapping.source + (n - mapping.destination)
		}
	}
	return translated
}

func parseSeedRanges(seedsStr string) []int {
	seeds := []int{}
	valueStrs := strings.Split(strings.TrimSpace(seedsStr), " ")
	for len(valueStrs) > 0 {
		start, _ := strconv.Atoi(valueStrs[0])
		length, _ := strconv.Atoi(valueStrs[1])
		for i := start; i < start+length; i++ {
			seeds = append(seeds, i)
		}
		valueStrs = valueStrs[2:]
	}
	return seeds
}

func parseSeeds(seedsStr string) []int {
	seeds := []int{}
	for _, seedStr := range strings.Split(strings.TrimSpace(seedsStr), " ") {
		seed, _ := strconv.Atoi(seedStr)
		seeds = append(seeds, seed)
	}
	return seeds
}

func parseMappings(str string) []Mapping {
	mappings := []Mapping{}
	for _, row := range strings.Split(strings.TrimSpace(str), "\n") {
		split := strings.Split(strings.TrimSpace(row), " ")
		source, _ := strconv.Atoi(split[0])
		destination, _ := strconv.Atoi(split[1])
		length, _ := strconv.Atoi(split[2])

		mappings = append(mappings, Mapping{
			source,
			destination,
			length,
		})
	}
	return mappings
}
