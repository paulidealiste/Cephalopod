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
	for _, dp := range input.Basic {
		var td, tdp float64
		for _, cp := range centroids {
			td = cephaloutils.EuclideanDistance(dp, cp)
			if td < tdp {
				dp.G = cp.G
			}
			tdp = td
		}
	}
}
