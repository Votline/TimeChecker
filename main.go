package main

import (
	"log"

	"runtime"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"

	"TimeCheck/shaders"
	"TimeCheck/digits"
)

const windowWidth = 250
const windowHeight = 100

func init() {
	runtime.LockOSThread()
}

func main(){
	if err := glfw.Init(); err != nil {
		log.Fatalln("GLFW init error. \nErr: ", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "TimeCheck", nil, nil)
	if err != nil {
		log.Fatalln("Create window error. \nErr: ", err)
	}
	window.MakeContextCurrent()
	
	vidMode := glfw.GetPrimaryMonitor().GetVideoMode()
	newPosX := int(float32(vidMode.Width) * (260.0 / float32(vidMode.Width)))
	newPosY := int(float32(vidMode.Height) * (1070.0 / float32(vidMode.Height)))
	window.SetPos(vidMode.Width-newPosX, vidMode.Height-newPosY)

	if err := gl.Init(); err != nil {
		log.Fatalln("OpenGL init error. \nErr: ", err)
	}
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	program := gl.CreateProgram()
	shaders.CompileAndAttachShaders(program)
	gl.LinkProgram(program)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	
	vertices1, linesQuan1 := digits.CreateVertexDigits(0, 0)
	vertices2, linesQuan2 := digits.CreateVertexDigits(6, 0.25)
	vertices3, linesQuan3 := digits.CreateVertexDigits(7, 0.4)
	vertices4, linesQuan4 := digits.CreateVertexDigits(9, 0.6)
	allVertices := append(vertices1, vertices2...)
	allVertices = append(allVertices, vertices3...)
	allVertices = append(allVertices, vertices4...)

	gl.BufferData(gl.ARRAY_BUFFER, len(allVertices)*4, gl.Ptr(allVertices), gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	gl.LineWidth(3.0)
	glfw.SwapInterval(1)
	gl.UseProgram(program)
	gl.ClearColor(0.0, 0.0, 0.0, 0.9)
	for !window.ShouldClose(){
		
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.DrawArrays(gl.LINE_STRIP, 0, linesQuan1)
		gl.DrawArrays(gl.LINE_STRIP, linesQuan1, linesQuan2)
		gl.DrawArrays(gl.LINE_STRIP, linesQuan1+linesQuan2, linesQuan3)
		gl.DrawArrays(gl.LINE_STRIP, linesQuan1+linesQuan2+linesQuan3, linesQuan4)

		if err := gl.GetError(); err != gl.NO_ERROR {
			log.Fatalln("OpenGL error. \nErr: ", err)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
