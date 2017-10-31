// Package cephalolinreg performs linear regression between arrays of DataPoint's x and y
package cephalolinreg

import (
	"math"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

func hypothesis(a float64, b float64, x float64) float64 {
	return a + b*x
}

func leastSquares(input []cephalobjects.DataPoint) cephalobjects.ModelSummary {
	var summary cephalobjects.ModelSummary
	desc := cephaloutils.CalculateDescriptors(input)
	ssdX := 0.0
	ssdY := 0.0
	ssdXY := 0.0
	for _, dp := range input {
		ssdX += math.Pow((dp.X - desc.MeanX), 2)
		ssdY += math.Pow((dp.Y - desc.MeanY), 2)
		ssdXY += (dp.X - desc.MeanX) * (dp.Y - desc.MeanY)
	}
	summary.B = ssdXY / ssdX
	summary.A = desc.MeanY - (desc.MeanX * summary.B)
	summary.R2 = math.Pow(ssdXY, 2) / (ssdX * ssdY)
	return summary
}

func cost(input []cephalobjects.DataPoint, modPar cephalobjects.ModelSummary) float64 {
	sum := 0.0
	for _, dp := range input {
		sum += math.Pow(hypothesis(modPar.A, modPar.B, dp.X)-dp.Y, 2)
	}
	return sum / float64(2.0*len(input))
}
