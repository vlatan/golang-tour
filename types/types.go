package types

import (
	"fmt"
	"math"
	"strings"
)

// Pointers
func Pointers() {
	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// Structs
func Structs() {
	type Vertex struct {
		X int
		Y int
	}
	fmt.Println(Vertex{1, 2})
}

// Struct Fields
func StructFields() {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	v.X = 4

	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(v)
}

// Pointers to structs
func StructPointers() {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

// Struct Literals
func StructLiterals() {
	type Vertex struct {
		X, Y int
	}

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, p, v2, v3)
}

// Arrays
func Array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// Slices
func Slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}

// Slices are like references to arrays
func SliceReference() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// Slice literals
func SliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// Slice defaults
func SliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Slice length and capacity
func SliceLengthCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

// Nil slices
func NilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	fmt.Println(s == nil)
}

func printMadeSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// Creating a slice with make
func MakeSlice() {
	a := make([]int, 5)
	printMadeSlice("a", a)

	b := make([]int, 0, 5)
	printMadeSlice("b", b)

	c := b[:2]
	printMadeSlice("c", c)

	d := c[2:5]
	printMadeSlice("d", d)
}

// Slices of slices
func SlicesOfSlices() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// Appending to a slice
func AppendToSlice() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	// We can also append a slice to the slice
	// 'append' is a variadic function which means it can accept variable number of arguments
	// As I understand the slice passed as an argument is unpacked and its elements are appended
	s = append(s, []int{1, 2}...)
	printSlice(s)
}

// Range
func Range() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// The range form of the for loop iterates over a slice or map.
	// giving the index (i) and a copy of an element at that index (v)
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 10)

	// If you only want the index, you can omit the second variable.
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// You can skip the index or value by assigning to _.
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// Exercise: Slices
// https://go.dev/tour/moretypes/18
func ExerciseSlices(dx, dy int) {
	result := make([][]uint8, dy)
	for x := range result {
		result[x] = make([]uint8, dx)
		for y := range result[x] {
			result[x][y] = uint8((x ^ y) * (x ^ y))
		}
	}
	fmt.Println(result)
}

// Maps
func Map() {
	type Vertex struct {
		Lat, Long float64
	}

	m := make(map[string]Vertex)

	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}

// Map literals
func MapLiterals() {
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

// Mutating Maps
func MutatingMaps() {
	m := make(map[string]int)
	fmt.Println(m)

	m["Answer"] = 42
	fmt.Println(m)
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println(m)
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println(m)
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// Exercise: Maps
func ExerciseMaps(s string) map[string]int {
	result := make(map[string]int)
	for _, word := range strings.Fields(s) {
		result[word] += 1
	}
	fmt.Println(result)
	return result
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Function values
func FunctionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Function closures
func FunctionClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := []int{}
	return func() int {
		num := 0
		if len(fib) == 1 {
			num = 1
		} else if len(fib) > 1 {
			num = fib[len(fib)-1] + fib[len(fib)-2]
		}
		fib = append(fib, num)
		return num
	}
}

// Exercise: Fibonacci closure
func ExcerciseClosure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
