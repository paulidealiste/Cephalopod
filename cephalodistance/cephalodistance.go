// Package cephalodistance offers several distance measures and returns distance matrices
package cephalodistance

import (
	"math"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
)

type pair struct {
	x1, y1, x2, y2 float64
}

// CalculateDistanceMatrix fills the distance property of a datastore based on the supplied metric
func CalculateDistanceMatrix(input *cephalobjects.DataStore, metric cephalobjects.DistanceMetric) {
	input.Distance = make([][]float64, len(input.Basic))
	for i := range input.Basic {
		input.Distance[i] = make([]float64, len(input.Basic))
		for j := range input.Basic {
			p := pair{x1: input.Basic[i].X, y1: input.Basic[i].Y, x2: input.Basic[j].X, y2: input.Basic[j].Y}
			input.Distance[i][j] = p.distance(metric)
		}
	}
}

func (p pair) distance(metric cephalobjects.DistanceMetric) float64 {
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
		distC = 0.0
	}
	return distC
}
