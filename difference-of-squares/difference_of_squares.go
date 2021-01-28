package diffsquares

import (
	"math"
)

func SquareOfSum(n int) int {
	sum := float64(n+1) * (float64(n) / float64(2))
	return int(math.Pow(sum, 2))
}

func SumOfSquares(n int) int {
	return int(float64((n*(n+1))*(n+(n+1))) / 6)
}

func difference(squareOfSum, SumOfSquares int) int {
	return squareOfSum - SumOfSquares
}
