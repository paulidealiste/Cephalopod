// Package cephalobjects define global Data structures
package cephalobjects

import (
	"errors"
	"time"
)

//STRTS

type CephaloTimeNode struct {
	Datetime time.Time
	Data     float64
	left     *CephaloTimeNode
	right    *CephaloTimeNode
}

type CephaloTimeSeries struct {
	Size int
	Root *CephaloTimeNode
}

//NewCTS creates an empty first-Root only time series // convenience
func NewCTS() CephaloTimeSeries {
	return CephaloTimeSeries{}
}

//Insert provides a way to insert new tree node in the appropriate place
func (cts *CephaloTimeSeries) Insert(dattime time.Time, Data float64) error {
	if cts.Root == nil {
		cts.Root = &CephaloTimeNode{Datetime: dattime, Data: Data}
		cts.Size++
		return nil
	}
	cts.Size++
	return cts.Root.insert(dattime, Data)
}

//Find offers fast retrieval of the desired Data point based on the supplied time
func (cts *CephaloTimeSeries) Find(dattime time.Time) (*CephaloTimeNode, bool) {
	if cts.Root == nil {
		return &CephaloTimeNode{}, false
	}
	return cts.Root.find(dattime)
}

//FindRange returns all of the tree nodes, i.e. tree Datapoints for the requested range
func (cts *CephaloTimeSeries) FindRange(start time.Time, end time.Time) ([]*CephaloTimeNode, error) {
	var fop []*CephaloTimeNode
	if cts.Root == nil || end.Before(start) {
		return fop, errors.New("No tree Root element or the end is before the start")
	}
	return cts.Root.findRange(start, end)
}

//Delete removes the designated Datapoint from the time series tree
func (cts *CephaloTimeSeries) Delete(dattime time.Time) error {
	if cts.Root == nil {
		return errors.New("Deletion can not be performed in an empty tree")
	}
	fakeParent := &CephaloTimeNode{right: cts.Root}
	err := cts.Root.delete(dattime, fakeParent)
	if err != nil {
		return err
	}
	if fakeParent.right == nil {
		cts.Root = nil
	}
	cts.Size--
	return nil
}

//TraversalMap offers inorder traversal of the timeseries with a callback/map function returning a Datapoint pointer
func (cts *CephaloTimeSeries) TraversalMap(ctn *CephaloTimeNode, callback func(*CephaloTimeNode)) {
	if ctn == nil {
		return
	}
	cts.TraversalMap(ctn.left, callback)
	callback(ctn)
	cts.TraversalMap(ctn.right, callback)
}

//EndpointsMap offers inorder traversal along the specified duration units
//(corresponding to last observation within duration unit) along with a callback/map function
func (cts *CephaloTimeSeries) EndpointsMap(period time.Duration, ctn *CephaloTimeNode, callback func(*CephaloTimeNode)) {
	runningnode, _ := ctn.findMin(nil)
	cts.TraversalMap(ctn, func(current *CephaloTimeNode) {
		controltime := runningnode.Datetime.Add(period)
		if current.Datetime.Equal(controltime) || current.Datetime.After(controltime) {
			callback(ctn)
			runningnode = current
		}
	})
}

//PeriodApply perfoms the series partial application of the supplied function, thus returning
//an enetierly new series of the period endpoints length (last duration unit) with Data transformed
//accordingly. It is expected that the applied function returns a new CephaloTimeNode instead of
//a pointer to an already used one
func (cts *CephaloTimeSeries) PeriodApply(period time.Duration, ctn *CephaloTimeNode, applied func([]*CephaloTimeNode) CephaloTimeNode) CephaloTimeSeries {
	nts := NewCTS()
	runningnode, _ := ctn.findMin(nil)
	var runnernodes []*CephaloTimeNode
	cts.TraversalMap(ctn, func(current *CephaloTimeNode) {
		controltime := runningnode.Datetime.Add(period)
		if current.Datetime.Equal(controltime) || current.Datetime.After(controltime) {
			calcnode := applied(runnernodes)
			nts.Insert(calcnode.Datetime, calcnode.Data)
			runnernodes = nil
			runningnode = current
		}
		runnernodes = append(runnernodes, current)
	})
	return nts
}

//Window-based methods (using FindRange for start-end window extraction)

//RollApply provides duration-based rolling window application of the callback function.
//Upon finishing it returns new CephaloTimeSeries, nearest resampled to the provided duration
func (cts *CephaloTimeSeries) RollApply(period time.Duration, ctn *CephaloTimeNode, minn int, applied func([]*CephaloTimeNode) CephaloTimeNode) CephaloTimeSeries {
	nts := NewCTS()
	cts.TraversalMap(ctn, func(current *CephaloTimeNode) {
		rollwinstart := current.Datetime.Add(-period) //TODO - left, right and center align
		rollwinend := current.Datetime
		inrange, _ := current.findRange(rollwinstart, rollwinend)
		if len(inrange) >= minn {
			calcnode := applied(inrange)
			nts.Insert(calcnode.Datetime, calcnode.Data)
		}
	})
	return nts
}

