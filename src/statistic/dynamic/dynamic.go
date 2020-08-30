package dynamic

import (
	"reflect"
	"log"
	"strconv"
)

type Dynamic struct {
	Type reflect.Type
	data interface {}
}

func (d *Dynamic) Init() *Dynamic {
	d.Type = nil
	return d
}

func (d *Dynamic) Data(data interface {}) *Dynamic {
	d.data = data
	d.Type = reflect.TypeOf(d.data)
	return d
}

func (d *Dynamic) SetData(data interface {}) {
	d.data = data
	d.Type = reflect.TypeOf(d.data)
}

func (d *Dynamic) Detect() float64 {
	switch d.Kind() {
	case reflect.Int:
		return float64(d.data.(int))
	case reflect.Float32:
		return float64(d.data.(float32))
	case reflect.Float64:
		return d.data.(float64)
	case reflect.String:
		dou, _ := strconv.ParseFloat(d.data.(string), 64)
		return dou
	default:
		log.Fatal("Data type can't detect")
	}
	return -999999999999999.99999
}

func (d *Dynamic) Kind() reflect.Kind {
	return d.Type.Kind()
}

func (d *Dynamic) String() string {
	if d.IsString() {
		return d.data.(string)
	}
	switch d.Kind() {
	case reflect.Int:
		return strconv.Itoa(d.data.(int))
	case reflect.Float64:
		return strconv.FormatFloat(d.data.(float64), 'E', -6, 64)
	case reflect.Float32:
		return strconv.FormatFloat(d.data.(float64), 'E', -6, 32)
	default:
		return "error"
	}
}

func (d *Dynamic) SetType(typ reflect.Type) {
	d.Type = typ
}

func (d *Dynamic) TypeDouble() (reflect.Type) {
	return reflect.TypeOf(0.3)
}

func (d *Dynamic) TypeFloat() (reflect.Type) {
	return reflect.TypeOf(float32(0.3))
}

func (d *Dynamic) TypeString() (reflect.Type) {
	return reflect.TypeOf("")
}

func (d *Dynamic) TypeInt() (reflect.Type) {
	return reflect.TypeOf(1)
}

func (d *Dynamic) Double() float64 {
	return d.Detect()
}

func (d *Dynamic) Integer() int {
	return int(d.Detect())
}

func (d *Dynamic) BigInt() int64 {
	return int64(d.Detect())
}

func (d *Dynamic) Float() float32 {
	return float32(d.Detect())
}

func (d *Dynamic) IsString() bool {
	if d.Kind() == reflect.String {
		return true
	}
	return false
}

func (d *Dynamic) IsInteger() bool {
	if d.Kind() == reflect.Int {
		return true
	}
	return false
}

func (d *Dynamic) IsFloat() bool {
	if d.Kind() == reflect.Float32 {
		return true
	}
	return false
}

func (d *Dynamic) IsDouble() bool {
	if d.Kind() == reflect.Float64 {
		return true
	}
	return false
}

func IsString(data interface {}) bool {
	dataType := reflect.TypeOf(data)
	if dt := dataType.Kind(); dt == reflect.String {
		return true
	}
	return false
}

func IsInteger(data interface {}) bool {
	dataType := reflect.TypeOf(data)
	if dt := dataType.Kind(); dt == reflect.Int {
		return true
	}
	return false
}

func IsFloat(data interface {}) bool {
	dataType := reflect.TypeOf(data)
	if dt := dataType.Kind(); dt == reflect.Float32 {
		return true
	}
	return false
}

func IsDouble(data interface {}) bool {
	dataType := reflect.TypeOf(data)
	if dt := dataType.Kind(); dt == reflect.Float64 {
		return true
	}
	return false
}