// Package cephalofile provides various I/O
package cephalofile

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

// ExportDataStore takes DataStore and writes its contents to a table
func ExportDataStore(input cephalobjects.DataStore, path string) {
	f, _ := os.Create(path)
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(encodeDataStore(input))
	w.Flush()
}

func encodeDataStore(input cephalobjects.DataStore) string {
	jo, _ := json.MarshalIndent(input.Basic, "", " ")
	return string(jo)
}

// ExportAnnonymusCSV writes 2D integer array data to a file of choice
func ExportAnnonymusCSV(input [][]int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, row := range input {
		if err := writer.Write(cephaloutils.IntsliceToString(row)); err != nil {
			return err
		}
	}
	return nil
}

// ExportTimeSeries takes an input cephalotimeseries and flushes its content
func ExportTimeSeries(input cephalobjects.CephaloTimeSeries, path string) {
	f, _ := os.Create(path)
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(encodeTimeSeries(input))
	w.Flush()
}

func encodeTimeSeries(input cephalobjects.CephaloTimeSeries) string {
	fullma := cephaloutils.CTSListMap(input)
	ts, _ := json.MarshalIndent(fullma, "", " ")
	return string(ts)
}

// ExportTimeSeriesList is a list-based version of cephalotimeseries to json export
func ExportTimeSeriesList(input []cephalobjects.CephaloTimeSeries, path string) {
	f, _ := os.Create(path)
	defer f.Close()

	w := bufio.NewWriter(f)
	timeserieslistmaps, _ := json.MarshalIndent(cephaloutils.TSListsFromTSTrees(input), "", " ")
	w.WriteString(string(timeserieslistmaps))
	w.Flush()
}
