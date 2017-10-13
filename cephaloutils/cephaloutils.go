// Package cephaloutils provides various utility functions (i.e. min, max, range, ...)
package cephaloutils

import (
	"errors"
	"math"
	"math/rand"
	"time"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephalorandom"
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

// TruncatedNormal generates truncated random normals
func TruncatedNormal(desc cephalobjects.Descriptors, l int) []cephalobjects.DataPoint {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	truncgen := make([]cephalobjects.DataPoint, l)
	upperBoundX := desc.MeanX + desc.SdX
	lowerBoundX := desc.MeanX - desc.SdX
	upperBoundY := desc.MeanY + desc.SdY
	lowerBoundY := desc.MeanY - desc.SdY
	for i := range truncgen {
		for {
			truncgen[i].X = math.Abs(random.NormFloat64())*desc.SdX + desc.MeanX
			truncgen[i].Y = math.Abs(random.NormFloat64())*desc.SdY + desc.MeanY
			truncgen[i].G = cephalorandom.RandStringBytes(random, 5)
			truncgen[i].A = truncgen[i].G
			if truncgen[i].X > lowerBoundX && truncgen[i].X < upperBoundX && truncgen[i].Y > lowerBoundY && truncgen[i].Y < upperBoundY {
				break
			}
		}
	}
	return truncgen
}

// EuclideanDistance returns the L2 norm of two DataPoints
func EuclideanDistance(p1 cephalobjects.DataPoint, p2 cephalobjects.DataPoint) float64 {
	ed := math.Sqrt(math.Pow((p1.X-p2.X), 2) + math.Pow((p1.Y-p2.Y), 2))
	return ed
}

// MinSliceIndex returns the position of the slice's smallest element
func MinSliceIndex(input []float64) int {
	m := input[0]
	var mi int
	for i, e := range input {
		if e < m {
			m = e
			mi = i
		}
	}
	return mi
}

// CheckIfSame checks if all the values in the input arrays are equal
func CheckIfSame(s1 []cephalobjects.DataPoint, s2 []cephalobjects.DataPoint) (bool, error) {
	if len(s1) != len(s2) {
		err := errors.New("input slices must be of the same length")
		return false, err
	}
	var counter int
	for i := range s1 {
		if s1[i].X == s2[i].X && s1[i].Y == s2[i].Y {
			counter++
		}
	}
	if counter == len(s1) {
		return true, nil
	}
	return false, nil
}
