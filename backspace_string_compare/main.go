package main

// Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.
// Note that after backspacing an empty text, the text will continue empty.
// I could solve this with a stack.. Putting in items, and if I encounter a `#` then take out the last item.
// There supposed to be a two pointer solution to this too.
func backspaceCompare(s string, t string) bool {
	build := func(v string) string {
		var b []byte
		for _, c := range v {
			if c == '#' && len(b) > 0 {
				b = b[:len(b)-1]
			} else if c != '#' {
				b = append(b, byte(c))
			}
		}
		return string(b)
	}
	return build(s) == build(t)
}

func backspaceCompareString(s string, t string) bool {
	build := func(v string) string {
		var b string
		for _, c := range v {
			if c == '#' && len(b) > 0 {
				b = b[:len(b)-1]
			} else if c != '#' {
				b += string(c)
			}
		}
		return b
	}
	return build(s) == build(t)
}
