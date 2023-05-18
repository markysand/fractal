package main

import (
	"fractal/frac"
	"image/color"
)

var (
	colorLight = color.Gray{0}
	colorDark  = color.Gray{10}
	colorPink  = color.RGBA{
		0, 200, 0, 255,
	}
	colorRed = color.RGBA{
		50, 200, 0, 255,
	}
	colorOrange = color.RGBA{
		100, 200, 0, 255,
	}
	colorYellow = color.RGBA{
		150, 200, 0, 255,
	}
	colorGreen = color.RGBA{
		200, 200, 25, 255,
	}
	colorBlue = color.RGBA{
		250, 200, 110, 255,
	}
)

var gen = frac.Generator{
	UpScaler: func(pixel int8) frac.Shape {
		switch pixel {
		case 1:
			return frac.Shape{
				{4, 0, 2},
				{1, 5, 1},
				{0, 1, 0},
			}
		case 2:
			return frac.Shape{
				{0, 3, 1},
				{2, 5, 0},
				{0, 2, 3},
			}
		case 3:
			return frac.Shape{
				{0, 3, 0},
				{3, 5, 3},
				{4, 0, 2},
			}
		case 4:
			return frac.Shape{
				{1, 4, 0},
				{0, 5, 4},
				{3, 4, 0},
			}
		case 5:
			return frac.Shape{
				{1, 2, 3},
				{4, 0, 4},
				{3, 2, 1},
			}
		default:
			return *frac.NewShape(3, 3)
		}

	},
	ShapePixelToColor: func(pixel int8) color.Color {
		switch pixel {
		case 1:
			return colorRed
		case 2:
			return colorOrange
		case 3:
			return colorYellow
		case 4:
			return colorGreen
		case 5:
			return colorBlue
		default:
			return colorLight
		}
	},
}

func main() {
	img := frac.Sierpinski.CreateImage(6)
	frac.SaveImage("new-frac.png", img)
}
