package frac

import (
	"image"
	"image/png"
	"os"
)

func SaveImage(name string, img image.Image) error {
	f, err := os.Create(name)

	if err != nil {
		return (err)
	}

	defer f.Close()

	err = png.Encode(f, img)

	if err != nil {
		return (err)
	}

	return nil
}
