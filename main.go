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

var devilsFrac = frac.Generator{
	UpScaler: func(pixel int8) frac.Shape {
		switch pixel {
		case 1:
			return frac.Shape{
				{4, 0, 2},
				{1, 1, 1},
				{0, 1, 0},
			}
		case 2:
			return frac.Shape{
				{0, 2, 1},
				{2, 2, 0},
				{0, 2, 3},
			}
		case 3:
			return frac.Shape{
				{0, 3, 0},
				{3, 3, 3},
				{4, 0, 2},
			}
		case 4:
			return frac.Shape{
				{1, 4, 0},
				{0, 4, 4},
				{3, 4, 0},
			}
		default:
			return *frac.NewShape(3, 3)
		}

	},
	ShapePixelToColor: func(pixel int8) color.Color {
		if pixel == 0 {
			return colorLight
		}
		return colorDark
	},
}

func main() {
	img := devilsFrac.CreateImage(4)
	frac.SaveImage("new-frac.png", img)
}
