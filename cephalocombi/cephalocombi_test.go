package cephalocombi

import (
	"testing"
)

func TestCartesianProduct(t *testing.T) {

	fl := []int{1, 2, 3, 4, 5}
	forcarts := [][]int{fl, fl, fl, fl, fl, fl, fl}

	test := CartesianProduct(forcarts)

	if len(test) != 5*5*5*5*5*5*5 {
		t.Error("Haven't generated all of the possible combinations")
	}

}
