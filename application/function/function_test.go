package function

import (
	"log"
	"testing"
)

func TestSliceUnique(t *testing.T) {
	a := new(Slice).Unique([]uint8{1, 1, 1, 1, 2, 3, 4, 5, 6, 8, 89, 9, 9, 9}).Uint8()
	log.Println(a)

	b := new(Slice).Unique([]uint{1, 1, 1, 1, 2, 3, 4, 5, 6, 8, 89, 9, 9, 9}).Uint()
	log.Println(b)

	c := new(Slice).Unique([]int64{1, 1, 1, 1, 2, 3, 4, 5, 6, 8, 89, 9, 9, 9}).Int64()
	log.Println(c)

	d := new(Slice).Unique([]int8{1, 1, 1, 1, 2, 3, 4, 5, 6, 8, 89, 9, 9, 9}).Int8()
	log.Println(d)

	e := new(Slice).Unique([]string{"1", "2", "1", "2", "3"}).String()
	log.Println(e)

	f := new(Slice).Unique([]bool{true, false, true, false}).Bool()
	log.Println(f)

	g := new(Slice).Unique([]int16{1, 2, 1, 1, 1, 2, 2, 2, 3, 4, 5, 8, 9, 65, 5, 5}).Int16()
	log.Println(g)
}

func TestSliceDiff(t *testing.T) {
	a := []uint64{1, 1, 1, 1, 2, 3, 4, 5, 6, 8, 89, 9, 9, 9}
	b := []uint64{1, 1, 1, 1, 2, 3, 4}
	c := []uint64{1, 1, 1, 1, 2, 3, 9}
	d := []uint64{1, 1, 1, 1, 2, 3, 4, 5, 6, 89}

	ss := new(Slice).Diff(a, b, c, d).Uint64()
	log.Println(ss)

	a1 := []string{"1", "2", "3", "4"}
	b1 := []string{"1"}
	c1 := []string{"5"}
	d1 := []string{"8"}

	ss1 := new(Slice).Diff(a1, b1, c1, d1).String()
	log.Println(ss1)

}

func TestSliceMerge(t *testing.T) {
	a := []uint64{1, 2}
	b := []uint64{1, 1, 3, 4}
	c := []uint64{1, 1, 3, 7, 4}
	d := []uint64{1, 1, 1, 1, 2, 3, 5, 6, 89}
	e := []uint64{100}

	ss := new(Slice).Merge(a, b, c, d, e).Uint64()
	log.Println(ss)
}
func TestSliceIntersect(t *testing.T) {
	a := []uint64{1, 2}
	b := []uint64{1, 1, 2, 3, 4}
	d := []uint64{1, 1, 1, 1, 2, 3, 5, 6, 89}
	E := []uint64{100, 2, 1}
	c := new(Slice).Intersect(a, b, d, E).Uint64()
	log.Println(c)
}

func TestInSlice(t *testing.T) {
	a := uint64(1)
	b := []uint64{1, 1, 2, 3, 4}
	c := new(Slice).InSlice(a, b)
	log.Println(c)
}