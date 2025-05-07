package utils

import (
	"fmt"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"

  "TimeCheck/digits"
  "TimeCheck/letters"
  "TimeCheck/primShapes"
)

const windowWidth = 210
const windowHeight = 90

type Btn struct {
  X1, Y1 float32
  X2, Y2 float32
  Text   string
}

type StopWatch struct {
  StartTime time.Time
  Running   bool
  Mode      bool
  SavedTime time.Duration
}

var Btns []Btn

func GetSW() *StopWatch {
	return &StopWatch{}
}

func CreateTime(sw *StopWatch) ([]float32, []int32) {
  var allVertices []float32
  var verticesQuan []int32
  offsetD := float32(0.11)
  s := time.Now().Format("15:04:05")

  if sw.Mode {
    if sw.Running {
      elapsed := time.Now().Sub(sw.StartTime)
      if elapsed >= time.Second {
        sw.SavedTime += elapsed
        sw.StartTime = time.Now()
      }
    }
    hours := int(sw.SavedTime.Hours())
    minutes := int(sw.SavedTime.Minutes()) % 60
    seconds := int(sw.SavedTime.Seconds()) % 60
    s = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
  }

  for _, ch := range s {
    if ch >= '0' && ch <= '9' {
      num := int(ch - '0')
      vertices1, verticesQuan1 := digits.CreateVertexDigits(num, offsetD)
      allVertices = append(allVertices, vertices1...)
      verticesQuan = append(verticesQuan, verticesQuan1)
      offsetD += 0.2
    } else if ch == ':' {
      vertices1, verticesQuan1 := digits.CreateVertexDigits(10, offsetD)
      vertices2, verticesQuan2 := digits.CreateVertexDigits(11, offsetD)
      allVertices = append(allVertices, vertices1...)
      verticesQuan = append(verticesQuan, verticesQuan1)
      allVertices = append(allVertices, vertices2...)
      verticesQuan = append(verticesQuan, verticesQuan2)
      offsetD += 0.1
    }
  }
  return allVertices, verticesQuan
}

func CreateButtons(allVertices *[]float32, verticesQuan *[]int32) {
  Btns = []Btn{}
  offsetL := float32(-0.02)
  s := "STOPW ONTOP"

  for _, ch := range s {
    if ch >= 'A' && ch <= 'Z' {
      vertices1, verticesQuan1, width := letters.CreateVertexLetters(ch, offsetL)
      *allVertices = append(*allVertices, vertices1...)
      *verticesQuan = append(*verticesQuan, verticesQuan1)
      offsetL += width
    } else if ch == ' ' {
      offsetL += 0.15
    }
  }
  timerButton := Btn{X1: -0.88, Y1: -0.3, X2: -0.08, Y2: -0.8, Text: "TIMER"}
  timerVertices := primShapes.CreateRect(
    timerButton.X1, timerButton.Y1,
    timerButton.X2, timerButton.Y2,
  )
  ontopButton := Btn{X1: 0.03, Y1: -0.3, X2: 0.85, Y2: -0.8, Text: "ONTOP"}
  ontopVertices := primShapes.CreateRect(
    ontopButton.X1, ontopButton.Y1,
    ontopButton.X2, ontopButton.Y2,
  )
  *allVertices = append(*allVertices, timerVertices...)
  *allVertices = append(*allVertices, ontopVertices...)
  *verticesQuan = append(*verticesQuan, 5)
  *verticesQuan = append(*verticesQuan, 5)
  Btns = append(Btns, timerButton)
  Btns = append(Btns, ontopButton)
}

func HoverOnBtns(w *glfw.Window, Btns []Btn) (hoverButton *Btn) {
  mX, mY := w.GetCursorPos()
  glX := float32(mX)/float32(windowWidth)*2 - 1
  glY := 1 - float32(mY)/float32(windowHeight)*2

  for _, Btn := range Btns {
    if glX >= Btn.X1 && glX <= Btn.X2 && glY <= Btn.Y1 && glY >= Btn.Y2 {
      return &Btn
    }
  }
  return nil
}
