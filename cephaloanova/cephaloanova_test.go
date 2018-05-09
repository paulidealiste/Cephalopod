package cephaloanova

import (
	"fmt"
	"testing"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
)

func TestAnova(t *testing.T) {
	hardinput := []cephalobjects.DataPoint{
		{X: 3.0, Y: 0.0, A: "plac", G: ""},
		{X: 2.0, Y: 0.0, A: "plac", G: ""},
		{X: 1.0, Y: 0.0, A: "plac", G: ""},
		{X: 1.0, Y: 0.0, A: "plac", G: ""},
		{X: 4.0, Y: 0.0, A: "plac", G: ""},
		{X: 5.0, Y: 0.0, A: "ldos", G: ""},
		{X: 2.0, Y: 0.0, A: "ldos", G: ""},
		{X: 4.0, Y: 0.0, A: "ldos", G: ""},
		{X: 2.0, Y: 0.0, A: "ldos", G: ""},
		{X: 3.0, Y: 0.0, A: "ldos", G: ""},
		{X: 7.0, Y: 0.0, A: "hdos", G: ""},
		{X: 4.0, Y: 0.0, A: "hdos", G: ""},
		{X: 5.0, Y: 0.0, A: "hdos", G: ""},
		{X: 3.0, Y: 0.0, A: "hdos", G: ""},
		{X: 6.0, Y: 0.0, A: "hdos", G: ""},
	}
	test := analysisOfVariance(hardinput)
	fmt.Println(test)
	if test.SST == test.SSM+test.SSE && test.Dft == test.Dfm+test.Dfe {
		t.Error("Total sums are not equal to model + error")
	}
}
