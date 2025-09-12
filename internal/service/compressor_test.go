package service

import (
	"reflect"
	"testing"
)

func TestCompressor(t *testing.T) {
	// Placeholder test
}

func TestDecompressor(t *testing.T) {
	// Placeholder test
}

func TestCharFrequency(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[rune]int
	}{
		{
			name:  "basic hello",
			input: "Hello Gophers!",
			expected: map[rune]int{
				'H': 1,
				'e': 2,
				'l': 2,
				'o': 2,
				' ': 1,
				'G': 1,
				'p': 1,
				'h': 1,
				'r': 1,
				's': 1,
				'!': 1,
			},
		},
		{
			name:     "empty string",
			input:    "",
			expected: map[rune]int{},
		},
		{
			name:  "repeated runes",
			input: "aaa",
			expected: map[rune]int{
				'a': 3,
			},
		},
		{
			name:  "unicode runes",
			input: "世世",
			expected: map[rune]int{
				'世': 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CharFrequency(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("\ninput: %q\nexpected: %#v\ngot:      %#v", tt.input, tt.expected, result)
			}
		})
	}
}
