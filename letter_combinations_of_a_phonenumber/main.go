package main

var letterMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	var result []string
	backtrack([]byte{}, &result, digits, 0)
	return result
}

func backtrack(path []byte, result *[]string, digits string, index int) {
	if index >= len(digits) {
		temp := make([]byte, len(path))
		copy(temp, path)
		*result = append(*result, string(temp))
		return
	}

	if digits[index] == '1' {
		return
	}

	for _, d := range letterMap[digits[index]] {
		path = append(path, d)
		backtrack(path, result, digits, index+1)
		path = path[:len(path)-1]
	}
}
