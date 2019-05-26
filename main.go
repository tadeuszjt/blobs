package main

import (
	"image/color"
	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tadeuszjt/blobs/geom"
	_ "net/http/pprof"
	"net/http"
	"log"
)

func init() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Mat3GeomToMgl32(m geom.Mat3) mgl32.Mat3 {
	return mgl32.Mat3{
		float32(m[0]), float32(m[3]), float32(m[6]),
		float32(m[1]), float32(m[4]), float32(m[7]),
		float32(m[2]), float32(m[5]), float32(m[8]),
	}
}

func blobsWindow() (*glfw.Window, error) {
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	return glfw.CreateWindow(640, 480, "Blobs", nil, nil)
}

func vertexData(positions []geom.Vec2, ages []int, colours []color.RGBA) []float32 {
	slice := make([]float32, 0, 12*len(positions))
	for i, pos := range positions {
		grown := 1 - 1/float64(ages[i])
		scale := grown * 10


		verts := [4]geom.Vec2{
			geom.Vec2{pos.X - scale, pos.Y -scale},
			geom.Vec2{pos.X + scale, pos.Y -scale},
			geom.Vec2{pos.X + scale, pos.Y +scale},
			geom.Vec2{pos.X - scale, pos.Y +scale},
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
				float32(verts[j].X),
				float32(verts[j].Y),
				float32(texcoords[j].X),
				float32(texcoords[j].Y),
				float32(colours[i].R),
				float32(colours[i].G),
				float32(colours[i].B),
				float32(colours[i].A),
			)
		}
	}

	return slice
}

func run() {
	defer func() {
		mainthread.Call(func() {
			glfw.Terminate()
		})
	}()

	mainthread.Call(func() {
		glhf.Init()

		glfw.Init()
		win, err := blobsWindow()
		check(err)
		win.MakeContextCurrent()
		win.SetFramebufferSizeCallback(sizeCallback)
		win.SetCursorPosCallback(mouseCallback)
		win.SetScrollCallback(mouseScrollCallback)

		gl.Enable(gl.BLEND)
		glhf.BlendFunc(glhf.SrcAlpha, glhf.OneMinusSrcAlpha)

		vertexFmt := glhf.AttrFormat{
			{Name: "position", Type: glhf.Vec2},
			{Name: "texcoord", Type: glhf.Vec2},
			{Name: "colour", Type: glhf.Vec4},
		}
		uniformFmt := glhf.AttrFormat{{Name: "matrix", Type: glhf.Mat3}}
		shader, err := glhf.NewShader(vertexFmt, uniformFmt, vertexShader, fragmentShader)
		check(err)

		slice := glhf.MakeVertexSlice(shader, 0, 0)

		pixels, err := loadImage("Circle.png")
		check(err)
		texture := glhf.NewTexture(640, 640, true, pixels)
		texture.Begin()
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
		gl.GenerateMipmap(gl.TEXTURE_2D)
		texture.End()

		for !win.ShouldClose() {
			update()

			glfw.PollEvents()

			glhf.Clear(1, 1, 1, 1)
			shader.Begin()
			shader.SetUniformAttr(
				0,
				Mat3GeomToMgl32(mat3WorldToGl()),
			)

			texture.Begin()
			data := vertexData(blobs[0].([]geom.Vec2), blobs[2].([]int), blobs[1].([]color.RGBA))
			slice.Begin()
			slice.SetLen(len(data) / 8)
			slice.SetVertexData(data)
			slice.Draw()
			slice.End()
			texture.End()
			shader.End()

			win.SwapBuffers()
		}
	})
}

func main() {
	mainthread.Run(run)
}

var vertexShader = `
#version 330 core
uniform mat3 matrix;
in vec2 position;
in vec2 texcoord;
in vec4 colour;
out vec2 Texcoord;
out vec4 Colour;
void main() {
	gl_Position = vec4(matrix * vec3(position, 1), 1.0);
	Texcoord = texcoord;
	Colour = colour;
}
`
var fragmentShader = `
#version 330 core
uniform sampler2D tex;
in vec2 Texcoord;
in vec4 Colour;
out vec4 outColor;
void main() {
	outColor = (Colour / 255) * texture(tex, Texcoord);
}
`
