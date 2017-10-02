// Package cephalokmeans provides fast k-means clustering
package cephalokmeans

import (
	"math"
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
	descriptors := cephaloutils.CalculateDescriptors(datarange)
	for i := range centroids {
		centroids[i].X = math.Abs(random.NormFloat64())*descriptors.SdX + descriptors.MeanX
		centroids[i].Y = math.Abs(random.NormFloat64())*descriptors.SdY + descriptors.MeanY
		centroids[i].G = cephalorandom.RandStringBytes(random, 5)
	}
	return centroids
}
