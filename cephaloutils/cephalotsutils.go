// Package cephaloutils provides various utility functions (i.e. min, max, range and timeseries utils ...)
package cephaloutils

import (
	"fmt"
	"time"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
)

//CTSListForm returns the slice-like TimeSeriesDataLike cepahlotimeseries representation
func CTSListForm(cts cephalobjects.CephaloTimeSeries) cephalobjects.TimeSeriesDataLike {
	var tslist []cephalobjects.TimeSeriesDataPoint
	pointcount := 0
	cts.TraversalMap(cts.Root, func(current *cephalobjects.CephaloTimeNode) {
		tspoint := cephalobjects.TimeSeriesDataPoint{ID: pointcount, Data: current.Data, Datetime: current.Datetime.Format(time.RFC3339)}
		tslist = append(tslist, tspoint)
		pointcount++
	})
	tsdl := cephalobjects.TimeSeriesDataLike{ID: cts.ID, Data: tslist}
	return tsdl
}

//TSListMap provides the transformation of a cephalotimeseries to an ID indexed TimeSeriesDataLike map
func CTSListMap(cts cephalobjects.CephaloTimeSeries) map[int]cephalobjects.TimeSeriesDataLike {
	tsdl := CTSListForm(cts)
	tsdlm := make(map[int]cephalobjects.TimeSeriesDataLike)
	tsdlm[tsdl.ID] = tsdl
	return tsdlm
}

//TSListsFromTSTrees enables the multi cephalotimeseries conversion to multi-id-keyed map of TimeSeriesDataLikes
func TSListsFromTSTrees(ctss []cephalobjects.CephaloTimeSeries) map[int]cephalobjects.TimeSeriesDataLike {
	tsdlsm := make(map[int]cephalobjects.TimeSeriesDataLike)
	jobs := make(chan cephalobjects.CephaloTimeSeries, 100)
	results := make(chan cephalobjects.TimeSeriesDataLike, 100)
	for w := 1; w <= 3; w++ {
		go tslistworker(w, jobs, results)
	}

	return tsdlsm
}

func tslistworker(id int, jobs <-chan cephalobjects.CephaloTimeSeries, results chan<- cephalobjects.TimeSeriesDataLike) {
	for cts := range jobs {
		fmt.Println("worker", id, "started  job", cts.ID)
		tsdl := CTSListForm(cts)
		fmt.Println("worker", id, "finished job", cts.ID)
		results <- tsdl
	}
}
