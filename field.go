package field

import (
	"reflect"
	"errors"
	. "github.com/emtabb/field/src/statistic/dynamic"
	. "github.com/emtabb/field/src/calculation/space"
	"github.com/emtabb/field/src/statistic"
	"github.com/emtabb/field/interaction"
)

const (
	STRING  = "STRING"
	DOUBLE  = "DOUBLE"
	FLOAT   = "FLOAT"
	INT     = "INT"
	BIGINT  = "BIGINT"
)

type Field struct {
	name string
	data []interface {}
	count int
	typef reflect.Type
	dynamic *Dynamic
	//
	vectorSpace VectorSpace
	interaction interaction.Interaction

	// Cache 
	doubleCache []float64
	floatCache []float32
	stringCache []string
	intCache []int
	
	// isGenerate
	isGenerate bool
}

func (f *Field) generate() {
	f.isGenerate = true
	f.name = "FIELD"
	f.data = make([]interface {}, 0)
	f.count = 0
}

func (f *Field) Init() *Field {
	f.generate()
	return f
}

func (f *Field) Name(name string) *Field {
	f.name = name
	return f
}

func (f *Field) GetName() string {
	return f.name
}

func (f *Field) Data(data []interface {}) *Field {
	f.setupData(data)
	return f
}

func (f *Field) SetData(data []interface {}) {
	f.setupData(data)
}

func (f *Field) GetData() []interface {} {
	tempData := make([] interface{}, f.count)
	switch v := f.typef.Kind(); v {
	case reflect.Float64:
		tempData = f.InterfaceToDouble()[:]
	case reflect.Float32:
		tempData = f.InterfaceToFloat()[:]
	case reflect.Int:
		tempData = f.InterfaceToInteger()[:]
	case reflect.String:
		tempData = f.InterfaceToString()[:]
	default:
		tempData = nil
	}
	return tempData
}

// Need upgrade to multithread after
func (f *Field) InterfaceToDouble() []interface {} {
	data := make([]interface {}, f.count)
	doubleData := f.Double()
	for i, double := range doubleData {
		data[i] = double
	}
	return data
}

func (f *Field) InterfaceToFloat() []interface {} {
	data := make([]interface {}, f.count)
	doubleData := f.Float()
	for i, double := range doubleData {
		data[i] = double
	}
	return data
}

func (f *Field) InterfaceToString() []interface {} {
	data := make([]interface {}, f.count)
	stringData := f.String()
	for i, strD := range stringData {
		data[i] = strD
	}
	return data
}

func (f *Field) InterfaceToInteger() []interface {} {
	data := make([]interface {}, f.count)
	intData := f.Integer()
	for i, intd := range intData {
		data[i] = intd
	}
	return data
}

func (f *Field) Statistic() statistic.StatisticType {
	return statistic.Statistic(f.data)
}

func (f *Field) clearCache() {
	f.floatCache = []float32 {}
	f.doubleCache = []float64 {}
	f.stringCache = []string {}
	f.intCache = []int {}
}

func (f *Field) Integer() []int {
	if len(f.intCache) == 0 {
		f.intCache = statistic.Integer(f.data)
	}
	return f.intCache
}

func (f *Field) Float() []float32 {
	if len(f.floatCache) == 0 {
		f.floatCache = statistic.Float(f.data)
	}
	return f.floatCache
}

func (f *Field) Double() []float64 {
	if len(f.doubleCache) == 0 {
		f.doubleCache = statistic.Double(f.data)
	}
	return f.doubleCache
}

func (f *Field) String() []string {
	if len(f.stringCache) == 0 {
		f.stringCache = statistic.String(f.data)
	}
	return f.stringCache
}

func (f *Field) SetType(typef string) error {
	switch typef {
	case STRING :
		f.typef = reflect.TypeOf("")
	case FLOAT  :
		f.typef = reflect.TypeOf(float32(0.0))
	case DOUBLE :
		f.typef = reflect.TypeOf(float64(0.0))
	case INT:
		f.typef = reflect.TypeOf(int(1))
	case BIGINT:
		f.typef = reflect.TypeOf(int64(1))
	default:
		f.typef = reflect.TypeOf(float64(0.0))
		err := errors.New("Type input is not define")
		return err
	}
	return nil
}

func (f *Field) GetType() string {
	stringType := "'"
	switch v := f.typef.Kind(); v {
	case reflect.String :
		stringType = STRING
	case reflect.Float32  :
		stringType = FLOAT
	case reflect.Float64 :
		stringType = DOUBLE
	case reflect.Int:
		stringType = INT
	case reflect.Int64:
		stringType = BIGINT
	default:
		stringType = "NaN"
	}
	return stringType
}

func (f *Field) setupData(data []interface {}) {
	f.data = data[:]
	f.count = len(f.data)
	f.SetType(DOUBLE)
	f.clearCache()
}

func (f *Field) SetVectorSpace(vectorSpace VectorSpace) {
	f.vectorSpace = vectorSpace
}

func (f *Field) VectorSpace() VectorSpace {
	return f.vectorSpace
}

func (f *Field) Distinct() []string {
	return statistic.Distinct(f.GetData())
}

func (f *Field) IsContainNaN() []int {
	return statistic.ContainNaN(f.GetData())
}

func (f *Field) IsContainINFINITY() []int {
	return statistic.ContainINFINITY(f.GetData())
}

func (f *Field) Find(data interface {}) int {
	return statistic.Find(data, f.GetData())
}

func (f *Field) FindArray(data []interface {}) []int {
	return statistic.FindArray(data, f.GetData())
}

func (f *Field) SetInteraction(interaction interaction.Interaction) {
	f.interaction = interaction
}

func (f *Field) Interaction() interaction.Interaction {
	return f.interaction
}

func (f *Field) Interact(vectors []Vector) []Vector {
	f.interaction.Init(vectors)
	return f.interaction.Generate()
}

func (f *Field) Vector() Vector {
	var v Vector
	v.SetCoordinate(f.Double())
	return v
}