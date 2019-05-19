package main

import (
	"github.com/tadeuszjt/blobs/geom"
	"github.com/tadeuszjt/blobs/table"
	"image/color"
	"math/rand"
)


var (
	blobs = table.T{
		[]geom.Vec2{},  // position
		[]color.RGBA{}, // colour
		[]int{},        // age
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
	blobs = table.Append(blobs, table.T{
		pos,
		colour,
		0,
	})
}

func collides(v geom.Vec2) bool {
	for _, p := range blobs[0].([]geom.Vec2) {
		if p.Minus(v).Len2() < 16*16 {
			return true
		}
	}
	return false
}

func breedPosition(parent geom.Vec2) geom.Vec2 {
	dist := rand.Float64() * 20 + 16
	vec := geom.Vec2RandNormal().Scaled(dist)
	return parent.Plus(vec)
}

func update() {
	positions := blobs[0].([]geom.Vec2)
	colours := blobs[1].([]color.RGBA)
	ages := blobs[2].([]int)

	children := blobs.Slice(0, 0)

	// spawn children
	for i, position := range positions {
		if rand.Intn(40) == 0 {
			childPos := breedPosition(position)
			if arena.Contains(childPos) && !collides(childPos) {
				colour := mutateColour(colours[i])
				spawnBlob(childPos, colour)
			}
		}
	}

	// increase age
	for i := range ages {
		ages[i]++
	}

	// die
	blobs = table.Filter(blobs, func(col table.T) bool {
		return rand.Intn(1000 - col[2].(int)) != 0
	})

	blobs = table.Append(blobs, children)
}
