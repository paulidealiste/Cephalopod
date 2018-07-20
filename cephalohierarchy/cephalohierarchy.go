// Package cephalohierarchy provides hierarchical clustering
package cephalohierarchy

import (
	"fmt"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephalostructures"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

// HierarchicalClustering performs said clustering and returns a list of cluster steps
func HierarchicalClustering(input *cephalobjects.DataStore) {
}

func constructTree(dmc cephalobjects.DataMatrix) {
	squarePusher := cephalostructures.Stack{}
	var dmm cephalobjects.DataMatrixExtreme
	dmm = cephaloutils.DataMatrixMin(dmc, true, false)
	squarePusher.Push(dmm)
	transformDataMatrix(&dmc, dmm)
	fmt.Println(dmc.Matrix)
}

func transformDataMatrix(dmc *cephalobjects.DataMatrix, dmm cephalobjects.DataMatrixExtreme) {
	pilaf, nlab := connectNearestLabels(dmc.Variables, dmm.RowName, dmm.ColName)
	dmc.Variables = pilaf
	inmat := make([][]float64, len(dmc.Variables)-1)
	fmt.Println(inmat)
	fmt.Println(dmc.Variables)
	fmt.Println(nlab)
}

func connectNearestLabels(varlabs []string, rc string, cc string) ([]string, string) {
	inlabs := make([]string, 0)
	nstrg := rc + "," + cc
	for _, lab := range varlabs {
		if lab != rc && lab != cc {
			inlabs = append(inlabs, lab)
		} else if lab == cc {
			inlabs = append(inlabs, nstrg)
		}
	}
	return inlabs, nstrg
}
