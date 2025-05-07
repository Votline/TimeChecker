package main

import (
	"log"
	"time"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"

	"TimeCheck/primShapes"
	"TimeCheck/shaders"
	"TimeCheck/letters"
	"TimeCheck/digits"
)

const windowWidth = 210
const windowHeight = 90

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
	newPosX := int(float32(vidMode.Width) * (220.0 / float32(vidMode.Width)))
	newPosY := int(float32(vidMode.Height) * (1075.0 / float32(vidMode.Height)))
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
	
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	gl.LineWidth(3.0)
	glfw.SwapInterval(1)
	gl.UseProgram(program)
	gl.ClearColor(0.0, 0.0, 0.0, 0.9)
	for !window.ShouldClose(){
		
		allVertices, verticesQuan := createArraysForDraw()
		gl.BufferData(gl.ARRAY_BUFFER, len(allVertices)*4, gl.Ptr(allVertices), gl.STATIC_DRAW)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		start := int32(0)
		for _, vertices := range verticesQuan {
			gl.DrawArrays(gl.LINE_STRIP, start, vertices)
			start += vertices
		}
		
		if err := gl.GetError(); err != gl.NO_ERROR {
			log.Fatalln("OpenGL error. \nErr: ", err)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func createArraysForDraw() ([]float32, []int32){
	var allVertices []float32
	var verticesQuan []int32
	offsetD := float32(0.09)
	offsetL := float32(0.0)
	s := time.Now().Format("15:04:05") + "TIMER ONTOP"

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			num := int(ch - '0')
			vertices1, verticesQuan1 := digits.CreateVertexDigits(num, offsetD)
			allVertices = append(allVertices, vertices1...)
			verticesQuan = append(verticesQuan, verticesQuan1)
			offsetD += 0.2	
		} else if ch == ':'{
			vertices1, verticesQuan1 := digits.CreateVertexDigits(10, offsetD)
			vertices2, verticesQuan2 := digits.CreateVertexDigits(11, offsetD)
			allVertices = append(allVertices, vertices1...)
			verticesQuan = append(verticesQuan, verticesQuan1)
			allVertices = append(allVertices, vertices2...)
			verticesQuan = append(verticesQuan, verticesQuan2)
			offsetD += 0.1
		} else if ch >= 'A' && ch <= 'Z' {
			vertices1, verticesQuan1, width := letters.CreateVertexLetters(ch, offsetL)
			allVertices = append(allVertices, vertices1...)
			verticesQuan = append(verticesQuan, verticesQuan1)
			offsetL += width
		} else if ch == ' ' {
			offsetL += 0.15
		}
	}
	timerButton := primShapes.CreateRect(
		-0.85, -0.3, 
		-0.2, -0.8,
	)
	ontopButton := primShapes.CreateRect(
    -0.08, -0.3,
    0.75, -0.8,
  )
	allVertices = append(allVertices, timerButton...)
	allVertices = append(allVertices, ontopButton...)
	verticesQuan = append(verticesQuan, 5)
	verticesQuan = append(verticesQuan, 5)
	return allVertices, verticesQuan
}
