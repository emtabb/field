package operator

// Norm chuẩn hoá dữ liệu theo chiều dọc.
func Mean(array []float64, numCPU int) float64{
	return Sum(array, numCPU) / float64(len(array))
}