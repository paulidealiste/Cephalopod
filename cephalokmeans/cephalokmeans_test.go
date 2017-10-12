package cephalokmeans

import (
	"testing"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
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
		if ce.X < ranger[0].X*2 || ce.Y < ranger[0].Y*2 {
			t.Error("Generated centroids escaped lower bound of the data range (x2)")
		}
		if ce.X > ranger[1].X*2 || ce.Y > ranger[1].Y*2 {
			t.Error("Generated centroids escaped upper bound of the data range (x2)")
		}
	}
}

// did centoid assignment assign each data point to a centroid-based group and whether centroids are recalculated accordingly
func TestCentroidAssignmentAndRecalculation(t *testing.T) {
	k := 3
	input, _ := cephalorandom.GenerateRandomDataStore(120, 3, 0.5)
	centroids := generateCentroids(&input, k)
	centroidsBak := make([]cephalobjects.DataPoint, len(centroids))
	copy(centroidsBak, centroids)
	assignCentroids(&input, centroids)
	recalculateCentroids(input, centroids)
	for _, dp := range input.Basic {
		if dp.G == "" {
			t.Error("Not all data points were assigned to corresponding centroids")
		}
	}
	var counter int
	for i := range centroids {
		if centroids[i].X == centroidsBak[i].X && centroids[i].Y == centroidsBak[i].Y {
			counter++
		}
	}
	if counter == len(centroids) {
		t.Error("Centroids were probably not recalculated")
	}
}
