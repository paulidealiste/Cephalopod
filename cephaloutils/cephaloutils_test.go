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
	dmc.Grep = make(map[string]cephalobjects.GrepFold)
	dmc.Variables = []string{"ohgr", "cevin", "skinn", "puppy", "remis"}
	// dmc.Matrix = [][]float64{
	// 	{0.00, 0.71, 5.66, 3.61, 4.24, 3.20},
	// 	{0.71, 0.00, 4.95, 2.92, 3.54, 2.50},
	// 	{5.66, 4.95, 0.00, 2.24, 1.41, 2.50},
	// 	{3.61, 2.92, 2.24, 0.00, 1.00, 0.50},
	// 	{4.24, 3.54, 1.41, 1.00, 0.00, 1.12},
	// 	{3.20, 2.50, 2.50, 0.50, 1.12, 0.00},
	// }
	dmc.Matrix = [][]float64{
		{0.00, 0.71, 5.66, 3.20, 4.24},
		{0.71, 0.00, 4.95, 2.50, 3.54},
		{5.66, 4.95, 0.00, 2.24, 1.41},
		{3.20, 2.50, 2.24, 0.00, 1.00},
		{4.24, 3.54, 1.41, 1.00, 0.00},
	}
	for i, vn := range dmc.Variables {
		for j, vni := range dmc.Variables {
			dmc.Grep[vn+" "+vni] = cephalobjects.GrepFold{Row: i, Col: j}
		}
	}

	test := DataMatrixMin(dmc, true, false)
	test1 := DataMatrixMin(dmc, false, true)
	fmt.Println("Datamatrix extremes")
	fmt.Println(test)
	fmt.Println(test1)

}

// Do you really know which string is the shortest?
func TestShortestString(t *testing.T) {
	teststrings := []string{"remork", "anak", "sterlinz", "amu", "pretkola"}
	test := ShortestString(teststrings)
	if test != "amu" {
		t.Error("Haven't found the shortest string element")
	}
}

//
// Timeseries utilities testing
//

// Does traversal ts -> list really work?

func TestTSListFormAndMap(t *testing.T) {
	testtree := cephalobjects.NewCTS()
	ad := time.Now()
	as := time.Now()
	for i := 0; i < 100; i++ {
		ad = ad.Add(10 * time.Minute)
		as = as.Add(-10 * time.Minute)
		testtree.Insert(ad, rand.Float64())
		testtree.Insert(as, rand.Float64())
	}
	testTSList := CTSListForm(testtree)
	if len(testTSList.Data) != testtree.Size {
		t.Error("Probably not an accurate traversal list")
	}
	testTSMap := CTSListMap(testtree)
	if len(testTSMap[testtree.ID].Data) != testtree.Size {
		t.Error("Probably not an accurate traversal map")
	}
}

func TestTSListsFromTSTrees(t *testing.T) {
	testtree := cephalobjects.NewCTS()
	ad := time.Now()
	as := time.Now()
	for i := 0; i < 1000; i++ {
		ad = ad.Add(10 * time.Minute)
		as = as.Add(-10 * time.Minute)
		testtree.Insert(ad, rand.Float64())
		testtree.Insert(as, rand.Float64())
	}
	testtree2 := cephalobjects.NewCTS()
	ad = time.Now()
	as = time.Now()
	for i := 0; i < 1000; i++ {
		ad = ad.Add(20 * time.Minute)
		as = as.Add(-20 * time.Minute)
		testtree.Insert(ad, rand.Float64())
		testtree.Insert(as, rand.Float64())
	}

	testtrees := []cephalobjects.CephaloTimeSeries{testtree, testtree2}
	defer Elapsed("MultiMapFromTree")()
	test := TSListsFromTSTrees(testtrees)
	if len(test[testtree.ID].Data) != testtree.Size {
		t.Error("Does the multi map really exist?")
	}
}
