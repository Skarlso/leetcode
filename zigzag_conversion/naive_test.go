package zigzagconversion

import "testing"

func TestConvert(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		rows   int
		result string
	}{
		{
			name:   "PAHNAPLSIIGYIR",
			input:  "PAYPALISHIRING",
			rows:   3,
			result: "PAHNAPLSIIGYIR",
		},
		{
			name:   "PINALSIGYAHRPI",
			input:  "PAYPALISHIRING",
			rows:   4,
			result: "PINALSIGYAHRPI",
		},
		{
			name:   "A",
			input:  "A",
			rows:   1,
			result: "A",
		},
		{
			name:   "AB",
			input:  "AB",
			rows:   1,
			result: "AB",
		},
	}

	for _, tc := range testcases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			got := convert(tc.input, tc.rows)
			if got != tc.result {
				t.Fatalf("got %s; want %s", got, tc.result)
			}
		})
	}
}
