package interfaces

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

// Images
func Img() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

type Image struct {
	Height, Width int
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Height, m.Width)
}

func (Image) At(x, y int) color.Color {
	c := uint8((x ^ y) * (x ^ y))
	return color.RGBA{c, c, 255, 255}
}

// ShowImage displays the image m
// when executed on the Go Playground.
// https://cs.opensource.google/go/x/tour/+/refs/tags/v0.1.0:pic/pic.go
func ShowImage(m image.Image) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	io.WriteString(w, "IMAGE:")
	b64 := base64.NewEncoder(base64.StdEncoding, w)
	err := (&png.Encoder{CompressionLevel: png.BestCompression}).Encode(b64, m)
	if err != nil {
		panic(err)
	}
	b64.Close()
	io.WriteString(w, "\n")
}

// Exercise: Images
func ExerciseImages() {
	m := Image{256, 256}
	ShowImage(m)
}
