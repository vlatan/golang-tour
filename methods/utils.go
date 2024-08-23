package methods

import (
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

// method on a non-struct type
// type must be defined in the same package as the method
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Method on a struct as reciever.
// The reciever can be a pointer or value (it doesn't matter).
// In this case "pointer indirection" will take place which means
// if the struct is assigned as value but needed by the method as pointer
// it vill be automatically passed as pointer, and vice-verca
// if it's assigned as pointer but needed by the method as value
// it will be automaticallt dereferenced and thus passes as value.
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods are just functions
// function that recieves a struct value
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// method on a struct pointer as reciever
func (v *Vertex) scale(f float64) {
	v.X *= f
	v.Y *= f
}

// function that recieves a pointer to struct
func scale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}
