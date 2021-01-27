package grains

import (
	"fmt"
	"math"
)

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("error")
	}
	return uint64(math.Exp2(float64(n - 1))), nil
}

func Total() uint64 {
	var sum uint64
	for i := 0; i < 64; i++ {
		sum += uint64(math.Exp2(float64(i)))
	}
	return sum
}
