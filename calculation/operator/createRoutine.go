package calculate

func addArrayRoutine(array []float64, addarray []float64, channel chan []float64) {
	for i := 0; i < len(array); i++ {
		array[i] += addarray[i]
	}
	channel <- array
}

func multiArrayRoutine(array []float64, multiArray []float64, channel chan float64) {
	total := 0.0
	for i := 0; i < len(array); i++ {
		total += array[i] * multiArray[i]
	}
	channel <- total
}

func sumRoutine(array[] float64, c chan float64) {
	total := 0.0
	for i := 0; i < len(array); i++ {
		total += array[i]
	}
	c <- total
}

func scaleRoutine(scale float64, array []float64, channel chan []float64) {
	for i := 0; i < len(array); i++ {
		array[i] *= scale
	}
	channel <- array
}