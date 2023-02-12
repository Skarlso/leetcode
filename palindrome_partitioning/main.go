package main

func partition(s string) [][]string {
	var result [][]string
	backtrack([]string{}, &result, s, 0)
	return result
}

func backtrack(path []string, result *[][]string, s string, index int) {
	if index >= len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for i := index; i < len(s); i++ {
		if isPalindrome(s, index, i) {
			path = append(path, string(s[index:i+1]))
			backtrack(path, result, s, i+1)
			path = path[:len(path)-1]
		}
	}

}

func isPalindrome(s string, start, end int) bool {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
