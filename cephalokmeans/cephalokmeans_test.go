package cephalokmeans

import (
	"testing"

	"github.com/paulidealiste/Cephalopod/cephalorandom"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

// wheter generated centroids fall within data bounds and wheter their number is the same as the number of groups
func TestCentroidsGeneratorSpread(t *testing.T) {
	k := 3
	input, _ := cephalorandom.GenerateRandomDataStore(126, 5, 0.5)
	ranger := cephaloutils.ExtremesRange(&input)
	test := generateCentroids(&input, k)
	if len(test) != k {
		t.Error("Did not generate the adequate ammount of centroids")
	}
	for _, ce := range test {
		if ce.X < ranger[0].X || ce.Y < ranger[0].Y {
			t.Error("Generated centroids escaped lower bound of the data range")
		}
		if ce.X > ranger[1].X || ce.Y > ranger[1].Y {
			t.Error("Generated centroids escaped upper bound of the data range")
		}
	}
}
