package frac

import (
	"image/color"
)

var (
	VFractal = Generator{
		UpScaler: func(pixel int8) Shape {
			if pixel == 0 {
				return *NewShape(3, 3)
			}

			if pixel == 1 {
				return Shape{
					{0, -1, 0},
					{-1, 1, -1},
					{1, 0, 1},
				}
			}

			return Shape{
				{-1, 0, -1},
				{1, -1, 1},
				{0, 1, 0},
			}
		},
		ShapePixelToColor: func(pixel int8) color.Color {
			switch pixel {
			case 1, -1:
				return color.Gray{0}
			default:
				return color.RGBA{100, 100, 100, 0}
			}
		},
	}

	Sierpinski = Generator{
		UpScaler: func(pixel int8) Shape {
			if pixel == 1 {
				return Shape{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				}
			}

			return *NewShape(3, 3)
		},
		ShapePixelToColor: func(pixel int8) color.Color {
			if pixel == 1 {
				return color.Gray{0}
			}
			return color.Gray{255}
		},
	}
)
