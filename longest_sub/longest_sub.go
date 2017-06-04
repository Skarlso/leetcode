package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	max := 0
	streakCount := 0
	streak := make(map[rune]int)
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if v, ok := streak[runes[i]]; ok {
			i = v + 1
			streak = make(map[rune]int)
			if streakCount > max {
				max = streakCount
			}
			streakCount = 0
		}
		streak[runes[i]] = i
		streakCount++
	}
	if streakCount > max {
		max = streakCount
	}
	return max
}

func main() {
	fmt.Println("Longest: ", lengthOfLongestSubstring(""))
}
