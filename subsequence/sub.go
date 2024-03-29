package subsequence

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	sIndex := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[sIndex] {
			sIndex++
			if sIndex >= len(s) {
				return true
			}
		}
	}

	return sIndex == len(s)
}
