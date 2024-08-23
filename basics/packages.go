package basics

import (
	"fmt"
	"math"
	"math/rand"
)

// Packages
func RandomNumber(num int) {
	fmt.Println("My favorite number is", rand.Intn(num))
}

// Imports
func SquareRoot(num float64) {
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(num))
}

// Exported names
func PrintPi() {
	fmt.Println(math.Pi)
}
