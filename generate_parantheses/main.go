package main

func generateParenthesis(n int) []string {
	var result []string
	paramBacktrack([]byte{}, &result, n, 0, 0)
	return result
}

func paramBacktrack(path []byte, result *[]string, end, open, closed int) {
	if open == end && closed == end {
		*result = append(*result, string(path))
		return
	}

	if closed < open {
		path = append(path, ')')
		paramBacktrack(path, result, end, open, closed+1)
		path = path[:len(path)-1]
	}
	if open < end {
		path = append(path, '(')
		paramBacktrack(path, result, end, open+1, closed)
		path = path[:len(path)-1]
	}
}
