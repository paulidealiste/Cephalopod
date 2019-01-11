package cephalofile

import (
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
	ExportDataStore(input)
	op, err := os.Stat("../dump.json")
	if op == nil || err != nil {
		t.Error("File not found or generated")
	}
}

func TestExportTimeseries(t *testing.T) {
	testtree := cephalobjects.NewCTS()
	ad := time.Now()
	as := time.Now()
	for i := 0; i < 100; i++ {
		ad = ad.Add(10 * time.Minute)
		as = as.Add(-10 * time.Minute)
		testtree.Insert(ad, rand.Float64())
		testtree.Insert(as, rand.Float64())
	}
	ExportTimeSeries(testtree)
}

// Is Iris data properly read from a .csv
func TestImportCSV(t *testing.T) {
	// rawiris := ImportCSV("iris.csv")
	// fmt.Println(rawiris[0])
	// if len(rawiris) == 0 {
	// 	t.Error("Probarbly not imported at all")
	// }
}
