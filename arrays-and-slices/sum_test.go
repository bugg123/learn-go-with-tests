package main

import "testing"

func assertEqual(actual, expected int, numbers []int, t *testing.T) {
	if actual != expected {
		t.Errorf("expected %d but got %d, %v", expected, actual, numbers)
	}
}

func TestSum(t *testing.T) {
	t.Run("sum array of 5 things", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		actual := Sum(numbers)
		expected := 15
		assertEqual(actual, expected, numbers, t)
	})
}
