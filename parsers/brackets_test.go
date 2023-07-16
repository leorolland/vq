package parsers_test

import (
	"reflect"
	"testing"

	"github.com/leorolland/vq"
	"github.com/leorolland/vq/parser"
	"github.com/leorolland/vq/parsers"
)

func TestBracketsSuccess(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    string
		expected vq.Node
	}{
		"empty brackets": {
			input:    "[]",
			expected: vq.Node{Kind: vq.Brackets},
		},
		"ascii inside": {
			input: "[foo]",
			expected: vq.Node{
				Kind: vq.Brackets,
				Children: []vq.Node{
					{Kind: vq.Text, Value: "foo"},
				},
			},
		},
		"unicode inside": {
			input: "[àà bb çç]",
			expected: vq.Node{
				Kind: vq.Brackets,
				Children: []vq.Node{
					{Kind: vq.Text, Value: "àà bb çç"},
				},
			},
		},
		"recursion empty": {
			input: "[a[]b]",
			expected: vq.Node{
				Kind: vq.Brackets,
				Children: []vq.Node{
					{Kind: vq.Text, Value: "a"},
					{Kind: vq.Brackets},
					{Kind: vq.Text, Value: "b"},
				},
			},
		},
		"1 direct recursion": {
			input: "[[foo]]",
			expected: vq.Node{
				Kind: vq.Brackets,
				Children: []vq.Node{
					{
						Kind: vq.Brackets,
						Children: []vq.Node{
							{Kind: vq.Text, Value: "foo"},
						},
					},
				},
			},
		},
		"2 direct recursions": {
			input: "[[[foo]]]",
			expected: vq.Node{
				Kind: vq.Brackets,
				Children: []vq.Node{
					{
						Kind: vq.Brackets,
						Children: []vq.Node{
							{
								Kind: vq.Brackets,
								Children: []vq.Node{
									{Kind: vq.Text, Value: "foo"},
								},
							},
						},
					},
				},
			},
		},
		"recursion with text": {
			input: "[a[bb]c]",
			expected: vq.Node{
				Kind: vq.Brackets,
				Children: []vq.Node{
					{Kind: vq.Text, Value: "a"},
					{
						Kind: vq.Brackets,
						Children: []vq.Node{
							{Kind: vq.Text, Value: "bb"},
						},
					},
					{Kind: vq.Text, Value: "c"},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			output, err := parser.Parse(parsers.Brackets(3), tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("expected %s, got %s", tc.expected, output)
			}
		})
	}
}
