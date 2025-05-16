// calculator_test.go
package calculator

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {

	// Table-driven tests
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"zero addition", 0, 0, 0},
		{"negative numbers", -1, -2, -3},
		{"mixed signs", -2, 3, 1},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

func ConcatPlus(s1, s2 string) string {
	return s1 + s2
}

func ConcatBuilder(s1, s2 string) string {
	var b strings.Builder
	b.WriteString(s1)
	b.WriteString(s2)
	return b.String()
}

func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ConcatPlus("hello", "world")
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ConcatBuilder("hello", "world")
	}
}
