package test

import (
	"fmt"
	"testing"
)
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
func Test(t *testing.T) {
	xzc := make([][] float64, 3)
	for i := 0; i < len(xzc); i++ {
		xzc[i] = make([] float64, 3)
	}
	for i := 0; i < len(xzc); i++ {
		for j := 0; j < len(xzc[0]); j++ {
			xzc[i][j] = float64(i)
			fmt.Print(xzc[i][j])
		}
		fmt.Println()
	}
	c := Transpose(xzc)
	for i := 0; i < len(xzc); i++ {
		for j := 0; j < len(xzc[0]); j++ {
			fmt.Print(c[i][j])
		}
		fmt.Println()
	}
}