package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{1, 3, 7})
	want := []int{3, 11}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Make sums of some slices", func(t *testing.T) {

		got := SumAllTails([]int{1, 4, 2}, []int{2, 3})
		want := []int{6, 3}

		checkSums(t, got, want)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{2, 3})
		want := []int{0, 3}

		checkSums(t, got, want)
	})
}
