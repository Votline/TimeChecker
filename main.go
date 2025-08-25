package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"TimeChecker/internal/ui"
	"TimeChecker/internal/render"
	"TimeChecker/internal/config"
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

	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("Config getting error: %v", err)
	}

	win := ui.CreateWin(cfg.WinW, cfg.WinH, cfg.AlX, cfg.AlY)
	pg, ofL := render.Setup(cfg.TextC, cfg.BackC)
	sw := ui.CreateSW(pg, ofL)
	win.SetMouseButtonCallback(ui.BtnCallback(sw, cfg.WinW, cfg.WinH))

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
