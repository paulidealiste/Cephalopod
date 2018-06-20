package cephalohierarchy

import (
	"testing"

	"github.com/paulidealiste/Cephalopod/cephalorandom"
)

func TestHierarchicalClustering(t *testing.T) {
	input, _ := cephalorandom.GenerateRandomDataStore(20, 3, 0.5)
	HierarchicalClustering(&input)
	input.Distance = [][]float64{
		{0.00, 0.71, 5.66, 3.61, 4.24, 3.20},
		{0.71, 0.00, 4.95, 2.92, 3.54, 2.50},
		{5.66, 4.95, 0.00, 2.24, 1.41, 2.50},
		{3.61, 2.92, 2.24, 0.00, 1.00, 0.50},
		{4.24, 3.54, 1.41, 1.00, 0.00, 1.12},
		{3.20, 2.50, 2.50, 0.50, 1.12, 0.00},
	}
}