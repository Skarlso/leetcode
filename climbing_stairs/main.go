package main

// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
// 4 =>
// 1. 1 + 1 + 1 + 1
// 2. 1 + 1 + 2
// 3. 2 + 1 + 1
// 4. 1 + 2 + 1
// 5. 2 + 2
// 5 =>
// 1. 1 + 1 + 1 + 1 + 1
// 2. 1 + 1 + 2 + 1
// 3. 2 + 1 + 1 + 1
// 4. 1 + 2 + 1 + 1
// 5. 2 + 2 + 1
// 6. 1 + 2 + 1 + 1
// 7. 1 + 1 + 2 + 1
// 8. 1 + 1 + 1 + 2
// 9. 2 + 1 + 2
// Time limit exceeded.
func climbStairsV1(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return climbStairsV1(n-1) + climbStairsV1(n-2)
}

func climbStairs(n int) int {
	one, two := 1, 1

	for i := 0; i < n-1; i++ {
		one, two = one+two, one
	}

	return one
}
