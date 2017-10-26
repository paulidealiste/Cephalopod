// Package cephalokmeans provides fast k-means clustering
package cephalokmeans

import (
	"math"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

func generateCentroids(input *cephalobjects.DataStore, k int) []cephalobjects.DataPoint {
	datarange := cephaloutils.ExtremesRange(input)
	descriptors := cephaloutils.CalculateDescriptors(datarange)
	centroids := cephaloutils.TruncatedNormal(descriptors, k)
	return centroids
}

func assignCentroids(input *cephalobjects.DataStore, centroids []cephalobjects.DataPoint) bool {
	checking := make([]bool, 0)
	for i, dp := range input.Basic {
		dists := make([]float64, len(centroids))
		for i, cp := range centroids {
			dists[i] = cephaloutils.EuclideanDistance(dp, cp)
		}
		var tempG = input.Basic[i].G
		input.Basic[i].G = centroids[cephaloutils.MinSliceIndex(dists)].G
		checking = append(checking, tempG == input.Basic[i].G)
	}
	return cephaloutils.CheckAllTrue(checking)
}

func recalculateCentroids(input cephalobjects.DataStore, centroids []cephalobjects.DataPoint) {
	storage := make(map[string][]cephalobjects.DataPoint)
	for _, dp := range input.Basic {
		storage[dp.G] = append(storage[dp.G], dp)
	}
	for i, cp := range centroids {
		descriptors := cephaloutils.CalculateDescriptors(storage[cp.G])
		if !math.IsNaN(descriptors.MeanX) && !math.IsNaN(descriptors.MeanY) {
			centroids[i].X = descriptors.MeanX
			centroids[i].Y = descriptors.MeanY
		}
	}
}

// Kmeans performs the clustering algorhithm and assigns clusters to a DataStore
func Kmeans(input *cephalobjects.DataStore, k int) {
	maxiter := 200
	curriter := 0
	centroids := generateCentroids(input, k)
	for {
		recalculateCentroids(*input, centroids)
		var noChanges = assignCentroids(input, centroids)
		curriter++
		if noChanges || curriter == maxiter {
			break
		}
	}
}
