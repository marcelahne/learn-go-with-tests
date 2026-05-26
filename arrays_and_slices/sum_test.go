package arrays_and_slices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("two collections of numbers", func(t *testing.T) {
		numbers_one := []int{1, 2}
		numbers_two := []int{0, 9}

		got := SumAll(numbers_one, numbers_two)
		want := []int{3, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("single collection of numbers", func(t *testing.T) {
		numbers := []int{1, 1, 1}

		got := SumAll(numbers)
		want := []int{3}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}

	for b.Loop() {
		Sum(numbers)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{1, 2, 3}, []int{5, 6, 7})
	}
}
