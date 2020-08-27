package field

import (
	. "emtabb/field/statistic/dynamic"
	. "emtabb/field/calculation/space"
	"emtabb/field/statistic"
	"reflect"
	"errors"
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
	typef reflect.Type
	dynamic *Dynamic
	space ISpace
	
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
}

func (f *Field) Init() *Field {
	f.generate()
	return f
}

func (f *Field) Name(name string) *Field {
	f.name = name
	return f
}

func (f *Field) Data(data []interface {}) *Field {
	f.setupData(data)
	return f
}

func (f *Field) SetData(data []interface {}) {
	f.setupData(data)
}

func (f *Field) clearCache() {
	f.floatCache = []float32 {}
	f.doubleCache = []float64 {}
	f.stringCache = []string {}
	f.intCache = []int {}
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

func (f *Field) setupData(data []interface {}) {
	f.data = data[:]
	f.SetType(DOUBLE)
	f.clearCache()
}
