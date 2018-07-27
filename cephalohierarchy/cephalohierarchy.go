// Package cephalohierarchy provides hierarchical clustering
package cephalohierarchy

import (
	"strings"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephalostructures"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

// HierarchicalClustering performs said clustering and returns a list of cluster steps
func HierarchicalClustering(input *cephalobjects.DataStore) {
}

func constructTree(hirstck cephalostructures.Stack) {

}

func constructStack(dmc cephalobjects.DataMatrix) cephalostructures.Stack {
	squarePusher := cephalostructures.Stack{}
	squarePusher.Push(dmc)
	for len(dmc.Matrix[0]) > 1 {
		dmm := cephaloutils.DataMatrixMin(dmc, true, false)
		transformDataMatrix(&dmc, dmm)
		squarePusher.Push(dmc)
	}
	return squarePusher
}

func transformDataMatrix(dmc *cephalobjects.DataMatrix, dmm cephalobjects.DataMatrixExtreme) {
	pilaf, nlab := connectNearestLabels(dmc.Variables, dmm.RowName, dmm.ColName)
	inmat := make([][]float64, len(pilaf))
	ingrep := make(map[string]cephalobjects.GrepFold)
	for i, rowname := range pilaf {
		inmat[i] = make([]float64, len(pilaf))
		for j, colname := range pilaf {
			if i == j {
				inmat[i][j] = 0.0
			} else if rowname == nlab {
				inmat[i][j] = valueSingleLinkage(nlab, colname, dmc)
			} else if colname == nlab {
				inmat[i][j] = valueSingleLinkage(nlab, rowname, dmc)
			} else {
				inmat[i][j] = grepValueFromMatrix(rowname, colname, dmc)
			}
			ingrep[rowname+" "+colname] = cephalobjects.GrepFold{Row: i, Col: j}
		}
	}
	dmc.Variables = pilaf
	dmc.Matrix = inmat
	dmc.Grep = ingrep
}

func grepValueFromMatrix(rowname string, colname string, dmco *cephalobjects.DataMatrix) float64 {
	cukey := rowname + " " + colname
	cupos := dmco.Grep[cukey]
	return dmco.Matrix[cupos.Row][cupos.Col]
}

func valueSingleLinkage(nlab string, alab string, dmco *cephalobjects.DataMatrix) float64 {
	targets := strings.Split(nlab, ",")
	slv1 := grepValueFromMatrix(targets[0], alab, dmco)
	slv2 := grepValueFromMatrix(targets[1], alab, dmco)
	if slv1 <= slv2 {
		return slv1
	}
	return slv2
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
