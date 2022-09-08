package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("should add 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		wanted := 15

		assertCorrectMessage(t, got, wanted)
	})

	t.Run("should add 8 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

		got := Sum(numbers)
		wanted := 36

		assertCorrectMessage(t, got, wanted)
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, wanted []int) {
		t.Helper()
		if !reflect.DeepEqual(got, wanted) {
			t.Errorf("got %v want %v", got, wanted)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		wanted := []int{2, 9}
		checkSums(t, got, wanted)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		wanted := []int{0, 9}
		checkSums(t, got, wanted)
	})

}

func assertCorrectMessage(t testing.TB, got int, wanted int) {
	t.Helper()
	if got != wanted {
		t.Errorf("\nActual: %d\nExpected: %d", got, wanted)
	}
}
