package interaction

import . "github.com/emtabb/field/src/calculation/space"

type Interaction interface {
	Init([]Vector)
	Generate() []Vector
}