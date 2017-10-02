// Package cephalokmeans provides fast k-means clustering
package cephalokmeans

import (
	"math/rand"
	"time"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephalorandom"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

func generateCentroids(input *cephalobjects.DataStore, k int) []cephalobjects.DataPoint {
	centroids := make([]cephalobjects.DataPoint, k)
	datarange := cephaloutils.ExtremesRange(input)
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	descriptors := cephaloutils.CalculateDescriptors(&datarange)
	for i := range centroids {
		centroids[i].X = random.NormFloat64()
		centroids[i].Y = random.NormFloat64()
		centroids[i].G = cephalorandom.RandStringBytes(random, 5)
	}
	return centroids
}
