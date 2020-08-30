package operator

func Add(array1 []float64, array2 []float64, numCPU int) []float64 {
	channels := createMultiChannelArrayFloat64(numCPU)
	return calculateAddArray(array1, array2, numCPU, channels)
}