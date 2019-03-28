package cephalobjects

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCTSforNodes(t *testing.T) {

	fmt.Println("Node tests")

	test := CephaloTimeNode{Datetime: time.Now(), Data: rand.Float64()}

	//Test targets
	dd := test.Datetime
	ds := test.Datetime
	ad := test.Datetime
	as := test.Datetime

	//Loop based insert
	for i := 0; i < 100; i++ {
		ad = ad.Add(10 * time.Minute)
		as = as.Add(-10 * time.Minute)
		test.insert(ad, rand.Float64())
		test.insert(as, rand.Float64())
	}

	//Find exact match
	if found, state := test.find(dd.Add(10 * time.Minute)); state != false {
		fmt.Println(found)
	}
	if found1, state1 := test.find(ds.Add(-10 * time.Minute)); state1 != false {
		fmt.Println(found1)
	}

	//Find range match
	_, err := test.findRange(dd.Add(10*time.Minute), ds.Add(-10*time.Minute))
	if err == nil {
		t.Error("Should have failed when end is before start")
	}
	rfound, _ := test.findRange(ds.Add(-10*time.Minute), dd.Add(10*time.Minute))
	fmt.Println(len(rfound))

	//Find max/min Datetime CephaloTimeNode
	max, parentMax := test.findMax(nil)
	min, parentMin := test.findMin(nil)
	fmt.Println("Max and parent")
	fmt.Println(max.Datetime)
	fmt.Println(parentMax.Datetime)
	fmt.Println("Min and parent")
	fmt.Println(min.Datetime)
	fmt.Println(parentMin.Datetime)

	//Delete tests
	test.delete(parentMax.Datetime, nil) //Delete half-leaf node with right child
	test.delete(max.Datetime, nil)       //Delete leaf node -- max is re-declared as its parent after the delete
	newmax, _ := test.findMax(nil)
	fmt.Println(newmax.Datetime)         //Since previous max is deleted new one can be obtained
	test.delete(parentMin.Datetime, nil) //Delete half-leaf node with left child

	test.delete(dd, nil)

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
	//Delete test when no root node
	err := testtree.Delete(ad)
	if err == nil {
		t.Error("Well not checking for the root node aren't we")
	}

	testtree.Insert(ad, rand.Float64())
	testtree.Delete(ad) // Delete test when only a root node

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
	fmt.Println(rfound)

	err = testtree.Delete(dd)
	fmt.Println(err)

	fmt.Println("Tree traversal")
	testtree.TraversalMap(testtree.Root, func(ctn *CephaloTimeNode) {
	})
	fmt.Println("Tree traversal with endpoints")
	testtree.EndpointsMap(time.Minute*21, testtree.Root, func(ctn *CephaloTimeNode) {
		// fmt.Println(ctn.Datetime)
	})

	fmt.Println("PeriodApplication")
	meantree := testtree.PeriodApply(time.Minute*20, testtree.Root, func(runo []*CephaloTimeNode) CephaloTimeNode {
		newnode := CephaloTimeNode{Datetime: runo[0].Datetime, Data: runo[0].Data}
		return newnode
	})
	fmt.Println(meantree.Size)
	fmt.Println("RollingApplication")
	rolltree := testtree.RollApply(time.Minute*20, testtree.Root, 1, func(nuno []*CephaloTimeNode) CephaloTimeNode {
		newnode := CephaloTimeNode{Datetime: nuno[0].Datetime, Data: nuno[0].Data}
		return newnode
	})
	fmt.Println(rolltree.Size)
	fmt.Println("RollingMeanApplication")
	meanrolltree := testtree.RollMean(time.Minute*20, 1)
	fmt.Println(meanrolltree.Size)
	fmt.Println("End of tree tests")

}

func TestAbsDuration(t *testing.T) {
	test := AbsDuration(-time.Minute * 21)
	if test < 0 {
		t.Error("Not really an absolute value is it?")
	}
}
