package leetcode

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:

P   A   H   N
A P L S I I G
Y   I   R

And then read line by line: "PAHNAPLSIIGYIR".
Write the code that will take a string and make this conversion given a number of rows:

string convert(string text, int nRows);
convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
*/

import (
	"bytes"
)

func convert1(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	bufAry := make([]bytes.Buffer, numRows+1)
	row, step, bound := 0, 1, numRows-1

	for i := range s {
		bufAry[row].WriteByte(s[i])

		if row == 0 {
			step = 1
		} else if row == bound {
			step = -1
		}

		row += step
	}

	for i := 0; i < numRows; i++ {
		bufAry[i].WriteTo(&bufAry[numRows])
	}

	return bufAry[numRows].String()
}

func convert2(s string, numRows int) string {
	if numRows <= 1 || numRows >= len(s) {
		return s
	}

	var buf bytes.Buffer
	slen, bound := len(s), numRows-1

	for i := 0; i < numRows; i++ {
		j := i
		step := 2 * (numRows - 1)
		step1 := 2 * (numRows - i - 1)
		step2 := 2 * i
		flip := true

		for j < slen {
			buf.WriteByte(s[j])

			if i == 0 || i == bound {
				j += step
			} else if flip {
				j += step1
			} else {
				j += step2
			}

			flip = !flip
		}
	}

	return buf.String()
}

func convert(s string, numRows int) string {
	return convert2(s, numRows)
}
