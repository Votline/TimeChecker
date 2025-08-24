package ui

import (
	"log"
	//"time"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const winW = 210
const winH = 90

type element interface {
	create() element
	
	getVtq() int32
	getVtc() []float32

	setRune(rune)
}
type digit struct {
	vtc []float32
	vtq int32
	rn rune
}
type letter struct {
	vtc []float32
	vtq int32
	rn rune
}
type btn struct {
	Pos [4]float32
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

//	win.SetMouseButtonCallback(btnCallback(sw))

	return win
}


/*func btnCallback(sw *StopWatch) func(w *glfw.Window, btn glfw.MouseButton, act glfw.Action, mod glfw.ModifierKey) {
	return func(w *glfw.Window, btn glfw.MouseButton, act glfw.Action, mod glfw.ModifierKey) {
		if btn == glfw.MouseButtonLeft && act == glfw.Press {
			if currBtn := utils.HoverOnBtns(w, utils.Btns); currBtn != nil {
				switch currBtn.Text {
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
					winState := w.GetAttrib(glfw.Floating)
					w.SetAttrib(glfw.Floating, 1-winState)
				}
			}
		}
	}
}
*/
