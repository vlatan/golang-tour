package interfaces

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

// Errors
func Error() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z, delta := x/2, 0.00000000000001
	for math.Abs(z*z-x) > delta {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

// Exercise: Errors
func ExerciseErrors() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
