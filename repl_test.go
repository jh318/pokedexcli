package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		if len(actual) != len(c.expected) {
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			t.Errorf("input: %q, actual length: %d \n expected words: %q, expected length: %d", c.input, len(actual), c.expected, len(c.expected) )
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				// if they don't match, use t.Errorf to print an error message
				// and fail the test
				t.Errorf("Input %q, index %d: got %q, expected %q", c.input, i, word, expectedWord)
			}
		}
	}
}