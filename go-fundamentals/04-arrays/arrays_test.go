package arrays_test

import (
	"arrays"
	"testing"
)

func TestArraysSum(t *testing.T) {
	num := [5]int{1, 2, 3, 4, 5}
	got := arrays.ArraysSum(num)
	want := 14

	if got != want {
		t.Errorf("got: '%d' , want: '%d' , given: '%v'", got, want, num)
	}
}
