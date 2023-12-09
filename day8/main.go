package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	content, _ := os.ReadFile("./input.txt")

	fmt.Println("Part 1", string(content))

	directions, lefts, rights := parseInput(string(content))

	part1(directions, lefts, rights)
	part2(directions, lefts, rights)
}

func parseInput(content string) (string, map[string]string, map[string]string) {
	rows := strings.Split(content, "\n")

	directions := rows[0]
	lefts, rights := parseNodes(rows[2:])

	return directions, lefts, rights
}

func part1(directions string, lefts map[string]string, rights map[string]string) {
	steps := calculateStepsToTerminal(directions, "AAA", lefts, rights)
	fmt.Println("Part 1 Steps", steps)
}

func part2(directions string, lefts map[string]string, rights map[string]string) {

	nodes := findStartNodes(lefts)

	fmt.Println("Start Nodes", nodes)

	stepsPerNode := []int64{}
	for _, node := range nodes {
		stepsPerNode = append(stepsPerNode, int64(calculateStepsToTerminal(directions, node, lefts, rights)))
	}

	steps := lcm(stepsPerNode...)

	fmt.Println("Part 2 Steps per node", stepsPerNode)
	fmt.Println("Part 2 Steps", steps)
}

func calculateStepsToTerminal(directions string, node string, lefts map[string]string, rights map[string]string) int {
	steps := 0
	for !isEndState([]string{node}) {
		for _, ch := range directions {
			if ch == 'L' {
				node = lefts[node]
			} else if ch == 'R' {
				node = rights[node]
			}

			steps++

			if node == "ZZZ" {
				break
			}
		}
	}
	return steps
}

func isEndState(nodes []string) bool {
	for _, node := range nodes {
		if node[len(node)-1] != 'Z' {
			return false
		}
	}
	return true
}

func findStartNodes(nodeMap map[string]string) []string {
	starts := []string{}
	for node := range nodeMap {
		if node[len(node)-1] == 'A' {
			starts = append(starts, node)
		}
	}
	return starts
}

func parseNodes(rows []string) (map[string]string, map[string]string) {
	r := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

	lefts := map[string]string{}
	rights := map[string]string{}

	for _, row := range rows {
		match := r.FindStringSubmatch(row)

		node := match[1]
		left := match[2]
		right := match[3]

		lefts[node] = left
		rights[node] = right
	}

	return lefts, rights
}
