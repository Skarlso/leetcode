package zigzagconversion

// func convert(s string, numRows int) string {
// 	if numRows < 2 {
// 		return s
// 	}

// 	var result []byte

// 	rows := make([][]byte, numRows)

// 	row := 0
// 	sub := false
// 	for _, c := range s {
// 		rows[row] = append(rows[row], byte(c))
// 		if sub {
// 			row--
// 			if row < 0 {
// 				row += 2
// 				sub = false
// 			}
// 		} else {
// 			row++
// 			if row >= numRows {
// 				sub = true
// 				row -= 2
// 			}
// 		}
// 	}

// 	for _, r := range rows {
// 		result = append(result, r...)
// 	}

// 	return string(result)
// }

// after looking at how the indexes are changing, the math approach was evident.
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	var result []byte

	// o(n)
	for row := 0; row < numRows; row++ {
		// dealing with the first or last row
		// increase by numRows -1 * 2 for the last and first row
		for i := row; i < len(s); i += (numRows - 1) * 2 {
			result = append(result, s[i])
			// deal with the in-between values. Catch the character in the middle and any
			// in between. it's always 2 times lower as we are going down the rows
			// so (numRows-1) * 2 AND ((numRows-1) * 2) - 2 * currentRow
			if row > 0 && row < numRows-1 {
				// fmt.Println(i)
				between := i + (numRows-1)*2 - (2 * row)
				if between < len(s) {
					result = append(result, s[between])
				}
			}
		}

	}

	return string(result)
}
