package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestParseNumbers(t *testing.T) {
	testCases := []struct {
		str      string
		expected []int
	}{
		{
			"41 48 83 86 17",
			[]int{41, 48, 83, 86, 17},
		},
		{
			"  1   21 5  100   6",
			[]int{1, 21, 5, 100, 6},
		},
	}
	for _, tc := range testCases {
		parsed := parseNumbers(tc.str)
		if !slices.Equal(parsed, tc.expected) {
			t.Error("test failed", parsed, tc)
		}
	}
}

func TestParseCard(t *testing.T) {
	testCases := []struct {
		input   string
		winners []int
		have    []int
	}{

		{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			[]int{41, 48, 83, 86, 17},
			[]int{83, 86, 6, 31, 17, 9, 48, 53},
		},
		{
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			[]int{13, 32, 20, 16, 61},
			[]int{61, 30, 68, 82, 17, 32, 24, 19},
		},
		{
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			[]int{1, 21, 53, 59, 44},
			[]int{69, 82, 63, 72, 16, 21, 14, 1},
		},
		{
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			[]int{41, 92, 73, 84, 69},
			[]int{59, 84, 76, 51, 58, 5, 54, 83},
		},
		{
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			[]int{87, 83, 26, 28, 32},
			[]int{88, 30, 70, 12, 93, 22, 82, 36},
		},
		{
			"Card 100: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			[]int{31, 18, 13, 56, 72},
			[]int{74, 77, 10, 23, 35, 67, 36, 11},
		},
	}
	for _, tc := range testCases {
		winners, have := parseCard(tc.input)
		if !slices.Equal(winners, tc.winners) || !slices.Equal(have, tc.have) {
			t.Error("test failed", winners, have, tc)
		}
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	sum := part2(input)
	if sum != 30 {
		t.Error("test failed", sum, 30)
	}
}
