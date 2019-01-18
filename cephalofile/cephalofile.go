// Package cephalofile provides various I/O
package cephalofile

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
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
	fullma := cephaloutils.TSListForm(input)
	ts, _ := json.MarshalIndent(fullma, "", " ")
	return string(ts)
}
