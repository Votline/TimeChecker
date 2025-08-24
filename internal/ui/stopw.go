package ui

import (
	"time"

	"TimeCheck/internal/render"
)

type elemMesh struct {
	vao uint32
	vtq int32
}

type StopWatch struct {
	pg uint32
	ofL int32

	digs []elemMesh
	lets map[rune]elemMesh

	StartTime time.Time
	Running bool
	Start bool
	SavedTime time.Duration
}

func createElement[T any, PT interface{*T; element}](chars []rune) []elemMesh {
	result := make([]elemMesh, len(chars))
	for i, ch := range chars {
		el := PT(new(T))
		el.setRune(ch)
		created := el.create()
		vao := render.CreateVAO(created.getVtc())
		result[i] = elemMesh{vao: vao, vtq: created.getVtq()}
	}
	return result
}

func CreateSW(pg uint32, ofl int32) *StopWatch {
	lets := make(map[rune]elemMesh)
	letsRunes := []rune{'S', 'T', 'O', 'P', 'W', 'N'}
	digs := createElement[digit]([]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', ':'})
	tempLets := createElement[letter](letsRunes)
	
	for i, ch := range letsRunes {
		lets[ch] = tempLets[i]
	}
	
	return &StopWatch{pg: pg, ofL: ofl, digs: digs, lets: lets}
}

func (sw *StopWatch) Render() {
	sw.renderTime()
	sw.renderBtns()
}

func (sw *StopWatch) renderTime() {
	offset := float32(0.0)
	currTime := time.Now().Format("15:04:05")
	for _, ch := range currTime {
		if ch == ':' {
			render.ElemRender(sw.pg, sw.ofL,
				sw.digs[10].vao, sw.digs[10].vtq, offset)
			render.ElemRender(sw.pg, sw.ofL,
				sw.digs[11].vao, sw.digs[11].vtq, offset)
			offset += 0.09
			continue
		}
		dgt := int(ch - '0')
		render.ElemRender(sw.pg, sw.ofL,
			sw.digs[dgt].vao, sw.digs[dgt].vtq, offset)
		offset += 0.15
	}
}

func (sw *StopWatch) renderBtns() {
	spaceIdx := int32(0)
	offset := float32(0.0)
	for _, ch := range "STOPW ON TOP" {
		if ch == ' ' {
			if spaceIdx > 0 {offset+=0.03} else {offset += 0.15}
			spaceIdx++
			continue
		}
		if mesh, exists := sw.lets[ch]; exists {
			render.ElemRender(sw.pg, sw.ofL,
				mesh.vao, mesh.vtq, offset)
			offset += 0.15
		}
	}
}
