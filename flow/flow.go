package flow

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// For loop
func ForLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

// While loop
func WhileLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// Forever loop
func ForeverLoop(seconds int64) {
	end := time.Now().Unix() + seconds

	for {
		now := time.Now().Unix()
		if now > end {
			fmt.Printf("Breaking the forever loop after %v seconds.\n", seconds)
			break
		}
	}
}

// If statement
func IfStatement(x float64) string {
	if x < 0 {
		return IfStatement(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// If with a short statement
func IfShort(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// // can't use v here, though
	return lim
}

// If and else
func IfElse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// Excercise Loops
func sqrt(x float64) float64 {
	z, delta, iter := x/2, 0.00000000001, 0
	for math.Abs(z*z-x) > delta {
		z -= (z*z - x) / (2 * z)
		iter += 1
	}
	// fmt.Println("Number of iterations:", iter)
	return z
}

func ExcerciseLoops(num float64) {
	fmt.Println(sqrt(num))
	fmt.Println(math.Sqrt(num))
}

// Switch
func Switch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows..
		fmt.Printf("%s\n", os)
	}
}

// Switch evaluation order
func SwitchEvaluationOrder() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// Switch with no condition.
func SwitchNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Defer
func Defer() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}

// Stacking defers
func DeferStacking() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
