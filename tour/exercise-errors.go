package main

import (
	"fmt"
)
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(-2)
	}
	res := 1.0
	for i := 0; i < 10; i++ {
		res = res - (res*res-x)/2/res
	}
	return res, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
