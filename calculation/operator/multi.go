package calculate

func Multi(array... []float64) float64 {
	if len(array) != 2 {
		return 0.0
	}
	var numCPU = 4	
	channels := createMultiChannelFloat64(numCPU)
	return calculateMultiArray(array[0], array[1], numCPU, channels)
}