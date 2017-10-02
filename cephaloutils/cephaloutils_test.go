package cephaloutils

import (
	"fmt"
	"testing"

	"github.com/paulidealiste/Cephalopod/cephalorandom"
)

// whether extremes exist and are in correct order - min DataPoint, max DataPoint
func TestExtremesRange(t *testing.T) {
	input, _ := cephalorandom.GenerateRandomDataStore(12, 3, 0.5)
	test := ExtremesRange(&input)
	if len(test) != 2 {
		t.Error("Range does not have exactly two DataPoints")
	}
	if test[0].X > test[1].X || test[0].Y > test[1].Y {
		t.Error("Maximum exceeds minimum")
	}
}

// wheter descriptors return values
func TestDescriptors(t *testing.T) {
	input, _ := cephalorandom.GenerateRandomDataStore(12, 3, 0.5)
	test := CalculateDescriptors(input.Basic)
	fmt.Println(test)
}
