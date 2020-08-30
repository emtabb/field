package operator

func createMultiChannelFloat64(numCPU int) []chan float64 {
	channels := make([]chan float64, numCPU)
	for i := 0; i < numCPU; i++ {
		channels[i] = make(chan float64)
	}
	return channels
}

func createMultiChannelArrayFloat64(numCPU int) []chan []float64 {
	channels := make([]chan []float64, numCPU)
	for i := 0; i < numCPU; i++ {
		channels[i] = make(chan []float64)
	}
	return channels
}