//RollMean provides the usual rolling mean (moving average) of the time series,
//but based on the period rather than the number of observations
func (cts *CephaloTimeSeries) RollMean(period time.Duration, minn int) CephaloTimeSeries {
	return cts.RollApply(period, cts.Root, minn, func(currents []*CephaloTimeNode) CephaloTimeNode {
		nctt := CephaloTimeNode{Datetime: currents[len(currents)-1].Datetime, Data: 0}
		var valData float64
		for _, cu := range currents {
			valData += cu.Data
		}
		nctt.Data = valData / float64(len(currents))
		return nctt
	})
}

//Node methods considered private (insert, find)
func (ctn *CephaloTimeNode) insert(dattime time.Time, Data float64) error {
	nctt := CephaloTimeNode{Datetime: dattime, Data: Data}
	switch {
	case nctt.Datetime.After(ctn.Datetime):
		//Check right
		if ctn.right == nil {
			ctn.right = &nctt
			return nil
		} else {
			ctn.right.insert(nctt.Datetime, nctt.Data)
		}
	case nctt.Datetime.Before(ctn.Datetime):
		//Check left
		if ctn.left == nil {
			ctn.left = &nctt
			return nil
		} else {
			ctn.left.insert(nctt.Datetime, nctt.Data)
		}
	}
	return nil
}

func (ctn *CephaloTimeNode) find(dattime time.Time) (*CephaloTimeNode, bool) {
	if ctn == nil {
		return &CephaloTimeNode{}, false
	}
	switch {
	case ctn.Datetime.Equal(dattime):
		return ctn, true
	case dattime.After(ctn.Datetime):
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
	if start.Before(ctn.Datetime) {
		findRangeInner(ctn.left, start, end, cb)
	}
	if (start.Before(ctn.Datetime) || start.Equal(ctn.Datetime)) && (end.After(ctn.Datetime) || end.Equal(ctn.Datetime)) {
		cb(ctn)
	}
	if end.After(ctn.Datetime) {
		findRangeInner(ctn.right, start, end, cb)
	}
}

func (ctn *CephaloTimeNode) findMax(parent *CephaloTimeNode) (*CephaloTimeNode, *CephaloTimeNode) {
	if ctn == nil {
		return &CephaloTimeNode{}, parent
	}
	if ctn.right == nil {
		return ctn, parent
	}
	return ctn.right.findMax(ctn)
}

func (ctn *CephaloTimeNode) findMin(parent *CephaloTimeNode) (*CephaloTimeNode, *CephaloTimeNode) {
	if ctn == nil {
		return &CephaloTimeNode{}, parent
	}
	if ctn.left == nil {
		return ctn, parent
	}
	return ctn.left.findMin(ctn)
}

func (ctn *CephaloTimeNode) replaceNode(parent, replacement *CephaloTimeNode) {
	if ctn == nil {
		return
	}
	if ctn == parent.left {
		parent.left = replacement
	}
	parent.right = replacement
}

func (ctn *CephaloTimeNode) delete(dattime time.Time, parent *CephaloTimeNode) error {
	if ctn == nil {
		return errors.New("Can't delete from a nil node")
	}
	switch {
	case dattime.Before(ctn.Datetime):
		return ctn.left.delete(dattime, ctn)
	case dattime.After(ctn.Datetime):
		return ctn.right.delete(dattime, ctn)
	default:
		//If node is leaf node it has no children then remove it from its parent
		if ctn.left == nil && ctn.right == nil {
			ctn.replaceNode(parent, nil)
			return nil
		}
		//If node is half-leaf it has one of the children, so replace node by its child node
		if ctn.left == nil {
			ctn.replaceNode(parent, ctn.right)
			return nil
		}
		if ctn.right == nil {
			ctn.replaceNode(parent, ctn.left)
			return nil
		}
		//If the node is inner then steps are:
		//1. in the left subtree find largest
		leftmax, leftmaxparent := ctn.left.findMax(ctn)
		//2. replace my value and Data with
		ctn.Datetime = leftmax.Datetime
		ctn.Data = leftmax.Data
		//3. remove replacement node
		return leftmax.delete(leftmax.Datetime, leftmaxparent)
	}
}

//AbsDuration returns the absolute value of the supplied time.Duration
func AbsDuration(duration time.Duration) time.Duration {
	if duration < 0 {
		return duration * -1
	}
	return duration
}
