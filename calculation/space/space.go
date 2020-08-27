package space

type ISpace interface {
	Add(Vector, Vector) Vector
	Scale(float64) Vector
	Multi(Vector, Vector) float64
	Transform(Vector) Vector
	Transpose(Vector) Vector
	Dimension() int
	Interactive()
}
