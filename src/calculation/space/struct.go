package space

type VectorSpaceStruct struct {
	VectorSpace
	dimension int
	thread int
}

func (vs *VectorSpaceStruct) Init() VectorSpace {
	vs.dimension = 1
	vs.thread = 1
	return vs
}

func (vs *VectorSpaceStruct) MultiThread(thread int) {
	vs.thread = thread
}

func (vs *VectorSpaceStruct) Add(vector1 Vector, vector2 Vector) Vector {
	return vector1.Add(vector2, vs.thread)
}

func (vs *VectorSpaceStruct) Scale(scale float64, vector Vector) Vector {
	return vector.Scale(scale, vs.thread)
}

func (vs *VectorSpaceStruct) Multi(vector1 Vector, vector2 Vector) float64 {
	return vector1.Multi(vector2, vs.thread)
}

func (vs *VectorSpaceStruct) Transform(vector Vector, vectors []Vector) (Vector, error) {
	return vector.Transform(vectors, vs.thread)
}

func (vs *VectorSpaceStruct) Matrix(vectors []Vector) [][]float64 {
	return Matrix(vectors)
}

func (vs *VectorSpaceStruct) Dimension(dimension int) {
	vs.dimension = dimension
}

func (vs *VectorSpaceStruct) Base(position int) Vector {
	var vector Vector
	return vector.Base(position)
}