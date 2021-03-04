package fib

import "math/big"

// Calculating the fibonacci sequence
func Fibonacci(y uint) *big.Int {
	if y <= 1 {
		return big.NewInt(int64(y))
	}
	var n2, n1 = big.NewInt(0), big.NewInt(1)
	for i := uint(1); i < y; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}
	return n1
}
