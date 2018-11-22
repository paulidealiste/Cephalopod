// Package cephalobjects define global data structures
package cephalobjects

import (
	"errors"
	"time"
)

//STRTS

type CephaloTimeNode struct {
	datetime time.Time
	data     float64
	left     *CephaloTimeNode
	right    *CephaloTimeNode
}

type CephaloTimeSeries struct {
	root *CephaloTimeNode
}

//NewCTS creates an empty first-root only time series // convenience
func NewCTS() CephaloTimeSeries {
	return CephaloTimeSeries{}
}

//Insert provides a way to insert new tree node in the appropriate place
func (cts *CephaloTimeSeries) Insert(dattime time.Time, data float64) error {
	if cts.root == nil {
		cts.root = &CephaloTimeNode{datetime: dattime, data: data}
		return nil
	}
	return cts.root.insert(dattime, data)
}

//Find offers fast retrieval of the desired data point based on the supplied time
func (cts *CephaloTimeSeries) Find(dattime time.Time) (*CephaloTimeNode, bool) {
	if cts.root == nil {
		return &CephaloTimeNode{}, false
	}
	return cts.root.find(dattime)
}

//FindRange returns all of the tree nodes, i.e. tree datapoints for the requested range
func (cts *CephaloTimeSeries) FindRange(start time.Time, end time.Time) ([]*CephaloTimeNode, error) {
	var fop []*CephaloTimeNode
	if cts.root == nil || end.Before(start) {
		return fop, errors.New("No tree root element or the end is before the start")
	}
	return cts.root.findRange(start, end)
}

//Node methods considered private (insert, find)
func (ctn *CephaloTimeNode) insert(dattime time.Time, data float64) error {
	nctt := CephaloTimeNode{datetime: dattime, data: data}
	switch {
	case nctt.datetime.After(ctn.datetime):
		//Check right
		if ctn.right == nil {
			ctn.right = &nctt
			return nil
		} else {
			ctn.right.insert(nctt.datetime, nctt.data)
		}
	case nctt.datetime.Before(ctn.datetime):
		//Check left
		if ctn.left == nil {
			ctn.left = &nctt
			return nil
		} else {
			ctn.left.insert(nctt.datetime, nctt.data)
		}
	}
	return nil
}

func (ctn *CephaloTimeNode) find(dattime time.Time) (*CephaloTimeNode, bool) {
	if ctn == nil {
		return &CephaloTimeNode{}, false
	}
	switch {
	case ctn.datetime.Equal(dattime):
		return ctn, true
	case dattime.After(ctn.datetime):
		return ctn.right.find(dattime)
	default:
		return ctn.left.find(dattime)
	}
}

func (ctn *CephaloTimeNode) findRange(start time.Time, end time.Time) ([]*CephaloTimeNode, error) {
	var fop []*CephaloTimeNode
	if end.Before(start) {
		return fop, errors.New("End can't come before start in find range")
	}
	findRangeInner(ctn, start, end, func(ctn *CephaloTimeNode) {
		fop = append(fop, ctn)
	})
	return fop, nil
}

func findRangeInner(ctn *CephaloTimeNode, start time.Time, end time.Time, cb func(ctn *CephaloTimeNode)) {
	if ctn == nil {
		return
	}
	if start.Before(ctn.datetime) {
		findRangeInner(ctn.left, start, end, cb)
	}
	if (start.Before(ctn.datetime) || start.Equal(ctn.datetime)) && (end.After(ctn.datetime) || end.Equal(ctn.datetime)) {
		cb(ctn)
	}
	if end.After(ctn.datetime) {
		findRangeInner(ctn.right, start, end, cb)
	}
}
