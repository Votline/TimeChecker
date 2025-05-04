package main

import (
	"log"

	"runtime"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"

	"TimeCheck/shaders"
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
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
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
	program := gl.CreateProgram()
	shaders.CompileAndAttachShaders(program)
	gl.LinkProgram(program)

	glfw.SwapInterval(1)
	gl.ClearColor(0.0, 0.0, 0.0, 0.5)
	for !window.ShouldClose(){
		gl.Clear(gl.COLOR_BUFFER_BIT)
		
		if err := gl.GetError(); err != gl.NO_ERROR {
			log.Fatalln("OpenGL error. \nErr: ", err)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
