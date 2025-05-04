package digits

func CreateVertexDigits(number int, offset float32) ([]float32, int32) {
	switch number{
	case 1:
		return []float32 {	
			-0.8 + offset, 0.0, 0.0,
			-0.7 + offset, 0.3, 0.0,
			-0.7 + offset, -0.275, 0.0,
		}, 3
	case 2:
		return []float32 {
			-0.9 + offset, 0.2, 0.0,
			-0.8 + offset, 0.3, 0.0,
			-0.7 + offset, 0.2, 0.0,
		
			-0.9 + offset, -0.2, 0.0,
			-0.6 + offset, -0.2, 0.0,
		}, 5
	default:
		return []float32 {
			0.0, 0.0, 0.0,
		}, 1 
	}
}
