// Package cephalohierarchy provides hierarchical clustering
package cephalohierarchy

import (
	"strings"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephalodistance"
	"github.com/paulidealiste/Cephalopod/cephalolambdas"
	"github.com/paulidealiste/Cephalopod/cephalostructures"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

// HierarchicalClustering performs said clustering and returns its graph representation
func HierarchicalClustering(input *cephalobjects.DataStore, linkage cephalobjects.LinkageCriteria) cephalostructures.Graph {
	cephalodistance.CalculateDistanceMatrix(input, cephalobjects.Euclidean)
	hs := constructStack(input.Distance, linkage)
	hg := constructGraph(hs)
	return hg
}

func constructGraph(hirstck cephalostructures.Stack) cephalostructures.Graph {
	hirgraph := cephalostructures.Graph{}
	var hirnodes []string
	var prevnodes []string
	for !hirstck.Empty() {
		po := hirstck.Pop()
		if t, ok := po.(cephalobjects.DataMatrix); ok { // Two-return syntax fot type assertion
			for _, dex := range t.Variables {
				hirgraph.InsertNode(dex, dex, po)
				hirnodes = append(hirnodes, dex)
			}
			ne := cephalolambdas.Beakdiff(hirnodes, prevnodes)
			od := cephalolambdas.Beakdiff(prevnodes, ne)
			prevnodes = hirnodes
			weaveEdge(&hirgraph, ne, od)
		}
	}
	return hirgraph
}

func weaveEdge(gp *cephalostructures.Graph, new []string, old []string) {
	cat := gp.GraphCatalog()
	for _, nnode := range new {
		bfufn := func(s string) bool { return strings.Contains(s, nnode) }
		cntns := cephalolambdas.Beakfilter(old, bfufn)
		shrtr := cephaloutils.ShortestString(cntns)
		if cat[shrtr] != nil {
			gp.DirectedEdge(cat[shrtr], cat[nnode])
		}
	}
}

func constructStack(dmc cephalobjects.DataMatrix, linkage cephalobjects.LinkageCriteria) cephalostructures.Stack {
	squarePusher := cephalostructures.Stack{}
	squarePusher.Push(dmc)
	for len(dmc.Matrix[0]) > 1 {
		dmm := cephaloutils.DataMatrixMin(dmc, true, false)
		transformDataMatrix(&dmc, dmm, linkage)
		squarePusher.Push(dmc)
	}
	return squarePusher
}

func transformDataMatrix(dmc *cephalobjects.DataMatrix, dmm cephalobjects.DataMatrixExtreme, linkage cephalobjects.LinkageCriteria) {
	pilaf, nlab := connectNearestLabels(dmc.Variables, dmm.RowName, dmm.ColName)
	inmat := make([][]float64, len(pilaf))
	ingrep := make(map[string]cephalobjects.GrepFold)
	for i, rowname := range pilaf {
		inmat[i] = make([]float64, len(pilaf))
		for j, colname := range pilaf {
			if i == j {
				inmat[i][j] = 0.0
			} else if rowname == nlab {
				inmat[i][j] = linkageFunction(nlab, colname, dmc, linkage)
			} else if colname == nlab {
				inmat[i][j] = linkageFunction(nlab, rowname, dmc, linkage)
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

func linkageFunction(nlab string, alab string, dmco *cephalobjects.DataMatrix, linkage cephalobjects.LinkageCriteria) float64 {
	var dist float64
	targets := strings.Split(nlab, ",")
	slv1 := grepValueFromMatrix(targets[0], alab, dmco)
	slv2 := grepValueFromMatrix(targets[1], alab, dmco)
	switch linkage {
	case cephalobjects.Single:
		dist = singleComparator(slv1, slv2)
	case cephalobjects.Complete:
		dist = completeComparator(slv1, slv2)
	case cephalobjects.Average:
		dist = averageComparator(slv1, slv2)
	}
	return dist
}

func singleComparator(slv1 float64, slv2 float64) float64 {
	if slv1 <= slv2 {
		return slv1
	}
	return slv2
}

func completeComparator(slv1 float64, slv2 float64) float64 {
	if slv1 > slv2 {
		return slv1
	}
	return slv2
}

func averageComparator(slv1 float64, slv2 float64) float64 {
	return (slv1 + slv2) / 2
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
