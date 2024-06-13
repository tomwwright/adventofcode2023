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

	part1(string(content))
	part2(string(content))

}

func part1(content string) {
	records := parseInput(content)
	fmt.Println("Part 1", records)
}

func part2(content string) {

	fmt.Println("Part 2", "???")
}

type Record struct {
	groups []int
	data   string
}

func parseInput(content string) []Record {
	records := []Record{}
	r := regexp.MustCompile(`([\.#\?]+) ((?:\d,)*\d)`)
	for _, line := range strings.Split(content, "\n") {
		match := r.FindStringSubmatch(line)

		data := match[1]
		groups := []int{}
		for _, str := range strings.Split(match[2], ",") {
			n, _ := strconv.Atoi(str)
			groups = append(groups, n)
		}

		records = append(records, Record{
			groups,
			data,
		})
	}
	return records
}
