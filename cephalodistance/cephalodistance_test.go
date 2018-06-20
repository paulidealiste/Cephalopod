package cephalodistance

import (
	"testing"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephalorandom"
)

func TestCalculateDistanceMatrix(t *testing.T) {
	input, _ := cephalorandom.GenerateRandomDataStore(10, 3, 0.5)
	CalculateDistanceMatrix(&input, cephalobjects.Euclidean)
	CalculateDistanceMatrix(&input, cephalobjects.SquaredEuclidean)
	CalculateDistanceMatrix(&input, cephalobjects.Manhattan)
	CalculateDistanceMatrix(&input, cephalobjects.Maximum)
	CalculateDistanceMatrix(&input, cephalobjects.Mahalanobis)
}
