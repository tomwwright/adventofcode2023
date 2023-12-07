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
			values: []int{ranks["onepair"] + values['3'], values['K'], values['T'], values['2']},
			bid:    765,
		},
		{
			values: []int{ranks["twopair"] + values['J'], values['T'], values['K']},
			bid:    220,
		},
		{
			values: []int{ranks["twopair"] + values['K'], values['7'], values['6']},
			bid:    28,
		},
		{
			values: []int{ranks["threeofakind"] + values['5'], values['J'], values['T']},
			bid:    684,
		},
		{
			values: []int{ranks["threeofakind"] + values['Q'], values['A'], values['J']},
			bid:    483,
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

	if !slices.Equal(a.values, b.values) {
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
			values: []int{ranks["fullhouse"] + values['7'], values['8']},
			bid:    1,
			hand:   "77788",
		},
		{
			values: []int{ranks["fullhouse"] + values['8'], values['7']},
			bid:    2,
			hand:   "77888",
		},
		{
			values: []int{ranks["fourofakind"] + values['A'], values['2']},
			bid:    3,
			hand:   "2AAAA",
		},
		{
			values: []int{ranks["fourofakind"] + values['3'], values['2']},
			bid:    4,
			hand:   "33332",
		},
	}

	if !slices.EqualFunc(hands, expected, isEqualHand) {
		t.Error("test failed", hands, expected)
	}
}

func TestParseHand(t *testing.T) {
	testCases := []struct {
		input    string
		expected []int
	}{
		{
			"QQQQQ",
			[]int{ranks["fiveofakind"] + values['Q']},
		},
		{
			"8888T",
			[]int{ranks["fourofakind"] + values['8'], values['T']},
		},
		{
			"TT444",
			[]int{ranks["fullhouse"] + values['4'], values['T']},
		},
		{
			"32T3K",
			[]int{ranks["onepair"] + values['3'], values['K'], values['T'], values['2']},
		},
		{
			"T55J5",
			[]int{ranks["threeofakind"] + values['5'], values['J'], values['T']},
		},
		{
			"KTJJT",
			[]int{ranks["twopair"] + values['J'], values['T'], values['K']},
		},
		{
			"KK677",
			[]int{ranks["twopair"] + values['K'], values['7'], values['6']},
		},
		{
			"QQQJA",
			[]int{ranks["threeofakind"] + values['Q'], values['A'], values['J']},
		},
		{
			"23456",
			[]int{ranks["highcard"] + values['6'], values['5'], values['4'], values['3'], values['2']},
		},
	}
	for _, tc := range testCases {
		parsed := parseHand(tc.input)
		if !slices.Equal(parsed, tc.expected) {
			t.Error("test failed", parsed, tc)
		}
	}
}
