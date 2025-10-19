package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

func Setup(textC, backC []float32) (uint32, int32, int32) {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	pg := gl.CreateProgram()
	shaders := attachShaders(pg)

	gl.LinkProgram(pg)
	gl.UseProgram(pg)
	
	ofC := gl.GetUniformLocation(pg, gl.Str("tColor\x00"))
	ofP := gl.GetUniformLocation(pg, gl.Str("offsetPos\x00"))
	scale := gl.GetUniformLocation(pg, gl.Str("scale\x00"))
	detachShaders(pg, shaders)

	gl.ClearColor(backC[0], backC[1], backC[2], backC[3])
	gl.Uniform4f(ofC, textC[0], textC[1], textC[2], textC[3])

	return pg, ofP, scale
}

func CreateVAO(vtc []float32) uint32 {
	var vao, vbo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	gl.BindVertexArray(vao)
		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.BufferData(gl.ARRAY_BUFFER, len(vtc)*4, gl.Ptr(vtc),gl.STATIC_DRAW)

		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
		gl.EnableVertexAttribArray(0)
	gl.BindVertexArray(0)

	return vao
}

func ElemRender(pg uint32, ofP, sc int32, vao uint32, vtq int32, scale, offsetPos []float32) {
	gl.UseProgram(pg)

	gl.Uniform2f(sc, scale[0], scale[1])
	gl.Uniform2f(ofP, offsetPos[0], offsetPos[1])
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.LINE_STRIP, 0, vtq)
}
