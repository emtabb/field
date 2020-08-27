package calculate

func Add(array... []float64) []float64 {
	if len(array) != 2 {
		return nil
	}
	var numCPU = 4	
	channels := createMultiChannelArrayFloat64(numCPU)
	return calculateAddArray(array[0], array[1], numCPU, channels)
}