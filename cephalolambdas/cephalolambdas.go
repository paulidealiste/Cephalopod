// Package cephalolambdas provides lodash-like functions for the Cephalopod
package cephalolambdas

// Beakdiff returns elements present in the first and absent in the second array
func Beakdiff(collection []string, targetcollection []string) []string {
	dfunc := func(target string) bool {
		for _, s := range targetcollection {
			if target == s {
				return false
			}
		}
		return true
	}
	diff := Beakfilter(collection, dfunc)
	return diff
}

// Beakfilter filters the elements from the array based on the condition supplied
func Beakfilter(collection []string, beakfun func(string) bool) []string {
	var flt []string
	for _, ff := range collection {
		if beakfun(ff) {
			flt = append(flt, ff)
		}
	}
	return flt
}

// Beakkeys returns the map keys as a string array
func Beakkeys(collection map[string]string) []string {
	bks := make([]string, 0, len(collection))
	for key := range collection {
		bks = append(bks, key)
	}
	return bks
}

// Beakindex returns the index/position of the target within the collection
func Beakindex(collection []string, target string) int {
	for i, item := range collection {
		if item == target {
			return i
		}
	}
	return -1
}
