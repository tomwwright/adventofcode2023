package main

import (
	"reflect"
	"strings"
	"testing"
)

var input = strings.TrimSpace(`
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`)

func TestParseInput(t *testing.T) {

	records := parseInput(input)
	expected := []Record{
		{
			[]int{1, 1, 3},
			"???.###",
		},
		{
			[]int{1, 1, 3},
			".??..??...?##.",
		},
		{
			[]int{1, 3, 1, 6},
			"?#?#?#?#?#?#?#?",
		},
		{
			[]int{4, 1, 1},
			"????.#...#...",
		},
		{
			[]int{1, 6, 5},
			"????.######..#####.",
		},
		{
			[]int{3, 2, 1},
			"?###????????",
		},
	}

	if !reflect.DeepEqual(records, expected) {
		t.Error("test failed", "\n", true, "\n", false)
	}
}
