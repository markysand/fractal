package main

import (
	"fractal/frac"
	"image/color"
)

var (
	colorLight = color.Gray{255}
	colorDark  = color.Gray{100}
	colorPink  = color.RGBA{
		200, 100, 0, 255,
	}
	colorYellow = color.RGBA{
		100, 200, 0, 255,
	}
)

func main() {
	img := vFractal.CreateImage(4)
	frac.SaveImage("new-frac.png", img)
}
