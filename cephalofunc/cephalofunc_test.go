package cephalofunc

import (
	"fmt"
	"testing"
)

func TestSliceIndex(t *testing.T) {
	source := []string{"a", "b", "c", "d", "e"}
	test := SliceIndex(len(source), func(i int) bool {
		return source[i] == "b"
	})
	test1 := SliceIndex(len(source), func(i int) bool {
		return source[i] == "j"
	})
	if test != 1 {
		t.Error("Proper index not found")
	}
	if test1 != -1 {
		t.Error("Improper index found as proper")
	}
}

func TestSliceFilter(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	test := SliceFilter(len(source), func(i int) interface{} {
		return source[i]
	}, func(i int) bool {
		return source[i] > 3
	})
	fmt.Println("--Filter--")
	fmt.Println(test)
	if len(test) != 2 {
		t.Error("Slice not filtered")
	}
}

func TestSliceMap(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	test := SliceMap(len(source), func(i int) interface{} {
		return source[i]
	}, func(e interface{}) interface{} {
		return e.(int) + 1
	})
	fmt.Println("--Map--")
	fmt.Println(test)
	if test[0] != 2 {
		t.Error("Source slice not mapped properly")
	}
}

func TestSliceSpliceish(t *testing.T) {
	source := []int{10, 11, 12, 13, 14}
	test := SliceSpliceish(len(source), func(i int) interface{} {
		return source[i]
	}, 2, true, 99)
	fmt.Println("--Splice addition--")
	fmt.Println(test)
	if len(test) != 6 {
		t.Error("Element was not inserted in the source")
	}
	test1 := SliceSpliceish(len(source), func(i int) interface{} {
		return source[i]
	}, 1, false, 0)
	fmt.Println("--Splice deletion--")
	fmt.Println(test1)
	if len(test1) != 4 {
		t.Error("Element was not deleted from the source")
	}
}
