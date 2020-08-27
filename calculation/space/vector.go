package space

import "emtabb/field/calculation/operator"

type Vector struct {
	coor []float64
}

func (v Vector) Add(vector Vector) Vector {
	v.coor = calculate.Add(v.coor, vector.coor)
	return v
}

func (v Vector) Scale(scale float64) Vector {
	v.coor = calculate.Scale(scale, v.coor)
	return v
}

func (v Vector) Multi(vector Vector) float64 {
	return calculate.Multi(v.coor, vector.coor)
}