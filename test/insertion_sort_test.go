package test

import (
	"mubashir/sorting"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{input: []int{100, 80, 40, 60, 20, 30}, expected: []int{20, 30, 40, 60, 80, 100}},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, test := range tests {
		output, _ := sorting.InsertionSort(test.input)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("got %v want %v", output, test.expected)
		}
	}
}
