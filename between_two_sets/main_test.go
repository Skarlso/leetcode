package main

import "testing"

func TestGetTotalX(t *testing.T) {
	num := getTotalX([]int32{2, 6}, []int32{24, 36})
	if num != 2 {
		t.Fatal("num should have been 2 but got: ", num)
	}

	num = getTotalX([]int32{2, 4}, []int32{16, 32, 96})
	if num != 3 {
		t.Fatal("num should have been 3 but got: ", num)
	}
}
