package interfaces

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v, (%v years)", p.Name, p.Age)
}

// Stringers
func Stringer() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

type IPAddr [4]byte

// Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	var addr []string
	for _, b := range ip {
		addr = append(addr, strconv.Itoa(int(b)))
	}
	return strings.Join(addr, ".")
}

// Exercise: Stringers
func ExerciseStringer() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
