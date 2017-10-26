// Package cephalofile provides various I/O
package cephalofile

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
)

// ExportDataStore takes DataStore and writes its contents to a table in txt file (no batteries error checks)
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
