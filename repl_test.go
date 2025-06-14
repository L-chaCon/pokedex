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
			input:    "    ",
			expected: []string{},
		},
		{
			input:    "umbreon eevee",
			expected: []string{"umbreon", "eevee"},
		},
		{
			input:    "umbReoN eEvEe",
			expected: []string{"umbreon", "eevee"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Error("Lengh of input is different from expected")
		}

		for i := range actual {
			word := actual[i]
			expectWord := c.expected[i]
			if word != expectWord {
				t.Errorf(`word: %s is not expected: %s`, word, expectWord)
			}
		}
	}
}
