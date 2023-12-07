package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

var ranks = map[string]int{
	"fiveofakind":  6,
	"fourofakind":  5,
	"fullhouse":    4,
	"threeofakind": 3,
	"twopair":      2,
	"onepair":      1,
	"highcard":     0,
}

type Hand struct {
	hand string
	rank int
	bid  int
}

func main() {
	content, _ := os.ReadFile("./input.txt")

	part1 := part1(string(content))
	fmt.Println("Part 1", part1)

	part2 := part2(string(content))
	fmt.Println("Part 2", part2)
}

func part1(content string) int {
	hands := parseInput(content)

	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}

	return sum
}

func part2(content string) int {
	hands := parseInputWithJokers(content)

	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}

	return sum
}

func parseInput(content string) []Hand {
	hands := []Hand{}
	for _, row := range strings.Split(strings.TrimSpace(content), "\n") {
		split := strings.Split(row, " ")

		hand := split[0]
		rank := parseRank(split[0])
		bid, _ := strconv.Atoi(split[1])

		hands = append(hands, Hand{
			hand,
			rank,
			bid,
		})
	}

	slices.SortFunc(hands, sortHands)

	return hands
}

func parseInputWithJokers(content string) []Hand {
	hands := []Hand{}
	for _, row := range strings.Split(strings.TrimSpace(content), "\n") {
		split := strings.Split(row, " ")

		hand := split[0]
		rank := parseRankWithJokers(split[0])
		bid, _ := strconv.Atoi(split[1])

		hands = append(hands, Hand{
			hand,
			rank,
			bid,
		})
	}

	slices.SortFunc(hands, sortHandsWithJokers)

	return hands
}

var values = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func sortHands(a Hand, b Hand) int {

	// use 100s to sort if different type
	if a.rank-b.rank != 0 {
		return a.rank - b.rank
	}

	// otherwise do bullshit comparison based on order
	for i, chA := range a.hand {
		valueA := values[chA]
		valueB := values[rune(b.hand[i])]

		if valueA-valueB != 0 {
			return valueA - valueB
		}
	}

	return 0
}

var valuesWithJokers = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

func sortHandsWithJokers(a Hand, b Hand) int {

	// use 100s to sort if different type
	if a.rank-b.rank != 0 {
		return a.rank - b.rank
	}

	// otherwise do bullshit comparison based on order
	for i, chA := range a.hand {
		valueA := valuesWithJokers[chA]
		valueB := valuesWithJokers[rune(b.hand[i])]

		if valueA-valueB != 0 {
			return valueA - valueB
		}
	}

	return 0
}

func parseRankWithJokers(hand string) int {

	// store occurences by rune
	counts := map[rune]int{
		'A': 0,
		'K': 0,
		'Q': 0,
		'J': 0,
		'T': 0,
		'9': 0,
		'8': 0,
		'7': 0,
		'6': 0,
		'5': 0,
		'4': 0,
		'3': 0,
		'2': 0,
	}
	for _, ch := range hand {
		counts[ch]++
	}

	// store values by frequency of occurrences, count jokers
	frequencies := [5][]rune{}
	jokers := counts['J']
	for ch, n := range counts {
		if ch != 'J' && n > 0 {
			frequencies[n-1] = append(frequencies[n-1], ch)
		}
	}

	// is it five of a kind?
	if len(frequencies[4]) == 1 {
		return ranks["fiveofakind"]
	}

	// is it four of a kind?
	if len(frequencies[3]) == 1 {
		switch jokers {
		case 1:
			return ranks["fiveofakind"]
		default:
			return ranks["fourofakind"]
		}
	}

	// is it full house?
	if len(frequencies[2]) == 1 && len(frequencies[1]) == 1 {
		return ranks["fullhouse"]
	}

	// is it three of a kind?
	if len(frequencies[2]) == 1 {
		switch jokers {
		case 2:
			return ranks["fiveofakind"]
		case 1:
			return ranks["fourofakind"]
		default:
			return ranks["threeofakind"]
		}
	}

	// is it two pair?
	if len(frequencies[1]) == 2 {
		switch jokers {
		case 1:
			return ranks["fullhouse"]
		default:
			return ranks["twopair"]
		}
	}

	// is it one pair?
	if len(frequencies[1]) == 1 {
		switch jokers {
		case 3:
			return ranks["fiveofakind"]
		case 2:
			return ranks["fourofakind"]
		case 1:
			return ranks["threeofakind"]
		default:
			return ranks["onepair"]
		}
	}

	// must be high card
	switch jokers {
	case 5:
		return ranks["fiveofakind"]
	case 4:
		return ranks["fiveofakind"]
	case 3:
		return ranks["fourofakind"]
	case 2:
		return ranks["threeofakind"]
	case 1:
		return ranks["onepair"]
	default:
		return ranks["highcard"]
	}
}

func parseRank(hand string) int {

	// store occurences by rune
	counts := map[rune]int{
		'A': 0,
		'K': 0,
		'Q': 0,
		'J': 0,
		'T': 0,
		'9': 0,
		'8': 0,
		'7': 0,
		'6': 0,
		'5': 0,
		'4': 0,
		'3': 0,
		'2': 0,
	}
	for _, ch := range hand {
		counts[ch]++
	}

	// store values by frequency of occurrences
	frequencies := [5][]rune{}
	for ch, n := range counts {
		if n > 0 {
			frequencies[n-1] = append(frequencies[n-1], ch)
		}
	}

	// is it five of a kind?
	if len(frequencies[4]) == 1 {
		return ranks["fiveofakind"]
	}

	// is it four of a kind?
	if len(frequencies[3]) == 1 {
		return ranks["fourofakind"]
	}

	// is it full house?
	if len(frequencies[2]) == 1 && len(frequencies[1]) == 1 {
		return ranks["fullhouse"]
	}

	// is it three of a kind?
	if len(frequencies[2]) == 1 {
		return ranks["threeofakind"]
	}

	// is it two pair?
	if len(frequencies[1]) == 2 {
		return ranks["twopair"]
	}

	// is it one pair?
	if len(frequencies[1]) == 1 {
		return ranks["onepair"]
	}

	// must be high card
	return ranks["highcard"]
}
