package util

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandHexString(t *testing.T) {
	testCases := []struct {
		name   string
		length int
	}{
		{
			name:   "Hex string with length 1",
			length: 1,
		},
		{
			name:   "Hex string with length 8",
			length: 8,
		},
		{
			name:   "Hex string with length 10",
			length: 10,
		},
		{
			name:   "Hex string with length 31",
			length: 31,
		},
	}

	const iterations = 5

	for _, tc := range testCases {
		for i := 0; i < iterations; i++ {
			t.Run(tc.name, func(t *testing.T) {
				result := RandHexString(tc.length)

				assert.Equal(t, tc.length, len(result), "Iteration %d: expected length %d, but got %d", tc.length, len(result))

				// Check if the result contains only valid hex characters
				isHex := regexp.MustCompile(`^[0-9A-F]*$`).MatchString
				assert.True(t, isHex(result), "Iteration %d: expected hex string, but got %s", result)
			})
		}
	}
}

func BenchmarkRandHexString(b *testing.B) {
	benchmarkCases := []struct {
		name   string
		length int
	}{
		{name: "Length_5", length: 5},
		{name: "Length_10", length: 10},
		{name: "Length_16", length: 16},
		{name: "Length_16", length: 20},
		{name: "Length_32", length: 32},
	}

	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RandHexString(bc.length)
			}
		})
	}
}
