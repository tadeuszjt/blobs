package main

import (
	"github.com/tadeuszjt/blobs/geom/geom32"
	"image/color"
	"math/rand"
)

var (
	blobs struct {
		position []geom.Vec2
		colour   []color.RGBA
		age      []int
	}
	arena geom.Rect = geom.RectCentered(2000, 2000, geom.Vec2{})
)

func init() {
	for i := 0; i < 10; i++ {
		position := geom.Vec2Rand(arena)
		if !collides(position) {
			spawnBlob(position, randColour())
		}
	}
}

func spawnBlob(pos geom.Vec2, colour color.RGBA) {
	blobs.position = append(blobs.position, pos)
	blobs.colour = append(blobs.colour, colour)
	blobs.age = append(blobs.age, 0)
}

func collides(v geom.Vec2) bool {
	for _, p := range blobs.position {
		if p.Minus(v).Len2() < 16*16 {
			return true
		}
	}
	return false
}

func breedPosition(parent geom.Vec2) geom.Vec2 {
	dist := rand.Float32()*20 + 16
	vec := geom.Vec2RandNormal().ScaledBy(dist)
	return parent.Plus(vec)
}

func update() {
	// spawn children
	for i, position := range blobs.position {
		if rand.Intn(80) == 0 {
			childPos := breedPosition(position)
			if arena.Contains(childPos) && !collides(childPos) {
				colour := mutateColour(blobs.colour[i])
				spawnBlob(childPos, colour)
			}
		}
	}

	// increase age
	for i := range blobs.age {
		blobs.age[i]++
	}

	// die
	for i := 0; i < len(blobs.position); i++ {
		if rand.Intn(800) == 0 {
			end := len(blobs.position) - 1
			if i < end {
				blobs.position[i] = blobs.position[end]
				blobs.colour[i] = blobs.colour[end]
				blobs.age[i] = blobs.age[end]
			}

			blobs.position = blobs.position[:end]
			blobs.colour = blobs.colour[:end]
			blobs.age = blobs.age[:end]
		}
	}
}
