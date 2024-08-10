package testcase

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	testCases := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			name:     "Single integer",
			input:    []any{1},
			expected: "01",
		},
		{
			name:     "Integer and string",
			input:    []any{1, "test"},
			expected: "01 test",
		},
		{
			name:     "Slice of strings",
			input:    []any{1, []string{"test", "value"}},
			expected: "01 test value",
		},
		{
			name:     "Map of strings",
			input:    []any{1, map[string]string{"key1": "value1", "key2": "value2"}},
			expected: "01 value1 value2",
		},
		{
			name: "Struct with fields",
			input: []any{1, struct {
				Field1 string
				Field2 int
				Field3 string
			}{"value1", 2, ""}},
			expected: "01 value1 2",
		},
		{
			name: "Different types",
			input: []any{1, "test", struct {
				Field1 string
				Field2 int
				Field3 string
			}{"value1", 2, ""}},
			expected: "01 test value1 2",
		},
		{
			name:     "Long string",
			input:    []any{1, strings.Repeat("a", 70)},
			expected: "01 " + strings.Repeat("a", 61) + "...",
		},
		{
			name:     "Special characters",
			input:    []any{1, "test", "/path"},
			expected: "01 test %path",
		},
		{
			name:     "Empty struct",
			input:    []any{1, struct{}{}},
			expected: "01",
		},
		{
			name:     "Nil pointer",
			input:    []any{1, (*string)(nil)},
			expected: "01",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Name(tc.input...)
			assert.Equal(t, tc.expected, result, "Unexpected result for case: %s", tc.name)
		})
	}
}
