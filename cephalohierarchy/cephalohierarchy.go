// Package cephalohierarchy provides hierarchical clustering
package cephalohierarchy

import (
	"fmt"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

// HierarchicalClustering performs said clustering and returns a list of cluster steps
func HierarchicalClustering(input *cephalobjects.DataStore) {
}

func constructTree(dmc cephalobjects.DataMatrix) {
	var dmi cephalobjects.DataMatrix
	dmi.Variables = make([]string, len(dmc.Variables))
	dmm := cephaloutils.DataMatrixMin(dmc, true, false)
	for {
		if len(dmi.Variables) > 1 {
			fmt.Println(dmm)
			// dmm = cephaloutils.DataMatrixMin(dmi, true, false)
		} else {
			//Should break here
		}
		break
	}
}
