package ui

import (
	"log"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const winW = 210
const winH = 90

type element interface {
	create() element
	
	getVtq() int32
	getVtc() []float32

	setData(rune)
}
type digit struct {
	rn rune
	vtq int32
	vtc []float32
}
type letter struct {
	rn rune
	vtq int32
	vtc []float32
}
type btn struct {
	vtq int32
	vtc []float32
	pos [4]float32
	Text string
}

func CreateWin() *glfw.Window {
	glfw.WindowHint(glfw.RefreshRate, 60)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.DoubleBuffer, glfw.True)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)

	win, err := glfw.CreateWindow(winW, winH, "TimeChecker", nil, nil)
	if err != nil {
		log.Fatalf("Create window error: %v", err)
	}

	win.MakeContextCurrent()
	win.SetAttrib(glfw.Floating, 1)
	
	vidMode := glfw.GetPrimaryMonitor().GetVideoMode()
	win.SetPos(vidMode.Width-220, vidMode.Height-1075)

	return win
}

func BtnCallback(sw *StopWatch) func(w *glfw.Window, mBtn glfw.MouseButton, act glfw.Action, mod glfw.ModifierKey) {
	return func(w *glfw.Window, mBtn glfw.MouseButton, act glfw.Action, mod glfw.ModifierKey) {
		if mBtn == glfw.MouseButtonLeft && act == glfw.Press {
			for btn, _ := range sw.btns {
				if btn.Hover(w, winW, winH) {
					switch btn.Text {
					case "STOPW":
						if !sw.SWmode {
							sw.SWmode = true
							sw.Running = true
							sw.StartTime = time.Now()
							sw.SavedTime = 0
						} else if sw.SWmode && sw.Running {
							sw.Running = false
							sw.SavedTime += time.Since(sw.StartTime)
						} else if sw.SWmode {
							sw.SWmode = false
						}
					case "ONTOP":
						winState := w.GetAttrib(glfw.Floating)
						w.SetAttrib(glfw.Floating, 1-winState)
					}
				}
			}
		}
	}
}
