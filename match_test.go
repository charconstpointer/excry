package excry

import (
	"testing"
)

func TestExists(t *testing.T) {
	type test struct {
		name     string
		input    string
		key      string
		expected bool
	}
	tests := []test{
		{
			name:     "retrieves key on first level",
			input:    `{"foo":"bar"}`,
			key:      "foo",
			expected: true,
		},
		{
			name:     "retrieves key on second level",
			input:    `{"root":{"foo":"bar"}}`,
			key:      "foo",
			expected: true,
		},
	}
	for _, tt := range tests {
		exists := Exists(tt.input, tt.key)
		if exists != tt.expected {
			t.Errorf("expected to be %v instead got %v", tt.expected, exists)
		}
	}
}
