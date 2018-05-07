package cephaloanova

import (
	"testing"

	"github.com/paulidealiste/Cephalopod/cephalorandom"
)

func TestAnova(t *testing.T) {
	input, _ := cephalorandom.GenerateRandomDataStore(120, 3, 0.5)
	test := analysisOfVariance(input.Basic)
	if test.SST == test.SSM+test.SSE && test.Dft == test.Dfm+test.Dfe {
		t.Error("Total sums are not equal to model + error")
	}
}
