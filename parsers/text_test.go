package parsers_test

import (
	"reflect"
	"testing"

	"github.com/leorolland/vq"
	"github.com/leorolland/vq/parser"
	"github.com/leorolland/vq/parsers"
)

func TestTextError(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         string
		expectedError error
	}{
		"empty string": {
			input:         "",
			expectedError: parser.ErrNoMatch,
		},
		"brackets": {
			input:         "foo[]bar",
			expectedError: parser.ErrUnconsumedInput,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, err := parser.Parse(parsers.Text(), tc.input)
			if err != tc.expectedError {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestTextSuccess(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    string
		expected vq.Node
	}{
		"ascii": {
			input:    "foo",
			expected: vq.Node{Kind: vq.Text, Value: "foo"},
		},
		"ascii with space": {
			input:    "foo bar",
			expected: vq.Node{Kind: vq.Text, Value: "foo bar"},
		},
		"ascii with space and unicode": {
			input:    "丒专	且 é",
			expected: vq.Node{Kind: vq.Text, Value: "丒专	且 é"},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			output, err := parser.Parse(parsers.Text(), tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("expected %s, got %s", tc.expected, output)
			}
		})
	}
}
