package operator

func calculateAddArray(array []float64, addarray []float64, numCPU int, channels []chan []float64) []float64 {
	addChannel := make([]float64, len(array))
	for i := 0; i < numCPU; i++ {
		from := int(i * len(array) / numCPU)
		to := int((i + 1) * len(array) / numCPU)
		go addArrayRoutine(array[from:to], addarray[from:to], channels[i])
	}
	for i := 0; i < numCPU; i++ {
		addChannel = append(addChannel, <- channels[i]...)
	}
	return addChannel
}

func calculateSum(array []float64, numCPU int, channels []chan float64) float64 {
	sumChannel := 0.0
	for i := 0; i < numCPU; i++ {
		from := int(i * len(array) / numCPU)
		to := int((i + 1) * len(array) / numCPU)
		go sumRoutine(array[from:to], channels[i])
	}
	for i := 0; i < numCPU; i++ {
		sumChannel += <- channels[i]
	}
	return sumChannel
}

func calculateScale(scale float64, array []float64, numCPU int, channels []chan []float64) []float64 {
	scaleChannel := make([]float64, len(array))
	for i := 0; i < numCPU; i++ {
		from := int(i * len(array) / numCPU)
		to := int((i + 1) * len(array) / numCPU)
		go scaleRoutine(scale, array[from:to], channels[i])
	}
	for i := 0; i < numCPU; i++ {
		scaleChannel = append(scaleChannel, <- channels[i]...)
	}
	return scaleChannel
}

func calculateMultiArray(array []float64, multiArray []float64, numCPU int, channels []chan float64) float64 {
	multiChannel := 0.0
	for i := 0; i < numCPU; i++ {
		from := int(i * len(array) / numCPU)
		to := int((i + 1) * len(array) / numCPU)
		go multiArrayRoutine(array[from:to], multiArray[from:to], channels[i])
	}
	for i := 0; i < numCPU; i++ {
		multiChannel += <- channels[i]
	}
	return multiChannel
}

func calculateTransform(array []float64, matrix [][]float64, numCPU int, channels []chan []float64) []float64 {
	scaleChannel := make([]float64, len(array))
	for i := 0; i < numCPU; i++ {
		from := int(i * len(array) / numCPU)
		to := int((i + 1) * len(array) / numCPU)
		go matmulRoutine(array[from:to], matrix[from:to], channels[i])
	}
	for i := 0; i < numCPU; i++ {
		scaleChannel = append(scaleChannel, <- channels[i]...)
	}
	return scaleChannel
}