package space

type HilbertSpace struct {
	ISpace
}

func (hs *HilbertSpace) Add(vector1 Vector, vector2 Vector) Vector {
	return vector1.Add(vector2)	
}