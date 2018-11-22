package cephalobjects

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCTSforNodes(t *testing.T) {

	fmt.Println("Node tests")

	test := CephaloTimeNode{}

	//Test targets
	var dd time.Time
	var ds time.Time
	ad := time.Now()
	as := time.Now()

	//Loop based insert
	for i := 0; i < 100; i++ {
		if i == 1 {
			dd = ad
			ds = as
		}
		ad = ad.Add(10 * time.Minute)
		as = as.Add(-10 * time.Minute)
		test.insert(ad, rand.Float64())
		test.insert(as, rand.Float64())
	}

	//Find exact match
	if found, state := test.find(dd); state != false {
		fmt.Println(found)
	}
	if found1, state1 := test.find(ds); state1 != false {
		fmt.Println(found1)
	}

	//Find range match
	_, err := test.findRange(dd.Add(10*time.Minute), ds.Add(-10*time.Minute))
	if err == nil {
		t.Error("Should have failed when end is before start")
	}
	rfound, _ := test.findRange(ds.Add(-5*time.Minute), dd.Add(5*time.Minute))
	fmt.Println(rfound[0].datetime)

	fmt.Println("End of node tests")
}

func TestCTSforTimeSeries(t *testing.T) {

	fmt.Println("Tree tests")

	testtree := NewCTS()

	//Test targets
	var dd time.Time
	var ds time.Time
	ad := time.Now()
	as := time.Now()

	//Find test when no root node
	if found, state := testtree.Find(ad); state != false {
		fmt.Println(found)
	}
	rfoundBlac, _ := testtree.FindRange(ds.Add(-5*time.Minute), dd.Add(5*time.Minute))
	if len(rfoundBlac) != 0 {
		t.Error("Probably a bad error for find range of timeseries")
	}

	//Loop tests for timeseries creation
	for i := 0; i < 100; i++ {
		if i == 1 {
			dd = ad
			ds = as
		}
		ad = ad.Add(10 * time.Minute)
		as = as.Add(-10 * time.Minute)
		testtree.Insert(ad, rand.Float64())
		testtree.Insert(as, rand.Float64())
	}

	if found1, state1 := testtree.Find(ds); state1 != false {
		fmt.Println(found1)
	}

	rfound, _ := testtree.FindRange(ds.Add(-5*time.Minute), dd.Add(5*time.Minute))
	fmt.Println(rfound[0].datetime)

	fmt.Println("End of tree tests")

}
