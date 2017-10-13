package cephaloutils

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
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

// Euclidean distance properties test
func TestEuclideanDistance(t *testing.T) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	p1 := cephalobjects.DataPoint{X: random.NormFloat64(), Y: random.NormFloat64()}
	p2 := cephalobjects.DataPoint{X: random.NormFloat64(), Y: random.NormFloat64()}
	test := EuclideanDistance(p1, p2)
	internal := math.Sqrt(math.Pow((p1.X-p2.X), 2) + math.Pow((p1.Y-p2.Y), 2))
	if test != internal {
		t.Error("Obtained value is not an Euclidean distance")
	}
}

// Wheter descriptors return values and are generated numbers really truncated to the desired boundaries (mean + 2SDs)
func TestDescriptorsAndTruncatedNormal(t *testing.T) {
	input, _ := cephalorandom.GenerateRandomDataStore(12, 3, 0.5)
	desc := CalculateDescriptors(input.Basic)
	test := TruncatedNormal(desc, 3)
	upperBoundX := desc.MeanX + desc.SdX
	lowerBoundX := desc.MeanX - desc.SdX
	upperBoundY := desc.MeanY + desc.SdY
	lowerBoundY := desc.MeanY - desc.SdY
	for _, dp := range test {
		if dp.X > upperBoundX || dp.X < lowerBoundX || dp.Y > upperBoundY || dp.Y < lowerBoundY {
			t.Error("Generated data fell outside of desired boundaries (mean +/- 2SD)")
		}
	}
}

// Wheter min slice index is realy a slice minimum
func TestMinSliceIndex(t *testing.T) {
	input := []float64{5.1, 3.2, 4.2, 9.4}
	test := MinSliceIndex(input)
	if input[test] != 3.2 {
		t.Error("Minimum index not found")
	}
}

// Wheter sameness test really tests for sameness
func TestCheckIfSame(t *testing.T) {
	input, _ := cephalorandom.GenerateRandomDataStore(12, 3, 0.5)
	input2, _ := cephalorandom.GenerateRandomDataStore(12, 3, 0.5)
	clone1 := make([]cephalobjects.DataPoint, len(input.Basic))
	clone2 := make([]cephalobjects.DataPoint, len(input.Basic))
	clone3 := make([]cephalobjects.DataPoint, len(input.Basic)-1)
	copy(input.Basic, clone1)
	copy(input.Basic, clone2)
	copy(input.Basic[0:len(input.Basic)-1], clone3)
	test, _ := CheckIfSame(clone1, clone2)
	_, err := CheckIfSame(clone1, clone3)
	test2, _ := CheckIfSame(clone1, input2.Basic)
	if !test {
		t.Error("Same slices reported different")
	}
	if err.Error() != "input slices must be of the same length" {
		t.Error("Errors did not propagate when slices were of different size")
	}
	if test2 {
		t.Error("Different slices reported same")
	}
}
