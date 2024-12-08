package main

import (
	"reflect"
	"testing"
)

func TestGetData(t *testing.T) {
	got := getData("test.txt")
	want := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("got %v wanted %v", got, want)
	}
}

func TestCheckReport(t *testing.T) {
	got := checkReport([]int{7, 6, 4, 2, 1})
	want := true
	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
