package ui

import (
	"fmt"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"

	"TimeChecker/internal/render"
)

type elemMesh struct {
	vao uint32
	vtq int32
}

type StopWatch struct {
	pg uint32
	ofP int32
	scale int32

	digs []elemMesh
	btns map[*btn]elemMesh
	lets map[rune]elemMesh

	StartTime time.Time
	Running bool
	SWmode bool
	SavedTime time.Duration
}

func createElement[T any, PT interface{*T; element}](chars []rune) []elemMesh {
	result := make([]elemMesh, len(chars))
	for i, ch := range chars {
		el := PT(new(T))
		el.setData(ch)
		created := el.create()
		vao := render.CreateVAO(created.getVtc())
		result[i] = elemMesh{vao: vao, vtq: created.getVtq()}
	}
	return result
}

func createBtn(positions [][4]float32, texts []string) map[*btn]elemMesh {
	result := make(map[*btn]elemMesh)
	for i, pos := range positions {
		b := &btn{}
		b.setPos(pos)
		b.setText(texts[i])
		el := b.create()
		vao := render.CreateVAO(el.getVtc())
		result[b] = elemMesh{vao: vao, vtq: el.getVtq()}
	}
	return result
}

func CreateSW(pg uint32, ofP int32, scale int32) *StopWatch {
	lets := make(map[rune]elemMesh)
	letsRunes := []rune{'S', 'T', 'O', 'P', 'W', 'N'}
	tempLets := createElement[letter](letsRunes)
	
	for i, ch := range letsRunes {
		lets[ch] = tempLets[i]
	}
	pos1 := [4]float32{-0.88, -0.3, -0.04, -0.8}
	pos2 := [4]float32{0.03, -0.3, 0.88, -0.8}
	texts := []string{"STOPW", "ONTOP"}
	positions := [][4]float32{pos1, pos2}
	btns := createBtn(positions, texts)
	
	digs := createElement[digit]([]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', ':'})
	
	return &StopWatch{pg: pg, ofP: ofP, scale: scale, digs: digs, btns: btns, lets: lets}
}

func (sw *StopWatch) Render(win *glfw.Window) {
	sw.renderDate()
	sw.renderTime()
	sw.renderBtns()

	for btn, _ := range sw.btns {
		btn.Hover(win, 210, 90)
	}
}

func (sw *StopWatch) renderDigit(needed string, sc, offset []float32) {
	for _, ch := range needed {
		if ch == ':' || ch == '/' {
			render.ElemRender(sw.pg, sw.ofP, sw.scale,
				sw.digs[10].vao, sw.digs[10].vtq, sc, offset)
			render.ElemRender(sw.pg, sw.ofP, sw.scale,
				sw.digs[11].vao, sw.digs[11].vtq, sc, offset)
			offset[0] += 0.09
			continue
		}
		dgt := int(ch - '0')
		render.ElemRender(sw.pg, sw.ofP, sw.scale,
			sw.digs[dgt].vao, sw.digs[dgt].vtq, sc, offset)
		offset[0] += 0.15
	}
}

func (sw *StopWatch) renderDate() {
	date := time.Now().Format("02/01")
	sw.renderDigit(date, []float32{0.5, 0.5}, []float32{-0.05, 0.6})
}

func (sw *StopWatch) renderTime() {
	var neededTime string
	currTime := time.Now().Format("15:04:05")
	if sw.SWmode {
		if sw.Running {
			elapsed := time.Now().Sub(sw.StartTime)
			if elapsed >= time.Second {
				sw.SavedTime += elapsed
				sw.StartTime = time.Now()
			}
		}
		hours := int(sw.SavedTime.Hours())
		minute := int(sw.SavedTime.Minutes()) % 60
		seconds := int(sw.SavedTime.Seconds())% 60
		neededTime = fmt.Sprintf("%02d:%02d:%02d", hours, minute, seconds)
	} else {
		neededTime = currTime
	}

	sw.renderDigit(neededTime, []float32{1.0, 1.0}, []float32{0.0, 0.0})
}

func (sw *StopWatch) renderBtns() {
	spaceIdx := int32(0)
	offset := []float32{0.0, 0.0}
	for _, ch := range "STOPW ON TOP" {
		if ch == ' ' {
			if spaceIdx > 0 {offset[0]+=0.03} else {offset[0] += 0.15}
			spaceIdx++
			continue
		}
		if mesh, exists := sw.lets[ch]; exists {
			render.ElemRender(sw.pg, sw.ofP, sw.scale,
				mesh.vao, mesh.vtq, []float32{1.0, 1.0}, offset)
			offset[0] += 0.15
		}
	}
	for _, v := range sw.btns {
		render.ElemRender(sw.pg, sw.ofP, sw.scale,
			v.vao, v.vtq, []float32{1.0, 1.0}, []float32{0.0, 0.0})
	}
}
