package cephalohierarchy

import (
	"fmt"
	"testing"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephalorandom"
)

func TestHierarchicalClustering(t *testing.T) {
	input, _ := cephalorandom.GenerateRandomDataStore(12, 3, 0.5)
	//
	fmt.Println("Dum dummy tests!")
	HierarchicalClustering(&input, cephalobjects.Single)
	HierarchicalClustering(&input, cephalobjects.Complete)
	HierarchicalClustering(&input, cephalobjects.Average)
	//
	// var dmc cephalobjects.DataMatrix
	// dmc.Grep = make(map[string]cephalobjects.GrepFold)
	// dmc.Variables = []string{"ohgr", "cevin", "skinn", "puppy", "remis", "bites"}
	// dmc.Matrix = [][]float64{
	// 	{0.00, 0.71, 5.66, 3.61, 4.24, 3.20},
	// 	{0.71, 0.00, 4.95, 2.92, 3.54, 2.50},
	// 	{5.66, 4.95, 0.00, 2.24, 1.41, 2.50},
	// 	{3.61, 2.92, 2.24, 0.00, 1.00, 0.50},
	// 	{4.24, 3.54, 1.41, 1.00, 0.00, 1.12},
	// 	{3.20, 2.50, 2.50, 0.50, 1.12, 0.00},
	// }
	// for i, vn := range dmc.Variables {
	// 	for j, vni := range dmc.Variables {
	// 		dmc.Grep[vn+" "+vni] = cephalobjects.GrepFold{Row: i, Col: j}
	// 	}
	// }
	// hirstck := constructStack(dmc)
	// constructGraph(hirstck)
}
