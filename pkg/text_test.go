package vq_test

import (
	"testing"

	"github.com/leorolland/vq/parser"
	vq "github.com/leorolland/vq/pkg"
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
		"invalid char": {
			input:         "~",
			expectedError: parser.ErrNoMatch,
		},
		"brackets": {
			input:         "foo[]bar",
			expectedError: parser.ErrUnconsumedInput,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, err := parser.Parse(vq.Text(), tc.input)
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
		expected string
	}{
		"ascii": {
			input:    "foo",
			expected: "foo",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			output, err := parser.Parse(vq.Text(), tc.input)
			if err != nil {
				t.FailNow()
			}
			if output != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, output)
			}
		})
	}
}
