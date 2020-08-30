package operator

func Scale(scale float64, array []float64, numCPU int) []float64 {
	channels := createMultiChannelArrayFloat64(numCPU)
	return calculateScale(scale, array, numCPU, channels)
}