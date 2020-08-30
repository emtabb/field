package space

import (
	"github.com/emtabb/field/src/calculation/operator"
	"errors"
)

type Vector struct {
	coor []float64
}

func (v Vector) Add(vector Vector, numCPU int) Vector {
	v.coor = operator.Add(v.coor, vector.coor, numCPU)
	return v
}

func (v Vector) Scale(scale float64, numCPU int) Vector {
	v.coor = operator.Scale(scale, v.coor, numCPU)
	return v
}

func (v Vector) Multi(vector Vector, numCPU int) float64 {
	return operator.Multi(v.coor, vector.coor, numCPU)
}

func (v Vector) Transform(vectors []Vector, numCPU int) (Vector, error) {
	var baseVector Vector
	if len(vectors) == 1 {
		tempVectors := make([]Vector, len(v.Coordinate()))
		vectorFactor := vectors[0].Coordinate()
		for i := 0; i < len(v.Coordinate()); i++ {
			tempVectors[i] = baseVector.Base(i).Scale(vectorFactor[i], numCPU)
		}
		vectors = tempVectors[:]
	}
	if len(vectors) != len(v.Coordinate()) {
		return baseVector.Zero(1), errors.New("Errors in calculation vector")
	}
	v.coor = operator.Transform(v.Coordinate(), Matrix(vectors), numCPU)
	return v, nil
}

func (v Vector) Zero(dimension int) Vector {
	v.coor = make([]float64, dimension)
	for i := 0; i < dimension; i++ {
		v.coor[i] = 0.0
	}
	return v
}

func (v Vector) Ones(dimension int) Vector {
	v.coor = make([]float64, dimension)
	for i := 0; i < dimension; i++ {
		v.coor[i] = 1.0
	}
	return v
}

func (v Vector) Base(position int) Vector {
	v = new(Vector).Zero(len(v.coor))
	v.coor[position] = 1.0
	return v
}

func (v Vector) SetCoordinate(coor []float64) {
	v.coor = coor[:]
}

func (v Vector) Coordinate() []float64 {
	return v.coor
}

func Matrix(vectors []Vector) [][]float64 {
	dimension := len(vectors)
	matrix := make([][]float64, dimension)
	for i := 0; i < dimension; i++ {
		matrix[i] = make([]float64, dimension)
		matrix[i] = vectors[i].Coordinate()[:]
	}
	return matrix
}