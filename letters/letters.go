package letters

func CreateVertexLetters(letter rune, offset float32) ([]float32, int32, float32) {
	switch letter {
	case 'T':
		return []float32{
			-0.82 + offset, -0.4, 0.0,
			-0.75 + offset, -0.4, 0.0,

			-0.75 + offset, -0.71, 0.0,
			-0.75 + offset, -0.4, 0.0,

			-0.68 + offset, -0.4, 0.0,
		}, 5, 0.15
	case 'I':
		return []float32{
			-0.8 + offset, -0.4, 0.0,
			-0.8 + offset, -0.72, 0.0,
		}, 2, 0.06
	case 'M':
		return []float32{
			-0.8 + offset, -0.7, 0.0,
			-0.8 + offset, -0.4, 0.0,

			-0.75 + offset, -0.7, 0.0,
			-0.7 + offset, -0.4, 0.0,

			-0.7 + offset, -0.72, 0.0,
		}, 5, 0.15
	case 'E':
		return []float32{
			-0.73 + offset, -0.4, 0.0,
			-0.8 + offset, -0.4, 0.0,
			-0.8 + offset, -0.55, 0.0,

			-0.73 + offset, -0.55, 0.0,
			-0.8 + offset, -0.55, 0.0,

			-0.8 + offset, -0.7, 0.0,
			-0.73 + offset, -0.7, 0.0,
		}, 7, 0.1
	case 'R':
		return []float32{
			-0.8 + offset, -0.72, 0.0,
			-0.8 + offset, -0.4, 0.0,

			-0.75 + offset, -0.4, 0.0,
			-0.7 + offset, -0.53, 0.0,
			-0.8 + offset, -0.58, 0.0,

			-0.71 + offset, -0.715, 0.0,
		}, 6, 0.15
	case 'O':
		return []float32{
			-0.8 + offset, -0.4, 0.0,
			-0.8 + offset, -0.7, 0.0,

			-0.7 + offset, -0.7, 0.0,
			-0.7 + offset, -0.4, 0.0,

			-0.8 + offset, -0.4, 0.0,
		}, 5, 0.15
	case 'N':
		return []float32{
			-0.8 + offset, -0.7, 0.0,
			-0.8 + offset, -0.4, 0.0,

			-0.7 + offset, -0.7, 0.0,
			-0.7 + offset, -0.4, 0.0,
		}, 4, 0.2
	case 'P':
		return []float32{
			-0.8 + offset, -0.73, 0.0,
			-0.8 + offset, -0.4, 0.0,

			-0.75 + offset, -0.4, 0.0,
			-0.7 + offset, -0.55, 0.0,
			-0.8 + offset, -0.62, 0.0,
		}, 5, 0.15
	default:
		return []float32{
			0.0, 0.0, 0.0,
		}, 1, 0.15
	}
}
