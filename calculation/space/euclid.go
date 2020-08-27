package space

type EuclidSpace struct {
	ISpace
}

func (es *EuclidSpace) Add(vector1 Vector, vector2 Vector) Vector {
	return vector1.Add(vector2)
}

func (es *EuclidSpace) Scale(scale float64, vector Vector) Vector {
	return vector.Scale(scale)
}

func (es *EuclidSpace) Multi(vector1 Vector, vector2 Vector) float64 {
	return vector1.Multi(vector2)
}