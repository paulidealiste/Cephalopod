// Package cephaloutils provides various utility functions (i.e. min, max, range, ...)
package cephaloutils

import (
	"math"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
)

// ExtremesRange returns X and Y coordinates range
func ExtremesRange(input *cephalobjects.DataStore) []cephalobjects.DataPoint {
	extremes := make([]cephalobjects.DataPoint, 2)
	var minX, minY float64 = input.Basic[0].X, input.Basic[0].Y
	var maxX, maxY float64 = 0.0, 0.0
	for _, dp := range input.Basic {
		minX = math.Min(minX, dp.X)
		minY = math.Min(minY, dp.Y)
		maxX = math.Max(maxX, dp.X)
		maxY = math.Max(maxY, dp.Y)
	}
	extremes[0].X = minX
	extremes[0].Y = minY
	extremes[1].X = maxX
	extremes[1].Y = maxY
	return extremes
}

// CalculateDescriptors returns means and SDs of a DataPoint slice
func CalculateDescriptors(input []cephalobjects.DataPoint) cephalobjects.Descriptors {
	var meanX, meanY, sdX, sdY float64
	var sumX, sumY, ssX, ssY float64 = 0.0, 0.0, 0.0, 0.0
	li := float64(len(input))
	for _, dp := range input {
		sumX += dp.X
		sumY += dp.Y
	}
	meanX = sumX / li
	meanY = sumY / li
	for _, dp := range input {
		ssX += (dp.X - meanX) * (dp.X - meanX)
		ssY += (dp.Y - meanY) * (dp.Y - meanY)
	}
	sdX = math.Sqrt(ssX / (li - 1))
	sdY = math.Sqrt(ssY / (li - 1))
	descs := cephalobjects.Descriptors{
		MeanX: meanX,
		MeanY: meanY,
		SdX:   sdX,
		SdY:   sdY,
	}
	return descs
}
