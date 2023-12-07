package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

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

var ranks = map[string]int{
	"fiveofakind":  600,
	"fourofakind":  500,
	"fullhouse":    400,
	"threeofakind": 300,
	"twopair":      200,
	"onepair":      100,
	"highcard":     0,
}

func main() {
	content, _ := os.ReadFile("./input.txt")

	hands := parseInput(string(content))

	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}

	fmt.Println("Sum", sum)
}

type Hand struct {
	hand   string
	values []int
	bid    int
}

func parseInput(content string) []Hand {
	hands := []Hand{}
	for _, row := range strings.Split(strings.TrimSpace(content), "\n") {
		split := strings.Split(row, " ")

		hand := split[0]
		values := parseHand(split[0])
		bid, _ := strconv.Atoi(split[1])

		hands = append(hands, Hand{
			hand,
			values,
			bid,
		})
	}

	slices.SortFunc(hands, sortHands)

	return hands
}

func sortHands(a Hand, b Hand) int {
	// use 100s to sort if different type
	if (a.values[0]/100)-(b.values[0]/100) != 0 {
		return a.values[0] - b.values[0]
	}

	// otherwise do bullshit comparison based on first
	for i, chA := range a.hand {
		valueA := values[chA]
		valueB := values[rune(b.hand[i])]

		if valueA-valueB != 0 {
			return valueA - valueB
		}
	}

	return 0
}

func parseHand(hand string) []int {

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
	frequencies := [5][]int{}
	for ch, n := range counts {
		if n > 0 {

			value := values[ch]
			frequencies[n-1] = append(frequencies[n-1], value)
		}
	}

	// sort frequencies by ascending
	for i := 0; i < len(frequencies); i++ {
		slices.Sort(frequencies[i])
		slices.Reverse(frequencies[i])
	}

	// is it five of a kind?
	if len(frequencies[4]) == 1 {
		value := frequencies[4][0]
		return []int{
			ranks["fiveofakind"] + value,
		}
	}

	// is it four of a kind?
	if len(frequencies[3]) == 1 {
		value := frequencies[3][0]
		high := frequencies[0][0]
		return []int{
			ranks["fourofakind"] + value,
			high,
		}
	}

	// is it full house?
	if len(frequencies[2]) == 1 && len(frequencies[1]) == 1 {
		triple := frequencies[2][0]
		double := frequencies[1][0]
		return []int{
			ranks["fullhouse"] + triple,
			double,
		}
	}

	// is it three of a kind?
	if len(frequencies[2]) == 1 {
		value := frequencies[2][0]
		singles := frequencies[0]
		return []int{
			ranks["threeofakind"] + value,
			singles[0],
			singles[1],
		}
	}

	// is it two pair?
	if len(frequencies[1]) == 2 {
		high := frequencies[1][0]
		low := frequencies[1][1]
		single := frequencies[0][0]
		return []int{
			ranks["twopair"] + high,
			low,
			single,
		}
	}

	// is it one pair?
	if len(frequencies[1]) == 1 {
		value := frequencies[1][0]
		singles := frequencies[0]
		return []int{
			ranks["onepair"] + value,
			singles[0],
			singles[1],
			singles[2],
		}
	}

	// must be high card
	return frequencies[0]
}
