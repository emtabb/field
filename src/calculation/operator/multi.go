package operator

func Multi(array1 []float64, array2 []float64, numCPU int) float64 {	
	channels := createMultiChannelFloat64(numCPU)
	return calculateMultiArray(array1, array2, numCPU, channels)
}