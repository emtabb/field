package statistic

import "github.com/emtabb/field/src/statistic/dynamic"

const (
	NA = "NA"
	NN = "NN"
	NaN = "NaN"
	NotANumber = "Not a numbers"
	NEGATIVE_INFINITY = -999999999999999.99999
	INFINITY = 999999999999999.99999
)

type StatisticType struct {
	Count map[string] int
}

func (st StatisticType) Make() StatisticType {
	st.Count["Double"] = 0
	st.Count["Float"] = 0
	st.Count["String"] = 0
	st.Count["Int"] = 0
	st.Count["Others"] = 0
	return st
}

func (st StatisticType) CountInt() {
	st.Count["Int"]++
}
func (st StatisticType) CountDouble() {
	st.Count["Double"]++
}
func (st StatisticType) CountFloat() {
	st.Count["Float"]++
}
func (st StatisticType) CountString() {
	st.Count["String"]++
}
func (st StatisticType) CountOthers() {
	st.Count["Others"]++
}

func (st StatisticType) Array() []int {
	return []int { st.Count["String"], st.Count["Int"], st.Count["Float"], st.Count["Double"], st.Count["Others"] }
}

func (st StatisticType) MostCount() {
	
}

func Statistic(list []interface {}) StatisticType {
	dynamical := new(dynamic.Dynamic).Init()
	var statisticsType StatisticType
	statisticsType = statisticsType.Make()
	for _, data := range list {
		dynamical = dynamical.Data(data)
		if dynamical.IsString() {
			statisticsType.CountString()
			continue
		}
		if dynamical.IsDouble() {
			statisticsType.CountDouble()
			continue
		}
		if dynamical.IsFloat() {
			statisticsType.CountFloat()
			continue
		}
		if dynamical.IsInteger() {
			statisticsType.CountInt()
			continue
		}
		statisticsType.CountOthers()
	}
	return statisticsType
}

func Integer(list []interface {}) []int {
	dynamical := new(dynamic.Dynamic).Init()
	intList := make([]int, 0)
	for _, data := range list {
		intList = append(intList, dynamical.Data(data).Integer())
	}
	return intList
}

func Float(list []interface {}) []float32 {
	dynamical := new(dynamic.Dynamic).Init()
	floatList := make([]float32, 0)
	for _, data := range list {
		floatList = append(floatList, dynamical.Data(data).Float())
	}
	return floatList
}

func Double(list []interface {}) []float64 {
	dynamical := new(dynamic.Dynamic).Init()
	doubleList := make([]float64, 0)
	for _, data := range list {
		doubleList = append(doubleList, dynamical.Data(data).Double())
	}
	return doubleList
}

func String(list []interface {}) []string {
	dynamical := new(dynamic.Dynamic).Init()
	stringList := make([]string, 0)
	for _, data := range list {
		stringList = append(stringList, dynamical.Data(data).String())
	}
	return stringList
}

func Distinct(list []interface {}) []string {
	distinct := make([]string, 0)
	for _, str := range list {
		if !ContainString(str.(string), distinct) {
			distinct = append(distinct, str.(string))
		}
	}
	return distinct
}

func ContainNaN(list []interface {}) []int {
	positions := make([]int, 0)
	for i, data := range list {
		if IsContainNaNFromDefinition(data) {
			positions = append(positions, i)
		}
	}
	return positions
}

func ContainINFINITY(list []interface {}) []int {
	positions := make([]int, 0)
	for i, data := range list {
		if data == INFINITY || data == NEGATIVE_INFINITY {
			positions = append(positions, i)
		}
	}
	return positions
}

func NaNGroupDefinitation() []string{
	return []string { NA, NN, NaN, NotANumber } 
}

func IsContainNaNFromDefinition(data interface{}) bool {
	if !dynamic.IsString(data) {
		return false
	}
	group := NaNGroupDefinitation()
	return ContainString(data.(string), group)
}

func Contain(data interface{}, list []interface {}) (bool, int) {
	for position, define := range list {
		if data == define {
			return true, position
		}
	}
	return false, 0
}

func ContainString(data string, list []string) bool {
	for _, define := range list {
		if data == define {
			return true
		}
	}
	return false
}

func Find(data interface {}, list []interface {}) int {
	isContain, position := Contain(data, list)
	if isContain {
		return position
	}
	return -1
}

func FindArray(data []interface {}, list []interface {}) []int {
	result := make([]int, len(list))
	for _, subdata := range data {
		result = append(result, Find(subdata, list))
	}
	return result
}
