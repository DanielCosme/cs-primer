package protobufvarint

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	testCases := map[string]struct {
		input    uint64
		expected []byte
	}{
		// "With input 1": {
		// 	path:     "./1.uint64",
		// 	expected: []byte{0x1},
		// },
		"With input 150": {
			input:    150,
			expected: []byte{0x96, 0x01},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			res := encode(tc.input)
			if !bytes.Equal(res, tc.expected) {
				t.Fatalf("\nWant: %#x\n Got: %#x", tc.expected, res)
			}
		})
	}
}

func TestDecode(t *testing.T) {

}
