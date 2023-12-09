package main

import "testing"

func TestIsEndState(t *testing.T) {
	testCases := []struct {
		input    []string
		expected bool
	}{
		{
			[]string{"AAZ", "AAA", "BBB", "ZAZ", "AZA"},
			false,
		},
		{
			[]string{"AAZ"},
			true,
		},
		{
			[]string{"AAZ", "ZAZ", "AZA"},
			false,
		},
		{
			[]string{"AAZ", "ZAZ"},
			true,
		},
	}
	for _, tc := range testCases {
		parsed := isEndState(tc.input)
		if parsed != tc.expected {
			t.Error("test failed", parsed, tc)
		}
	}
}
