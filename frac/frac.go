package frac

import (
	"image"
	"image/color"
)

const InitialPixelValue = 1

type Shape [][]int8

func (s Shape) DimX() int {
	if len(s) > 0 {
		return len(s[0])
	}

	return 0
}

func (s Shape) DimY() int {
	return len(s)
}

func NewShape(x, y int) *Shape {
	s := make(Shape, y)

	for xPos := 0; xPos < y; xPos++ {
		s[xPos] = make([]int8, x)
	}

	return &s
}

type Generator struct {
	// UpScaler will return a shape, equal in dimensions to the baseShape
	UpScaler          func(pixel int8) Shape
	ShapePixelToColor func(pixel int8) color.Color
}

func (g Generator) CreateImage(iterations int, pixelSize int) image.Image {
	var (
		currentShape = g.UpScaler(InitialPixelValue)
		nextShape    Shape
		xMul         = currentShape.DimX()
		yMul         = currentShape.DimY()
	)

	for l := 0; l < iterations; l++ {

		nextShape = *NewShape(currentShape.DimX()*xMul, currentShape.DimY()*yMul)

		for y := 0; y < currentShape.DimY(); y++ {
			for x := 0; x < currentShape.DimY(); x++ {
				fillShape := g.UpScaler(currentShape[y][x])

				for yD := 0; yD < yMul; yD++ {
					for xD := 0; xD < xMul; xD++ {
						nextShape[y*yMul+yD][x*xMul+xD] = fillShape[yD][xD]
					}
				}
			}
		}

		currentShape = nextShape
	}

	if pixelSize == 0 {
		pixelSize = 1
	}

	img := image.NewRGBA(image.Rect(0, 0, currentShape.DimX()*pixelSize, currentShape.DimY()*pixelSize))

	for y := 0; y < currentShape.DimY(); y++ {
		for x := 0; x < currentShape.DimY(); x++ {
			for yD := 0; yD < pixelSize; yD++ {
				for xD := 0; xD < pixelSize; xD++ {
					img.Set(x*pixelSize+xD, y*pixelSize+yD, g.ShapePixelToColor(currentShape[y][x]))
				}
			}
		}
	}

	return img
}
