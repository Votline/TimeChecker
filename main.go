package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"TimeCheck/internal/ui"
	"TimeCheck/internal/render"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("GLFW init error. \nErr: ", err)
	}
	defer glfw.Terminate()

	if err := gl.Init(); err != nil {
		log.Fatalln("OpenGL init error. \nErr: ", err)
	}
	log.SetFlags(log.Lshortfile)

	win := ui.CreateWin()
	pg, ofL := render.Setup()
	sw := ui.CreateSW(pg, ofL)
	win.SetMouseButtonCallback(ui.BtnCallback(sw))

	glfw.SwapInterval(1)
	for !win.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		
		sw.Render(win)

		if err := gl.GetError(); err != gl.NO_ERROR {
			log.Println("OpenGL error. \nErr: ", err)
		}

		win.SwapBuffers()
		glfw.PollEvents()
	}
}
