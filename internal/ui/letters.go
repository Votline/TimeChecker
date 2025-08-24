package ui

const (
	letterMinX  = -0.8
	letterMaxX  = -0.7

	letterMaxY = -0.4
	letterMinY = -0.7
)

func (l *letter) create() element {
	switch l.rn {
	case 'S':
		return &letter{
			vtc: []float32{
				letterMaxX, letterMaxY, 0.0,
				letterMinX, (letterMaxY + letterMinY) / 2, 0.0,
				letterMaxX, letterMinY, 0.0,
				letterMinX, letterMinY, 0.0},
			vtq: int32(4)}
	case 'W':
		return &letter{
			vtc: []float32{
				letterMinX, letterMaxY, 0.0,
				(letterMinX + letterMaxX) * 0.55, letterMinY, 0.0,
				(letterMinX + letterMaxX) * 0.5, letterMaxY*1.2, 0.0,
				(letterMinX + letterMaxX) * 0.45, letterMinY, 0.0,
				letterMaxX, letterMaxY, 0.0},
			vtq: int32(5)}
	case 'T':
		return &letter{
			vtc: []float32{
				letterMinX - 0.02, letterMaxY, 0.0,
				letterMaxX + 0.02, letterMaxY, 0.0,

				(letterMinX + letterMaxX) / 2, letterMaxY, 0.0,
				(letterMinX + letterMaxX) / 2, letterMinY, 0.0,

				(letterMinX + letterMaxX) / 2, letterMaxY, 0.0},
			vtq: int32(5)}
	case 'O':
		return &letter{
			vtc: []float32{
				letterMinX, letterMaxY, 0.0,
				letterMinX, letterMinY, 0.0,

				letterMaxX, letterMinY, 0.0,
				letterMaxX, letterMaxY, 0.0,

				letterMinX, letterMaxY, 0.0},
			vtq: int32(5)}
	case 'N':
		return &letter{
			vtc: []float32{
				letterMinX, letterMinY, 0.0,
				letterMinX, letterMaxY, 0.0,

				letterMaxX, letterMinY, 0.0,
				letterMaxX, letterMaxY, 0.0},
			vtq: int32(4)}
	case 'P':
		return &letter{
			vtc: []float32{
				letterMinX, letterMinY, 0.0,
				letterMinX, letterMaxY, 0.0,

				letterMaxX, letterMaxY, 0.0,
				letterMaxX, (letterMaxY + letterMinY) / 2, 0.0,
				letterMinX, (letterMaxY + letterMinY) / 2, 0.0},
			vtq: int32(5)}
	}
	return &letter{}
}

func (l *letter) getVtc() []float32 {
	return l.vtc
}

func (l *letter) getVtq() int32 {
	return l.vtq
}

func (l *letter) setRune(r rune) {
	l.rn = r
}
