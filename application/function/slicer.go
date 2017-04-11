package function

import (
	"reflect"
)

//Slice 切片结构体
type Slice struct {
	data []interface{}
}

//Unique 切片去重复
func (S *Slice) Unique(sliceData interface{}) *Slice {
	for k := range toMap(sliceData) {
		S.data = append(S.data, k)
	}
	return S
}

//Diff 返回切片差集
func (S *Slice) Diff(slices ...interface{}) *Slice {
	snum := len(slices)
	if snum < 1 {
		return S
	}
	map1 := toMap(slices[0])
	tempData := make(map[interface{}]bool)
	for i := 1; i < snum; i++ {
		for v := range toMap((slices[i])) {
			tempData[v] = true
		}
	}
	for k := range map1 {
		if _, ok := tempData[k]; ok == false {
			S.data = append(S.data, k)
		}
	}
	return S
}

//Merge 多个切片的并集
func (S *Slice) Merge(slices ...interface{}) *Slice {
	snum := len(slices)
	if snum < 1 {
		return S
	}
	tempData := make(map[interface{}]bool)
	for i := 0; i < snum; i++ {
		for v := range toMap((slices[i])) {
			tempData[v] = true
		}
	}
	for k := range tempData {
		S.data = append(S.data, k)
	}
	return S
}

//Intersect 返回切片交集
func (S *Slice) Intersect(slices ...interface{}) *Slice {
	snum := len(slices)
	if snum < 1 {
		return S
	}
	tempData := make(map[interface{}]int)
	for i := 0; i < snum; i++ {
		for v := range toMap((slices[i])) {
			tempData[v]++
		}
	}
	for k, v := range tempData {
		if v == snum {
			S.data = append(S.data, k)
		}
	}
	return S
}

//InSlice 判断切片是否存在某个元素
func (S *Slice) InSlice(value interface{}, sliceData interface{}) bool {
	temp := toMap(sliceData)
	if ok := temp[value]; ok {
		return true
	}
	return false
}

//toMap 切片转map
func toMap(sliceData interface{}) map[interface{}]bool {
	mapData := make(map[interface{}]bool)
	var inPutData []interface{}
	rv := reflect.ValueOf(sliceData)
	for i := 0; i < rv.Len(); i++ {
		inPutData = append(inPutData, rv.Index(i).Interface())
	}
	for _, v := range inPutData {
		mapData[v] = true
	}
	return mapData
}

//Uint64 to []uint64
func (S *Slice) Uint64() []uint64 {
	var data []uint64
	for _, v := range S.data {
		if vv, ok := v.(uint64); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Uint32 to []uint32
func (S *Slice) Uint32() []uint32 {
	var data []uint32
	for _, v := range S.data {
		if vv, ok := v.(uint32); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Uint16 to []uint16
func (S *Slice) Uint16() []uint16 {
	var data []uint16
	for _, v := range S.data {
		if vv, ok := v.(uint16); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Uint8 to []uint8
func (S *Slice) Uint8() []uint8 {
	var data []uint8
	for _, v := range S.data {
		if vv, ok := v.(uint8); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Uint to []uint
func (S *Slice) Uint() []uint {
	var data []uint
	for _, v := range S.data {
		if vv, ok := v.(uint); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Int64 to []int64
func (S *Slice) Int64() []int64 {
	var data []int64
	for _, v := range S.data {
		if vv, ok := v.(int64); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Int32 to []int32
func (S *Slice) Int32() []int32 {
	var data []int32
	for _, v := range S.data {
		if vv, ok := v.(int32); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Int16 to []int16
func (S *Slice) Int16() []int16 {
	var data []int16
	for _, v := range S.data {
		if vv, ok := v.(int16); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Int8 to []int8
func (S *Slice) Int8() []int8 {
	var data []int8
	for _, v := range S.data {
		if vv, ok := v.(int8); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Int to []int
func (S *Slice) Int() []int {
	var data []int
	for _, v := range S.data {
		if vv, ok := v.(int); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Float64 to []float64
func (S *Slice) Float64() []float64 {
	var data []float64
	for _, v := range S.data {
		if vv, ok := v.(float64); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Float32 to []float32
func (S *Slice) Float32() []float32 {
	var data []float32
	for _, v := range S.data {
		if vv, ok := v.(float32); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Byte to []byte
func (S *Slice) Byte() []byte {
	var data []byte
	for _, v := range S.data {
		if vv, ok := v.(byte); ok {
			data = append(data, vv)
		}
	}
	return data
}

//String to []string
func (S *Slice) String() []string {
	var data []string
	for _, v := range S.data {
		if vv, ok := v.(string); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Bool to []bool
func (S *Slice) Bool() []bool {
	var data []bool
	for _, v := range S.data {
		if vv, ok := v.(bool); ok {
			data = append(data, vv)
		}
	}
	return data
}

//Interface to []interface
func (S *Slice) Interface() []interface{} {
	return S.data
}
