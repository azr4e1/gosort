package gosort_test

import (
	"github.com/google/go-cmp/cmp"

	gs "gosort"
	"testing"
)

type TestCase struct {
	Input []int
	Value []int
}

var testCases []TestCase = []TestCase{
	{Input: []int{1, 2, 3, 4}, Value: []int{1, 2, 3, 4}},
	{Input: []int{4, 3, 2, 3, 532, 43, 12, 5, 34, 23}, Value: []int{2, 3, 3, 4, 5, 12, 23, 34, 43, 532}},
	{Input: []int{}, Value: []int{}},
	{Input: []int{1}, Value: []int{1}},
	{Input: []int{9, 9, 9, 9, 9, 9, 9}, Value: []int{9, 9, 9, 9, 9, 9, 9}},
	{Input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, Value: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
}

func TestBubbleSort_SortsTheGivenArray(t *testing.T) {
	t.Parallel()
	for _, tc := range testCases {
		want := tc.Value
		got := make([]int, len(tc.Input))
		copy(got, tc.Input)
		gs.BubbleSort(got)

		if !cmp.Equal(got, want) {
			t.Error(cmp.Diff(got, want))
		}
	}
}

func TestInsertionSort_SortsTheGivenArray(t *testing.T) {
	t.Parallel()
	for _, tc := range testCases {
		want := tc.Value
		got := make([]int, len(tc.Input))
		copy(got, tc.Input)
		gs.InsertionSort(got)

		if !cmp.Equal(got, want) {
			t.Error(cmp.Diff(got, want))
		}
	}
}
