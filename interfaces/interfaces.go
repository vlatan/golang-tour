package interfaces

import (
	"fmt"
	"golang-tour/methods"
	"math"
)

// Interfaces
func Interfaces() {
	var a abser
	f := methods.MyFloat(-math.Sqrt2)
	v := methods.Vertex{X: 3, Y: 4}

	// The reciever is pointer but the method needs a value reciever.
	// Automatic pointer inderection (the pointer reciever is dereferenced).
	// The same as writing (*p).Abs()
	p := &f
	fmt.Println(p.Abs())

	// The reciever is value but the method needs a pointer.
	// Automatic pointer inderection (the value reciever is converted to pointer)
	// The same as writing (&v).Abs()
	fmt.Println(v.Abs())

	// "a" is of type "abser" interface,
	// which means its value ("f" in this case) must implement all the methods defined in this interface
	// "f" is of MyFloat type which indeed implements Abs() method as defined in "abser".
	// In short we say "MyFloat implements abser".
	a = f
	fmt.Println(a.Abs())

	// "a" is of type "abser" interface,
	// which means its value ("&v" in this case) must implement all the methods defined in this interface
	// "&v" is of *Vertex type which indeed implements Abs() method as defined in "abser".
	// In short we say "*Vertex implements abser".
	a = &v
	fmt.Println(a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement abser, meaning it will raise error.
	// Apparently interfaces care if their method's reciever is a pointer or value.
	// a = v
	// fmt.Println(a.Abs())
}

// Interfaces are implemented implicitly
func InterfaceIsImplicit() {
	var i I = &T{"hello"}
	i.M()
}

// Interface values
func InterfaceValues() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

// Interface values with nil underlying values
func InterfaceNilUnderlyingValue() {
	var i I  // i is I interface type
	var t *T // t is pointer to T type
	i = t    // assign t pointer to i
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

// Nil interface values
func NilIntrefaceValue() {
	var i I
	describe(i)
	// i.M() // this will panic because the interface is nil, has no underlying value

}

// The empty interface
func EmptyInterface() {
	var i interface{}
	describeEmptyInterface(i)

	i = 42
	describeEmptyInterface(i)

	i = "hello"
	describeEmptyInterface(i)
}
