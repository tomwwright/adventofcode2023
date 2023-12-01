package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	sum := 0
	for i, line := range lines {
		num := ParseCalibrationValueWithWords(line)
		sum += num
		fmt.Println(i, num, line)
	}
	fmt.Println(sum)
}

func ParseCalibrationValueWithWords(s string) int {
	matchers := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	first := 0
	firstFound := len(s)
	last := 0
	lastFound := -1
	for k, v := range matchers {
		index := strings.Index(s, k)
		lastIndex := strings.LastIndex(s, k)
		if index != -1 && index < firstFound {
			first = v
			firstFound = index
		}
		if lastIndex != -1 && lastIndex > lastFound {
			last = v
			lastFound = lastIndex
		}
	}
	return first*10 + last
}

func ParseCalibrationValue(s string) int {
	var first int
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			num, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}
			first = num
			break
		}
	}

	var last int
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			num, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}
			last = num
			break
		}
	}

	return first*10 + last
}
