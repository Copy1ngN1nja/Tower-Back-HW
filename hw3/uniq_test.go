package main

import "testing"

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMake(t *testing.T) {
	tests := []struct {
		lines               []string
		c, d, u, i          bool
		numFields, numChars int
		expected            []string
	}{
		{ // test with nothing
			lines: []string{
				"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik.",
			},
			c:         false,
			d:         false,
			u:         false,
			i:         false,
			numFields: 0,
			numChars:  0,
			expected: []string{
				"I love music.", "", "I love music of Kartik.", "Thanks.", "I love music of Kartik.",
			},
		},
		{ // test with c enabled
			lines: []string{
				"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik.",
			},
			c:         true,
			d:         false,
			u:         false,
			i:         false,
			numFields: 0,
			numChars:  0,
			expected: []string{
				"3 I love music.", "1 ", "2 I love music of Kartik.",
				"1 Thanks.", "2 I love music of Kartik.",
			},
		},
		{ // test with d enabled
			lines: []string{
				"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik.",
			},
			c:         false,
			d:         true,
			u:         false,
			i:         false,
			numFields: 0,
			numChars:  0,
			expected: []string{
				"I love music.", "I love music of Kartik.", "I love music of Kartik.",
			},
		},
		{ // test with u enabled
			lines: []string{
				"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik.",
			},
			c:         false,
			d:         false,
			u:         true,
			i:         false,
			numFields: 0,
			numChars:  0,
			expected: []string{
				"", "Thanks.",
			},
		},
		{ // test with i enabled
			lines: []string{
				"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "", "I love MuSIC of Kartik.",
				"I love music of kartik.", "Thanks.", "I love music of kartik.", "I love MuSIC of Kartik.",
			},
			c:         false,
			d:         false,
			u:         false,
			i:         true,
			numFields: 0,
			numChars:  0,
			expected: []string{
				"I LOVE MUSIC.", "", "I love MuSIC of Kartik.", "Thanks.", "I love music of kartik.",
			},
		},
		{ // test with numFields
			lines: []string{
				"We love music.", "I love music.", "They love music.", "",
				"I love music of Kartik.", "We love music of Kartik.", "Thanks.",
			},
			c:         false,
			d:         false,
			u:         false,
			i:         false,
			numFields: 1,
			numChars:  0,
			expected: []string{
				"We love music.", "", "I love music of Kartik.", "Thanks.",
			},
		},
		{ // test with numChars
			lines: []string{
				"I love music.", "A love music.", "C love music.", "", "I love music of Kartik.",
				"We love music of Kartik.", "Thanks.",
			},
			c:         false,
			d:         false,
			u:         false,
			i:         false,
			numFields: 0,
			numChars:  1,
			expected: []string{
				"I love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks.",
			},
		},
	}
	for idx, test := range tests {
		result := makeAns(test.lines, test.c, test.d, test.u, test.i, test.numFields, test.numChars)
		if !slicesEqual(result, test.expected) {
			t.Errorf("Fail at test %d!\nInput: %v\nc = %t, d = %t, u = %t, i = %t, numFields = %d, numChars = %d\nExpected: %v\nGot %v",
				idx, test.lines, test.c, test.d, test.u, test.i, test.numFields, test.numChars, test.expected, result)
		}
	}
}
