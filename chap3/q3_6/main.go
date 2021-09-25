package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(px, py, mandelbrot(z))
		}
	}
	img = superSampling(img)
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, 255 - contrast*n, 255 - contrast*n, 255 - contrast*n}
		}
	}
	return color.RGBA{0xaa, 0xaa, 0, 0xff}
}

// https://github.com/relsa/exercise-gopl.io/blob/master/ch03/ex06/mandelblot.go
func superSampling(img *image.RGBA) *image.RGBA {
	rect := img.Bounds()
	width := rect.Max.X - rect.Min.X
	height := rect.Max.Y - rect.Min.Y
	out := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			out.Set(x, y, subPixel(img, x, y))
		}
	}
	return img
}

func subPixel(img *image.RGBA, x, y int) color.RGBA {
	pixels := []color.RGBA{
		img.RGBAAt(x, y),
		img.RGBAAt(x+1, y),
		img.RGBAAt(x, y+1),
		img.RGBAAt(x+1, y+1)}
	var r, g, b uint32
	for _, p := range pixels {
		r += uint32(p.R)
		g += uint32(p.G)
		b += uint32(p.B)
	}
	return color.RGBA{uint8(r / 4), uint8(g / 4), uint8(b / 4), 0xff}
}
