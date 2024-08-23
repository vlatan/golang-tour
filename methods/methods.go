package methods

import (
	"fmt"
	"math"
)

// Methods
func Method() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// Pointer receivers
func PointerRecievers() {
	v := Vertex{3, 4}
	v.scale(10)
	fmt.Println(v.Abs())
}

// Pointers and functions
func PointersFunctions() {
	v := Vertex{4, 3}
	scale(&v, 10)
	fmt.Println(Abs(v))
}

// Methods and pointer indirection
func PointerIndirection() {
	v := Vertex{3, 4}
	v.scale(2)
	scale(&v, 10)

	p := &Vertex{4, 3}
	p.scale(3)
	scale(p, 8)

	fmt.Println(v, p)

	v = Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	p = &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(Abs(*p))

	p = &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", p, p.Abs())
	p.scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", p, p.Abs())
}
