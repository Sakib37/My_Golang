package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	negativeNumber := fmt.Sprint(float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: %v", negativeNumber)
}

func SqrtError(x float64) (float64, error) {
	z := 1.0
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func errorsFunc() {
	nn := ErrNegativeSqrt(-5)
	fmt.Println(nn.Error())
	fmt.Println(SqrtError(2))
	fmt.Println(SqrtError(-2))
}
