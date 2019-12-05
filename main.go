package main

import (
	"github.com/tadeuszjt/blobs/geom/geom32"
	"github.com/tadeuszjt/blobs/gfx"
	"image/color"
)

func vertexData(positions []geom.Vec2, ages []int, colours []color.RGBA) []float32 {
	slice := make([]float32, 0, 8*len(positions))
	for i, pos := range positions {
		grown := 1 - 1/float32(ages[i])
		scale := grown * 10

		verts := [4]geom.Vec2{
			geom.Vec2{pos.X - scale, pos.Y - scale},
			geom.Vec2{pos.X + scale, pos.Y - scale},
			geom.Vec2{pos.X + scale, pos.Y + scale},
			geom.Vec2{pos.X - scale, pos.Y + scale},
		}

		texcoords := [4]geom.Vec2{
			{0, 0},
			{1, 0},
			{1, 1},
			{0, 1},
		}

		for _, j := range [...]int{0, 1, 2, 0, 2, 3} {
			slice = append(
				slice,
				verts[j].X,
				verts[j].Y,
				texcoords[j].X,
				texcoords[j].Y,
				float32(colours[i].R)/255,
				float32(colours[i].G)/255,
				float32(colours[i].B)/255,
				float32(colours[i].A)/255,
			)
		}
	}

	return slice
}

var (
	texID     gfx.TexID
	frameRect geom.Rect
	mousePos  geom.Vec2
	camera    = struct {
		zoom float32
		pos  geom.Vec2
	}{zoom: 4, pos: geom.Vec2{}}
)

func camRect() geom.Rect {
	return geom.RectCentered(
		frameRect.Width()*camera.zoom,
		frameRect.Height()*camera.zoom,
		camera.pos,
	)
}

func mouseWorld() geom.Vec2 {
	dispToWorld := geom.Mat3Camera2D(frameRect, camRect())
	return dispToWorld.TimesVec2(mousePos, 1).Vec2()
}

func resize(width, height int) {
	frameRect = geom.RectOrigin(float32(width), float32(height))
}

func mouse(w *gfx.Win, event gfx.MouseEvent) {
	switch ev := event.(type) {
	case gfx.MouseScroll:
		oldMouseWorld := mouseWorld()
		camera.zoom *= 1 + 0.04*(-ev.Dy)
		newMouseWorld := mouseWorld()
		camera.pos.PlusEquals(oldMouseWorld.Minus(newMouseWorld))

	case gfx.MouseMove:
		mousePos = ev.Position

	default:
	}
}

func setup(w *gfx.Win) error {
	var err error
	texID, err = w.LoadTexture("Circle.png")
	return err
}

func draw(w *gfx.WinDraw) {
	worldToDisplay := geom.Mat3Camera2D(camRect(), frameRect)
	w.SetMatrix(worldToDisplay)

	update()
	data := vertexData(blobs.position, blobs.age, blobs.colour)
	w.DrawVertexData(data, &texID)
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		SetupFunc:  setup,
		DrawFunc:   draw,
		MouseFunc:  mouse,
		ResizeFunc: resize,
		Title:      "Blobs",
	})
}
