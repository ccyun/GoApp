package function

import (
	"log"
	"reflect"
	"testing"
)

func ArrayValues(data interface{}, outData interface{}) {
	var (
		inPutData []interface{}
	)
	tempData := make(map[interface{}]bool)
	rv := reflect.ValueOf(data)
	for i := 0; i < rv.Len(); i++ {
		inPutData = append(inPutData, rv.Index(i).Interface())
	}
	for _, v := range inPutData {
		tempData[v] = true
	}

}

func TestValues(t *testing.T) {
	a := []uint64{1, 1, 1, 1, 2, 3, 4, 5, 6, 8, 89, 9, 9, 9}
	b := []uint64{}
	ArrayValues(a, &b)
	log.Println(b)
}
