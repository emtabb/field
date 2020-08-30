package space

type VectorSpace interface {
	Add(Vector, Vector) Vector
	Scale(float64, Vector) Vector
	Multi(Vector, Vector) float64
	Transform(Vector, []Vector) (Vector, error)
	Transpose([]Vector) Vector
	Matrix([]Vector) [][]float64
	Dimension(int)
	MultiThread(int)
	Base(int) Vector
}
