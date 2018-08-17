// Package cephalodistance offers several distance measures and returns distance matrices
package cephalodistance

import (
	"math"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

type pair struct {
	x1, y1, x2, y2 float64
}

// CalculateDistanceMatrix fills the distance property of a datastore based on the supplied metric
func CalculateDistanceMatrix(input *cephalobjects.DataStore, metric cephalobjects.DistanceMetric) {
	var dmc cephalobjects.DataMatrix
	dmc.Matrix = make([][]float64, len(input.Basic))
	dmc.Variables = make([]string, len(input.Basic))
	dmc.Grep = make(map[string]cephalobjects.GrepFold)
	var cummulative int
	for i, dp := range input.Basic {
		dmc.Matrix[i] = make([]float64, len(input.Basic))
		dmc.Variables[i] = dp.UID
		for j, dpi := range input.Basic {
			p := pair{x1: input.Basic[i].X, y1: input.Basic[i].Y, x2: input.Basic[j].X, y2: input.Basic[j].Y}
			dmc.Matrix[i][j] = p.distance(metric, input)
			dmc.Grep[dp.UID+" "+dpi.UID] = cephalobjects.GrepFold{Row: i, Col: j}
			cummulative++
		}
	}
	input.Distance = dmc
}

func (p pair) distance(metric cephalobjects.DistanceMetric, input *cephalobjects.DataStore) float64 {
	var distC float64
	switch metric {
	case cephalobjects.Euclidean:
		distC = math.Sqrt((p.x2-p.x1)*(p.x2-p.x1) + (p.y2-p.y1)*(p.y2-p.y1))
	case cephalobjects.SquaredEuclidean:
		distC = (p.x2-p.x1)*(p.x2-p.x1) + (p.y2-p.y1)*(p.y2-p.y1)
	case cephalobjects.Manhattan:
		distC = math.Abs(p.x2-p.x1) + math.Abs(p.y2-p.y1)
	case cephalobjects.Maximum:
		distC = math.Max(math.Abs(p.x2-p.x1), math.Abs(p.y2-p.y1))
	case cephalobjects.Mahalanobis:
		distC = mahalanobis(p, input)
	}
	return distC
}

func mahalanobis(p pair, input *cephalobjects.DataStore) float64 {
	desc := cephaloutils.CalculateDescriptors(input.Basic)
	invcovmat := cephaloutils.InverseMatrix(cephaloutils.CovarianceMatrix(desc))
	p1res := []float64{(p.x1 - desc.MeanX), (p.y1 - desc.MeanY)}

	var p1ma cephalobjects.DataMatrix
	p1ma.Matrix = [][]float64{
		cephaloutils.DotProduct(invcovmat, p1res),
	}
	mahalanobis := math.Sqrt(cephaloutils.DotProduct(p1ma, p1res)[0])
	return mahalanobis
}
