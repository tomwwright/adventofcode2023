package main

import "testing"

func TestExtrapolateSequence(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			[]int{0, 3, 6, 9, 12, 15},
			18,
		},
		{
			[]int{1, 3, 6, 10, 15, 21},
			28,
		},
		{
			[]int{10, 13, 16, 21, 30, 45},
			68,
		},
		{
			[]int{1, 1, 2, 4, 7},
			11,
			// 1 1 2 4 7 11
			//  0 1 2 3 4
			//   1 1 1 1
			//    0 0 0
		},
		{
			[]int{13, 9, 6, 4, 3, 3},
			4,
			// 13  9  6  4  3  3  4
			//   4  3  2  1  0  -1
			//    -1 -1 -1 -1 -1
			//      0 0 0 0  0
		},
	}
	for _, tc := range testCases {
		n := extrapolateSequence(tc.input)
		if n != tc.expected {
			t.Error("test failed", n, tc)
		}
	}
}

func TestExtrapolateSequenceBackwards(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			[]int{0, 3, 6, 9, 12, 15},
			-3,
		},
		{
			[]int{1, 3, 6, 10, 15, 21},
			//[0] 1 3 6 10 15 21 [28]
			// [1] 2 3 4  5  6 [7]
			//  [1] 1 1 1  1 [1]
			//   [0] 0 0 0 [0]
			0,
		},
		{
			[]int{10, 13, 16, 21, 30, 45},
			//   [5] 10  13  16  21  30  45 [68]
			//    [5]	 3   3   5   9  15  [23]
			//     [-2]  0   2   4   6  [8]
			//        [2]  2   2   2  [2]
			//	        [0]  0   0  [0]
			5,
		},
		{
			[]int{1, 1, 2, 4, 7},
			//[2] 1  1  2  4  7 [11]
			// [-1] 0  1  2  3 [4]
			//   [1] 1  1  1  [1]
			//    [0] 0  0  [0]
			2,
		},
		{
			[]int{13, 9, 6, 4, 3, 3},
			//[18] 13  9  6  4  3  3  [4]
			//   [5] 4  3  2  1  0  [-1]
			//   [-1] -1 -1 -1 -1 [-1]
			//          0 0 0 0  0
			18,
		},
	}
	for _, tc := range testCases {
		n := extrapolateSequenceBackwards(tc.input)
		if n != tc.expected {
			t.Error("test failed", n, tc)
		}
	}
}
