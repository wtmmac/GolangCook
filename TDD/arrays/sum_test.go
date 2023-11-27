package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collections of any size:", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d give, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9}, []int{2, 3})
	want := []int{3, 9, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
