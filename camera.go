package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tadeuszjt/blobs/geom"
)

var (
	camera = struct {
		zoom     float64
		position geom.Vec2
	}{zoom: 1}

	mouse struct {
		windowPos, worldPos geom.Vec2
	}

	windowRect geom.Rect
	glRect     geom.Rect = geom.RectCentered(2, -2, geom.Vec2{})
)

func mat3WorldToGl() geom.Mat3 {
	cameraRect := geom.RectCentered(
		camera.zoom*windowRect.Width(),
		camera.zoom*windowRect.Height(),
		camera.position,
	)
	return geom.Mat3Camera2D(cameraRect, glRect)
}

func mat3WindowToWorld() geom.Mat3 {
	cameraRect := geom.RectCentered(
		camera.zoom*windowRect.Width(),
		camera.zoom*windowRect.Height(),
		camera.position,
	)
	return geom.Mat3Camera2D(windowRect, cameraRect)
}

func sizeCallback(w *glfw.Window, width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	windowRect = geom.RectOrigin(float64(width), float64(height))
}

func mouseCallback(w *glfw.Window, xpos, ypos float64) {
	pos := geom.Vec2{xpos, ypos}
	mouse.windowPos = pos
	mouse.worldPos = mat3WindowToWorld().TimesVec2(pos, 1).Vec2()
}

func mouseScrollCallback(w *glfw.Window, xOffset, yOffset float64) {
	oldMouseWorld := mouse.worldPos
	camera.zoom *= 1 - yOffset*0.04
	newMouseWorld := mat3WindowToWorld().TimesVec2(mouse.windowPos, 1).Vec2()
	camera.position.PlusEquals(oldMouseWorld.Minus(newMouseWorld))
}
