package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	t.Run("sum all slices past in", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := []int{1, 2, 3}
		got := SumAll(slice1, slice2)
		want := []int{15, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but wanted %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 9}, []int{1, 2}, []int{1,2,3})
		want := []int{9, 2, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but wanted %v", got, want)
		}
	})
}
