package calculate

func Mean(array []float64) float64{
	return Sum(array) / float64(len(array))
}