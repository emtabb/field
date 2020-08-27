package calculate

func Scale(scale float64, array []float64) []float64 {
	var numCPU = 4	
	channels := createMultiChannelArrayFloat64(numCPU)
	return calculateScale(scale, array, numCPU, channels)
}