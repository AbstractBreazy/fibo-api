package fib

import (
	"testing"
)

// FibonacciTest is a simple test for Fibonacci function
func TestFibonacci(t *testing.T) {
	values := []struct {
		number   uint
		expected int64
	}{
		{0, 0}, {1, 1}, {50, 12586269025},
		{15, 610}, {25, 75025}, {60, 1548008755920},
	}
	for _, v := range values {
		if result := Fibonacci(v.number); result.Int64() != v.expected {
			t.Errorf("Wrong Fibonacci value for: %v, got: %v, want: %v", v.number, result, v.expected)
		}
	}
}
