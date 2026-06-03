package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "MJ is the G.O.A.T.",
			expected: []string{"mj", "is", "the", "g.o.a.t."},
		},
		{
			input:    "Yesterday I woke up sucking a lemon",
			expected: []string{"yesterday", "i", "woke", "up", "sucking", "a", "lemon"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		
		if len(actual) != len(c.expected) {
			t.Errorf("FAIL: actual output size does not match expected")
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("FAIL: input word(s) do not match expected")
			}
		}
	}
}