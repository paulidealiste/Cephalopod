// Package cephalokmeans provides fast k-means clustering
package cephalokmeans

import (
	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

func generateCentroids(input *cephalobjects.DataStore, k int) []cephalobjects.DataPoint {
	datarange := cephaloutils.ExtremesRange(input)
	descriptors := cephaloutils.CalculateDescriptors(datarange)
	centroids := cephaloutils.TruncatedNormal(descriptors, k)
	return centroids
}

func assignCentroids(input *cephalobjects.DataStore, centroids []cephalobjects.DataPoint) {
	for i, dp := range input.Basic {
		dists := make([]float64, len(centroids))
		for i, cp := range centroids {
			dists[i] = cephaloutils.EuclideanDistance(dp, cp)
		}
		input.Basic[i].G = centroids[cephaloutils.MinSliceIndex(dists)].G
	}
}
