package main

import (
	"bufio"
	"fmt"
	"image"
	"os"

	_ "image/png"
)

func loadImage(path string) ([]uint8, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bufio.NewReader(file))
	if err != nil {
		return nil, err
	}

	switch trueim := img.(type) {
	case *image.RGBA:
		return trueim.Pix, nil
	case *image.NRGBA:
		return trueim.Pix, nil
	}
	return nil, fmt.Errorf("unhandled image format")
}
