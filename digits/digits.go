package digits

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

func CreateVertexDigits(number int) int32 {
	vertices, linesQuantity := createVertices(number)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)
	
	return linesQuantity
}

func createVertices(number int) ([]float32, int32) {
	switch number{
	case 1:
		return []float32 {	
			-0.9, 0.0, 0.0,
			-0.7, 0.3, 0.0,
			-0.7, -0.3, 0.0,
		}, 3
	case 2:
		return []float32 {
			-0.9, 0.3, 0.0,
			-0.8, 0.4, 0.0,
			-0.7, 0.3, 0.0,
		
			-0.9, -0.3, 0.0,
			-0.6, -0.3, 0.0,
		}, 5
	case 3:
    	return []float32 {
				-0.9, 0.6, 0.0,
				-0.8, 0.7, 0.0,
				-0.7, 0.7, 0.0,
				-0.5, 0.6, 0.0,
			
				-0.5, 0.4, 0.0,
				-0.5, -0.2, 0.0,
				-0.7, -0.5, 0.0,
				-0.9, -0.7, 0.0,
			}, 8
	}
	return []float32 {
		0.0, 0.0, 0.0,
		0.0, 0.0, 0.0,
		0.0, 0.0, 0.0,
	}, 3 
}
