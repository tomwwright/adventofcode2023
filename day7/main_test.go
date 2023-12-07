package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestParseInput(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	hands := parseInput(input)
	expected := []Hand{
		{
			rank: ranks["onepair"],
			bid:  765,
			hand: "32T3K",
		},
		{
			rank: ranks["twopair"],
			bid:  220,
			hand: "KTJJT",
		},
		{
			rank: ranks["twopair"],
			bid:  28,
			hand: "KK677",
		},
		{
			rank: ranks["threeofakind"],
			bid:  684,
			hand: "T55J5",
		},
		{
			rank: ranks["threeofakind"],
			bid:  483,
			hand: "QQQJA",
		},
	}

	if !slices.EqualFunc(hands, expected, isEqualHand) {
		t.Error("test failed", hands, expected)
	}

}

func isEqualHand(a Hand, b Hand) bool {
	if a.bid != b.bid {
		return false
	}

	if a.rank != b.rank {
		return false
	}

	if a.hand != b.hand {
		return false
	}

	return true
}

func TestBullshitSorting(t *testing.T) {
	input := `
33332 4
77888 2
77788 1
2AAAA 3
`

	hands := parseInput(input)
	expected := []Hand{
		{
			rank: ranks["fullhouse"],
			bid:  1,
			hand: "77788",
		},
		{
			rank: ranks["fullhouse"],
			bid:  2,
			hand: "77888",
		},
		{
			rank: ranks["fourofakind"],
			bid:  3,
			hand: "2AAAA",
		},
		{
			rank: ranks["fourofakind"],
			bid:  4,
			hand: "33332",
		},
	}

	if !slices.EqualFunc(hands, expected, isEqualHand) {
		t.Error("test failed", hands, expected)
	}
}

func TestParseRank(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			"QQQQQ",
			ranks["fiveofakind"],
		},
		{
			"8888T",
			ranks["fourofakind"],
		},
		{
			"TT444",
			ranks["fullhouse"],
		},
		{
			"32T3K",
			ranks["onepair"],
		},
		{
			"T55J5",
			ranks["threeofakind"],
		},
		{
			"KTJJT",
			ranks["twopair"],
		},
		{
			"KK677",
			ranks["twopair"],
		},
		{
			"QQQJA",
			ranks["threeofakind"],
		},
		{
			"23456",
			ranks["highcard"],
		},
	}
	for _, tc := range testCases {
		parsed := parseRank(tc.input)
		if parsed != tc.expected {
			t.Error("test failed", parsed, tc)
		}
	}
}

func TestParseRankWithJokers(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			"QQQQQ",
			ranks["fiveofakind"],
		},
		{
			"8888T",
			ranks["fourofakind"],
		},
		{
			"TJ444",
			ranks["fourofakind"],
		},
		{
			"32TJK",
			ranks["onepair"],
		},
		{
			"T55J5",
			ranks["fourofakind"],
		},
		{
			"KTJJT",
			ranks["fourofakind"],
		},
		{
			"KK677",
			ranks["twopair"],
		},
		{
			"QQQJA",
			ranks["fourofakind"],
		},
		{
			"23456",
			ranks["highcard"],
		},
	}
	for _, tc := range testCases {
		parsed := parseRankWithJokers(tc.input)
		if parsed != tc.expected {
			t.Error("test failed", parsed, tc)
		}
	}
}
