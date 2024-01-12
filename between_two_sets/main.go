package main

import "fmt"

func getTotalX(a []int32, b []int32) int32 {
	// something something GCD
	// Write your code here
	// 6 can be divided by all the elements of a
	// and 6 can divide all the elements of b
	// we can be sure that a[0] == min and b[len(b)-1] == max

	var num int32
loop:
	for i := a[len(a)-1]; i <= b[0]; i++ {
		for _, n := range a {
			if i%n != 0 {
				continue loop
			}
		}

		for _, n := range b {
			if n%i != 0 {
				continue loop
			}
		}

		fmt.Println(i)
		num++
	}

	return num
}
