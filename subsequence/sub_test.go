package subsequence

import "testing"

func TestSubsequence(t *testing.T) {
	tcs := []struct {
		name   string
		input  string
		target string
		want   bool
	}{
		{
			name:   "abc",
			input:  "abc",
			target: "ahbgdc",
			want:   true,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := isSubsequence(tc.input, tc.target)
			if !got {
				t.Fatal("input didn't match target:", tc.input, tc.target)
			}
		})
	}
}
