package main

import (
	"strings"
	"unicode"
)

// func letterCasePermutation(s string) []string {
// 	return nil
// }

func letterCasePermutation(s string) []string {
	var result []string
	if s == "" {
		result = append(result, "")
		return result
	}

	perms := letterCasePermutation(s[1:])
	for _, v := range perms {
		if unicode.IsDigit(rune(s[0])) {
			result = append(result, string(s[0])+v)
		} else {
			result = append(result, strings.ToLower(string(s[0]))+v)
			result = append(result, strings.ToUpper(string(s[0]))+v)
		}
	}
	return result
}
