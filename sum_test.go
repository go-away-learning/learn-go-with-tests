package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum numbers slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", want, got, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSlicesSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("sums tails of slices", func(t *testing.T) {
		got := SumAllTais([]int{0, 1, 2}, []int{9, 8, 7}, []int{100})
		want := []int{3, 15, 0}

		checkSlicesSum(t, got, want)
	})

	t.Run("sums tails with empty slice", func(t *testing.T) {
		got := SumAllTais([]int{0, 2}, []int{0, 5, 5}, []int{})
		want := []int{2, 10, 0}

		checkSlicesSum(t, got, want)
	})
}
