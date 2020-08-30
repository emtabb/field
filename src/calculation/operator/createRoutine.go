package operator

func addArray(array []float64, addarray []float64) []float64 {
	for i := 0; i < len(array); i++ {
		array[i] += addarray[i]
	}
	return array
}

func addArrayRoutine(array []float64, addarray []float64, channel chan []float64) {
	array = addArray(array, addarray)
	channel <- array
}

func multiArray(array []float64, multiArray []float64) float64 {
	total := 0.0
	for i := 0; i < len(array); i++ {
		total += array[i] * multiArray[i]
	}
	return total
}

func multiArrayRoutine(array []float64, multiarray []float64, channel chan float64) {
	channel <- multiArray(array, multiarray)
}

func sum(array []float64) float64 {
	total := 0.0
	for i := 0; i < len(array); i++ {
		total += array[i]
	}
	return total
}

func sumRoutine(array []float64, c chan float64) {
	c <- sum(array)
}

func scaleRoutine(scale float64, array []float64, channel chan []float64) {
	for i := 0; i < len(array); i++ {
		array[i] *= scale
	}
	channel <- array
}

func matmulRoutine(array []float64, matrix [][]float64, channel chan []float64) {
	tempArray := array[:]
	for i := 0; i < len(array); i++ {
		array[i] = multiArray(matrix[i], tempArray)
	}
	channel <- array
}