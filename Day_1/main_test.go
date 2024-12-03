package main

import (
	"reflect"
	"testing"
)

func TestGetData(t *testing.T) {
	got := getData("test.txt")
	want := []int{3, 4, 4, 3, 2, 5, 1, 3, 3, 9, 3, 3}
	if reflect.DeepEqual(got, want) != true {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSplitArray(t *testing.T) {
	twoArrays := getData("test.txt")
	gotFirst, gotSecond := splitArray(twoArrays)
	wantFirst, wantSecond := []int{1, 2, 3, 3, 3, 4}, []int{3, 3, 3, 4, 5, 9}
	pass := false
	for i := range gotFirst {
		if gotFirst[i] == wantFirst[i] && gotSecond[i] == wantSecond[i] {
			pass = true
		}
	}
	if pass != true {
		t.Errorf("got %q & %q want %q & %q", gotFirst, gotSecond, wantFirst, wantSecond)
	}
}

func TestSumLists(t *testing.T) {
	got := sumLists(splitArray(getData("test.txt")))
	var want int = 11

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func TestMultiplyLists(t *testing.T) {
	got := multiplyLists(splitArray(getData("test.txt")))
	want := 31
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
