package cephaloutils

import (
	"fmt"
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
	fmt.Println(desc)
	covmat := CovarianceMatrix(desc)
	fmt.Println(covmat)
	InverseMatrix(covmat)
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

// Tests the knownd matrix vector product

func TestDotProduct(t *testing.T) {
	testvec := []float64{-2, 40, 4}
	var testmat cephalobjects.DataMatrix
	testmat.Matrix = [][]float64{
		{3.6885, 0.0627, -1.2821},
		{0.0627, 0.0022, -0.024},
		{-1.2821, -0.024, 0.4588},
	}
	product := DotProduct(testmat, testvec)
	prodVec := []float64{-9.9974, -0.1334, 3.4394}
	if product[2] != prodVec[2] {
		t.Error("Product not valid")
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

// Does plucking really returns string values from DataPoint objects (tests A instead of G)
func TestPluckStringValues(t *testing.T) {
	input, _ := cephalorandom.GenerateRandomDataStore(12, 3, 0.5)
	test := PluckStringValues(input.Basic, cephalobjects.Actual)
	test2 := PluckStringValues(input.Basic, cephalobjects.Grouped)
	if test[0] != input.Basic[0].A && test[1] != input.Basic[1].A || test2[0] != input.Basic[0].G && test2[1] != input.Basic[1].G {
		t.Error("String values were not plucked correctly")
	}
}

// Whether all values in the input array are true
func TestCheckAllTrue(t *testing.T) {
	input := []bool{true, true, true, false}
	test := CheckAllTrue(input)
	if test == true {
		t.Error("Checking for all true not true")
	}
}

// Does the min function really returns the minimal value info
func TestDataMatrixMin(t *testing.T) {
	var dmc cephalobjects.DataMatrix
	dmc.Grep = make(map[string]int)
	dmc.Variables = []string{"ohgr", "cevin", "puppy"}
	dmc.Matrix = [][]float64{
		{10.00, 0.71, 3.73},
		{0.71, 10.00, -1.87},
		{3.73, -1.87, 10.00},
	}
	var cummulative int
	for _, vn := range dmc.Variables {
		for _, vni := range dmc.Variables {
			dmc.Grep[vn+" "+vni] = cummulative
			cummulative++
		}
	}

	test := DataMatrixMin(dmc, true)
	test1 := DataMatrixMin(dmc, false)
	fmt.Println("Datamatrix extremes")
	fmt.Println(test)
	fmt.Println(test1)

}
