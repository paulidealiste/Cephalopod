package cephalofile

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephalokmeans"
	"github.com/paulidealiste/Cephalopod/cephalorandom"
)

// Whether exported file exists and contains the exported data
func TestExportDataStore(t *testing.T) {
	input, _ := cephalorandom.GenerateRandomDataStore(300, 5, 0.5)
	cephalokmeans.Kmeans(&input, 3)
	ExportDataStore(input, "../dump.json")
	op, err := os.Stat("../dump.json")
	if op == nil || err != nil {
		t.Error("File not found or generated")
	}
}

// Export cephalotimeseries to multi-series json test
func TestExportTimeseries(t *testing.T) {
	testtree := cephalobjects.NewCTS()

	ad := time.Now()
	as := time.Now()

	testtree.Insert(as, rand.Float64())

	for i := 0; i < 50; i++ {
		ad = ad.Add(10 * time.Minute)
		as = as.Add(-10 * time.Minute)
		testtree.Insert(ad, rand.Float64())
		testtree.Insert(as, rand.Float64())
	}
	ExportTimeSeries(testtree, "../tsdump.json")
	op, err := os.Stat("../tsdump.json")
	if op == nil || err != nil {
		t.Error("Export json file not found or generated")
	}

	testtree2 := testtree.RollMean(time.Minute*20, 1)

	testtreeslice := []cephalobjects.CephaloTimeSeries{testtree, testtree2}
	ExportTimeSeriesList(testtreeslice, "../tssdump.json")
	op, err = os.Stat("../tssdump.json")
	if op == nil || err != nil {
		t.Error("ExportList json file not found or generated")
	}
}

func TestExportAnnonymusCSV(t *testing.T) {
	fl := []int{1, 2, 4, 4, 5}
	in := [][]int{fl, fl, fl, fl, fl, fl, fl}
	err := ExportAnnonymusCSV(in, "boolion.csv")
	if err != nil {
		t.Error("CSV export did not work properly.")
		fmt.Println(err)
	}
}
