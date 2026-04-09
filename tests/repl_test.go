package main

import (
	"testing"

	"github.com/Eval-99/pokedexcli/internal/repl"
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
			input:    "  hello  world, how are you today?",
			expected: []string{"hello", "world,", "how", "are", "you", "today?"},
		},
	}

	for _, c := range cases {
		actual := repl.CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Actual result length does not equal expected length")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word does not equal expected word")
			}
		}
	}
}
