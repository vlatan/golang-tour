package basics

// Functions
func Add(x int, y int) int {
	return x + y
}

// Functions continued
func Addition(x, y int) int {
	return x + y
}

// Multiple results
func Swap(x, y string) (string, string) {
	return y, x
}

// Named return values
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
