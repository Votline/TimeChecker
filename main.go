package main

import (
	"log"
	"time"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"TimeCheck/utils"
	"TimeCheck/shaders"
)

const windowWidth = 210
const windowHeight = 90

func init() {
	runtime.LockOSThread()
}


func main() {
	sw := utils.GetSW()
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
	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton,
		action glfw.Action, mod glfw.ModifierKey) {
		if button == glfw.MouseButtonLeft && action == glfw.Press {
			if currentBtn := utils.HoverOnBtns(w, utils.Btns); currentBtn != nil {
				switch currentBtn.Text {
				case "TIMER":
					if !sw.Mode {
						sw.Mode = true
						sw.Running = true
						sw.StartTime = time.Now()
						sw.SavedTime = 0
					} else if sw.Mode && sw.Running {
						sw.Running = false
						sw.SavedTime += time.Since(sw.StartTime)
					} else if sw.Mode {
						sw.Mode = false
					}
				case "ONTOP":
					currentState := window.GetAttrib(glfw.Floating)
					window.SetAttrib(glfw.Floating, 1-currentState)
				}
			}
		}
	})

	vidMode := glfw.GetPrimaryMonitor().GetVideoMode()
	window.SetPos(vidMode.Width-220, vidMode.Height-1075)

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
	for !window.ShouldClose() {

		allVertices, verticesQuan := utils.CreateTime(sw)
		utils.CreateButtons(&allVertices, &verticesQuan)
		gl.BufferData(gl.ARRAY_BUFFER, len(allVertices)*4, gl.Ptr(allVertices), gl.STATIC_DRAW)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		start := int32(0)
		for _, vertices := range verticesQuan {
			gl.DrawArrays(gl.LINE_STRIP, start, vertices)
			start += vertices
		}

		if err := gl.GetError(); err != gl.NO_ERROR {
			log.Println("OpenGL error. \nErr: ", err)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
