package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestParseSeeds(t *testing.T) {
	testCases := []struct {
		input    string
		expected []int
	}{
		{
			" 79 14 55 13",
			[]int{79, 14, 55, 13},
		},
	}
	for _, tc := range testCases {
		parsed := parseSeeds(tc.input)
		if !slices.Equal(parsed, tc.expected) {
			t.Error("test failed", parsed, tc)
		}
	}
}

func TestParseSeedRanges(t *testing.T) {
	testCases := []struct {
		input    string
		expected []int
	}{
		{
			" 79 3 55 1",
			[]int{79, 80, 81, 55},
		},
		{
			"11 8 55 1 6 3",
			[]int{11, 12, 13, 14, 15, 16, 17, 18, 55, 6, 7, 8},
		},
	}
	for _, tc := range testCases {
		parsed := parseSeedRanges(tc.input)
		if !slices.Equal(parsed, tc.expected) {
			t.Error("test failed", parsed, tc)
		}
	}
}

func TestTranslate(t *testing.T) {
	mappings := []Mapping{
		{
			source:      50,
			destination: 98,
			length:      2,
		},
		{
			source:      52,
			destination: 50,
			length:      48,
		},
	}
	testCases := []struct {
		input    int
		expected int
	}{
		{
			79, 81,
		},
		{
			14, 14,
		},
		{
			55, 57,
		},
		{
			13, 13,
		},
		{
			100, 100,
		},
	}
	for _, tc := range testCases {
		n := translate(tc.input, mappings)
		if n != tc.expected {
			t.Error("test failed", n, tc)
		}
	}
}

func TestParseMapping(t *testing.T) {
	testCases := []struct {
		input    string
		expected []Mapping
	}{
		{
			"50 98 2\n52 50 48",
			[]Mapping{
				{
					source:      50,
					destination: 98,
					length:      2,
				},
				{
					source:      52,
					destination: 50,
					length:      48,
				},
			},
		},
	}
	for _, tc := range testCases {
		parsed := parseMappings(tc.input)
		if !slices.Equal(parsed, tc.expected) {
			t.Error("test failed", parsed, tc)
		}
	}
}
