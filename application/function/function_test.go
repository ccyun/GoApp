package function

import (
	"log"
	"testing"
)

func TestSliceUnique(t *testing.T) {
	a := []uint8{1, 1, 1, 1, 2, 3, 4, 5, 6, 8, 89, 9, 9, 9}
	//b := SliceUnique(a).Uint64()
	a = new(Slice).Unique(a).Uint8()
	log.Println(a)
}

func TestSliceDiff(t *testing.T) {
	a := []uint64{1, 1, 1, 1, 2, 3, 4, 5, 6, 8, 89, 9, 9, 9}
	b := []uint64{1, 1, 1, 1, 2, 3, 4}
	d := []uint64{1, 1, 1, 1, 2, 3, 4, 5, 6, 89}
	//b := SliceUnique(a).Uint64()
	c := new(Slice).Diff(a, b, d).Uint64()
	log.Println(c)
}

func TestSliceMerge(t *testing.T) {
	a := []uint64{1, 2}
	b := []uint64{1, 1, 3, 4}
	d := []uint64{1, 1, 1, 1, 2, 3, 5, 6, 89}
	E := []uint64{100}
	//b := SliceUnique(a).Uint64()
	c := new(Slice).Merge(a, b, d, E).Uint64()
	log.Println(c)
}
func TestSliceIntersect(t *testing.T) {
	a := []uint64{1, 2}
	b := []uint64{1, 1, 2, 3, 4}
	d := []uint64{1, 1, 1, 1, 2, 3, 5, 6, 89}
	E := []uint64{100, 2, 1}
	//b := SliceUnique(a).Uint64()
	c := new(Slice).Intersect(a, b, d, E).Uint64()

	log.Println(c)
}

func TestInSlice(t *testing.T) {
	a := uint64(8)
	b := []uint64{1, 1, 2, 3, 4}
	//b := SliceUnique(a).Uint64()
	c := new(Slice).InSlice(a, b)
	log.Println(c)
}
