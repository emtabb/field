package operator

import (
	"runtime"
)

var _ = runtime.NumCPU()

func Sum(array[] float64, numCPU int) float64 {
	channels := createMultiChannelFloat64(numCPU)
	return calculateSum(array, numCPU, channels)
}

func NormalSum(array[] float64) float64 {
	sum := 0.0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}