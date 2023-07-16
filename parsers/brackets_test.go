package parsers_test

import (
	"testing"

	"github.com/leorolland/vq/parser"
	"github.com/leorolland/vq/parsers"
)

func TestBracketsSuccess(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    string
		expected string
	}{
		"empty brackets": {
			input:    "[]",
			expected: "Brackets()",
		},
		"ascii inside": {
			input:    "[foo]",
			expected: "Brackets(foo)",
		},
		"unicode inside": {
			input:    "[àà bb çç]",
			expected: "Brackets(àà bb çç)",
		},
		"recursion empty": {
			input:    "[a[]b]",
			expected: "Brackets(a, Brackets(), b)",
		},
		"1 direct recursion": {
			input:    "[[foo]]",
			expected: "Brackets(Brackets(foo))",
		},
		"2 direct recursions": {
			input:    "[[[foo]]]",
			expected: "Brackets(Brackets(Brackets(foo)))",
		},
		"recursion with text": {
			input:    "[a[bb]c]",
			expected: "Brackets(a, Brackets(bb), c)",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			output, err := parser.Parse(parsers.Brackets(2), tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if output != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, output)
			}
		})
	}
}
