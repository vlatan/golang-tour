package interfaces

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Reader() {
	// strings.NewReader() returns a pointer to strings.Reader{} struct
	// with values: s string, i int64 as current reading index,
	// and prevRune int as index of previous rune or < 0
	r := strings.NewReader("Hello, Reader!")

	// create slice with len 8 that accepts bytes
	b := make([]byte, 8)

	// infinite loop
	for {
		// If a struct implements a Read() method than that struct implements io.Reader interface
		// So strings.Reader implements io.Reader interface
		// Read(b) populates the given byte slice with data and returns the number of bytes populated
		// and an error value. It returns an io.EOF error when the stream ends.
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:%v] = %q\n", n, b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct{}

// Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func Validate(r io.Reader) {
	b := make([]byte, 1024, 2048)
	i, o := 0, 0
	for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
		n, err := r.Read(b)
		for i, v := range b[:n] {
			if v != 'A' {
				fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
				return
			}
		}
		o += n
		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			return
		}
	}
	if o == 0 {
		fmt.Fprintf(os.Stderr, "read zero bytes after %d Read calls\n", i)
		return
	}
	fmt.Println("OK!")
}

// Exercise: Readers
func ExerciseReaders() {
	Validate(MyReader{})
}

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

func (ro rot13Reader) Read(b []byte) (n int, err error) {
	n, err = ro.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return
}

// Exercise: rot13Reader
func ExerciseRot1Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
