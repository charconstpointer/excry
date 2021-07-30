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
		err      bool
	}
	tests := []test{
		{
			name:     "retrieves key on first level",
			input:    `{"foo":"bar"}`,
			key:      "foo",
			expected: true,
			err:      false,
		},
		{
			name:     "retrieves key on second level",
			input:    `{"root":{"foo":"bar"}}`,
			key:      "foo",
			expected: true,
			err:      false,
		},
		{
			name:     "retrieves key on third level",
			input:    `{"root":{"notyet":{"foo":"bar"}}}`,
			key:      "foo",
			expected: true,
			err:      false,
		},
	}
	for _, tt := range tests {
		exists, err := Exists(tt.input, tt.key)
		if err != nil && !tt.err {
			t.Errorf("%s: Exists(%s, %s): unexpected error: %v", tt.name, tt.input, tt.key, err)
		}
		if exists != tt.expected {
			t.Errorf("expected to be %v instead got %v", tt.expected, exists)
		}
	}
}

func TestExistsWithVal(t *testing.T) {
	type test struct {
		name     string
		input    string
		key      string
		val      string
		expected bool
		err      bool
	}
	tests := []test{
		{
			name:     "retrieves key on first level",
			input:    `{"foo":"bar"}`,
			val:      "bar",
			key:      "foo",
			expected: true,
			err:      false,
		},
		{
			name:     "retrieves key on second level",
			input:    `{"root":{"foo":"bar"}}`,
			key:      "foo",
			val:      "bar",
			expected: true,
			err:      false,
		},
		{
			name:     "retrieves key on third level",
			input:    `{"root":{"notyet":{"foo":"baz"}}}`,
			key:      "foo",
			val:      "baz",
			expected: true,
			err:      false,
		},
	}
	for _, tt := range tests {
		ok, err := ExistsWithVal(tt.input, tt.key, tt.val)
		if err != nil && !tt.err {
			t.Errorf("%s: Exists(%s, %s): unexpected error: %v", tt.name, tt.input, tt.key, err)
		}
		if ok != tt.expected {
			t.Errorf("expected to be %v instead got %v", tt.expected, ok)
		}

	}
}
