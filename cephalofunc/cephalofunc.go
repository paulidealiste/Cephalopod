// Package cephalofunc provides functional-flavored utilities
package cephalofunc

// SliceIndex finds the index of the supplied element in the target slice
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// SliceFilter returns new slice containing all elements that satisfy the predicate
func SliceFilter(limit int, element func(i int) interface{}, predicate func(i int) bool) []interface{} {
	slf := make([]interface{}, 0)
	for i := 0; i < limit; i++ {
		if predicate(i) {
			slf = append(slf, element(i))
		}
	}
	return slf
}

// SliceMap returns new slice containing the results of function to each element of the souce slice
func SliceMap(limit int, element func(i int) interface{}, predicate func(interface{}) interface{}) []interface{} {
	slm := make([]interface{}, 0)
	for i := 0; i < limit; i++ {
		slm = append(slm, predicate(element(i)))
	}
	return slm
}

// SliceSpliceish functions similarily to splice, but only one element can be deleted or inserted by index
func SliceSpliceish(limit int, element func(i int) interface{}, start int, add bool, new interface{}) []interface{} {
	sls := make([]interface{}, 0)
	for i := 0; i < limit; i++ {
		if add == true {
			sls = append(sls, element(i))
			if i == start {
				sls = append(sls, new)
			}
		} else {
			if i != start {
				sls = append(sls, element(i))
			}
		}
	}
	return sls
}
