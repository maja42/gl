package main

/*
	WEB example:

<html>
	<head>
		<meta name="viewport" content="width=device-width, initial-scale=1, user-scalable=no">
	</head>
	<body>
		<script src="motionblur.js" type="text/javascript"></script>
	</body>
</html>
*/

import (
	"encoding/binary"
	"fmt"
	"github.com/maja42/gl"
	"github.com/maja42/gl/glutil"
	"github.com/maja42/gl/render"
	"github.com/maja42/glfw"
	"golang.org/x/mobile/exp/f32"
	"log"
)

const vertexSource = `
#version 100
attribute vec2 position;
attribute vec3 color;
varying vec3 vColor;

void main(void) {
    gl_Position = vec4(position, 0.0, 1.0);
    vColor = color;
}
`

const fragmentSource = `
precision mediump float;
varying vec3 vColor;

void main(void) {
	gl_FragColor = vec4(vColor, 1);
}
`

func main() {
	renderThread := render.New()
	defer renderThread.Terminate()

	if err := glfw.Init(renderThread, gl.ContextWatcher); err != nil {
		log.Fatalln(err)
	}
	defer glfw.Terminate()

	gl.Init(renderThread)

	window, err := glfw.CreateWindow(1024, 768, "", nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	window.MakeContextCurrent()

	fmt.Printf("OpenGL: %s %s %s; %v samples.\n", gl.GetString(gl.VENDOR), gl.GetString(gl.RENDERER), gl.GetString(gl.VERSION), gl.GetInteger(gl.SAMPLES))
	fmt.Printf("GLSL: %s.\n", gl.GetString(gl.SHADING_LANGUAGE_VERSION))

	glfw.SwapInterval(1)
	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		window.SetShouldClose(true)
	})

	gl.ClearColor(0.1, 0.1, 0.1, 1)

	program, err := glutil.CreateProgram(vertexSource, fragmentSource)
	if err != nil {
		log.Fatalln(err)
	}
	defer gl.DeleteProgram(program)

	gl.ValidateProgram(program)
	if gl.GetProgrami(program, gl.VALIDATE_STATUS) != gl.TRUE {
		log.Fatalf("gl validate status: %s", gl.GetProgramInfoLog(program))
	}
	gl.UseProgram(program)

	vertexPositionAttrib := gl.GetAttribLocation(program, "position")
	gl.EnableVertexAttribArray(vertexPositionAttrib)
	vertexColorAttrib := gl.GetAttribLocation(program, "color")
	gl.EnableVertexAttribArray(vertexColorAttrib)

	triangleVertexPositionBuffer := gl.CreateBuffer()
	defer gl.DeleteBuffer(triangleVertexPositionBuffer)

	gl.BindBuffer(gl.ARRAY_BUFFER, triangleVertexPositionBuffer)
	vertices := f32.Bytes(binary.LittleEndian,
		-0.5, -0.3, 1, 0, 0,
		0.5, -0.3, 0, 1, 0,
		0, 0.5, 0, 0, 1,
	)
	gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.DYNAMIC_DRAW)

	if err := gl.GetError(); err != 0 {
		log.Fatalf("gl error: %v", err)
	}

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.VertexAttribPointer(vertexPositionAttrib, 2, gl.FLOAT, false, 20, 0)
		gl.VertexAttribPointer(vertexColorAttrib, 3, gl.FLOAT, false, 20, 8)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.SwapBuffers()
		glfw.PollEvents()
		renderThread.Sync()

		if err := gl.CheckError(); err != nil {
			log.Fatalln(err)
		}
	}
}
