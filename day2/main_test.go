package main

import (
	"fmt"
	"testing"
)

func TestParseRound(t *testing.T) {
	testCases := []struct {
		round    string
		expected Round
	}{
		{
			"5 blue, 2 green, 7 red",
			Round{
				blue:  5,
				red:   7,
				green: 2,
			},
		},
		{
			" 14 red, 15 green, 1 blue",
			Round{
				blue:  1,
				red:   14,
				green: 15,
			},
		},
		{
			"1 green, 3 blue, 14 red",
			Round{
				blue:  3,
				red:   14,
				green: 1,
			},
		},
	}
	for _, tc := range testCases {
		parsed := ParseRound(tc.round)
		fmt.Println(parsed)
		if parsed != tc.expected {
			t.Error("test failed", parsed, tc)
		}
	}
}
