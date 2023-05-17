package main

import (
	"fractal/frac"
	"image/color"
)

var (
	vFractal = frac.Generator{
		UpScaler: func(pixel int8) frac.Shape {
			if pixel == 0 {
				return *frac.NewShape(3, 3)
			}

			if pixel == 1 {
				return frac.Shape{
					{0, 1, 0},
					{-1, 1, -1},
					{1, 0, 1},
				}
			}

			return frac.Shape{
				{-1, 0, -1},
				{1, -1, 1},
				{0, 1, 0},
			}
		},
		ShapePixelToColor: func(pixel int8) color.Color {
			switch pixel {
			case 1, -1:
				return colorDark
			default:
				return colorLight
			}
		},
	}
)
