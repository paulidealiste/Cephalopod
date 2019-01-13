// Package cephalofile provides various I/O
package cephalofile

import (
	"bufio"
	"encoding/json"
	"os"
	"time"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
)

// ExportDataStore takes DataStore and writes its contents to a table
func ExportDataStore(input cephalobjects.DataStore) {
	f, _ := os.Create("../dump.json")
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(encodeDataStore(input))
	w.Flush()
}

func encodeDataStore(input cephalobjects.DataStore) string {
	jo, _ := json.MarshalIndent(input.Basic, "", " ")
	return string(jo)
}

// ExportTimeSeries takes an input TimeSeries and flushes its content
func ExportTimeSeries(input cephalobjects.CephaloTimeSeries) {
	f, _ := os.Create("../tsdump.json")
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(encodeTimeSeries(input))
	w.Flush()
}

func encodeTimeSeries(input cephalobjects.CephaloTimeSeries) string {
	var fullma []cephalobjects.TimeSeriesDataPoint
	pointcount := 0
	input.TraversalMap(input.Root, func(current *cephalobjects.CephaloTimeNode) {
		tsdp := cephalobjects.TimeSeriesDataPoint{ID: pointcount, Datetime: current.Datetime.Format(time.RFC3339), Data: current.Data}
		fullma = append(fullma, tsdp)
		pointcount++
	})

	ts, _ := json.MarshalIndent(fullma, "", " ")
	return string(ts)
}
