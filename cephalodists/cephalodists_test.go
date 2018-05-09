package cephalodists

import (
	"testing"
)

func TestProbabilityF(t *testing.T) {
	df1 := 3.0
	df2 := 15.0
	x := 2.2
	test := ProbabiltyF(x, df1, df2)
	if test == 0.0 {
		t.Error("F PDF probably not calculated right")
	}
}
