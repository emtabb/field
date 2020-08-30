package operator

func Transform(array []float64, matrix [][]float64, numCPU int) []float64 {
	channels := createMultiChannelArrayFloat64(numCPU)
	return calculateTransform(array, Transpose(matrix), numCPU, channels)
}

func Transpose(matrix [][]float64) [][]float64 {
	dimension := len(matrix)
	newMatrix := make([][]float64, dimension)
	for i := 0; i < dimension; i++ {
		newMatrix[i] = make([]float64, dimension)
	}
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			newMatrix[i][j] = matrix[j][i]
		}
	}
	return newMatrix
}