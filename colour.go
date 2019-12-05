package main

import (
	"image/color"
	"math/rand"
)

func randColour() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(254)),
		A: 255,
	}
}

func mutateColour(c color.RGBA) color.RGBA {
	bitChance := 400
	mutateBits := func(bits *uint8) {
		for i := uint(0); i < 8; i++ {
			if rand.Intn(bitChance) == 0 {
				*bits ^= 0x1 << i
			}
		}
	}

	mutateBits(&c.R)
	mutateBits(&c.G)
	mutateBits(&c.B)

	return c
}
