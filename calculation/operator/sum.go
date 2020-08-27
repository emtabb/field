package calculate

import (
	"runtime"
	"fmt"
)

var _ = runtime.NumCPU()
var _ = fmt.Sprintf("")

func Sum(array[] float64) float64 {
	var numCPU = 4	
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