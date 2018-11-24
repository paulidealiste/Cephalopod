package cephalofile

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

//Record represents one record from the iris dataset
type Record struct {
	sl float64
	sw float64
	pl float64
	pw float64
	sp string
	ee error
}

//ImportCSV reads a CSV file and parses it to a DataStore
func ImportCSV(path string) []Record {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 5
	var parsed []Record
	for {
		row, err := reader.Read()
		cph := Record{}
		for i, rec := range row {
			if i == 4 {
				if rec == "" {
					log.Printf("Unexpected type in column %d\n", i)
					cph.ee = fmt.Errorf("Empty string value")
					break
				}
				cph.sp = rec
				continue
			}

			flt, err := strconv.ParseFloat(rec, 64)
			if err != nil {
				log.Printf("Unexpected type in column %d\n", i)
				cph.ee = fmt.Errorf("Unparsable float")
				break
			}
			switch i {
			case 0:
				cph.sl = flt
			case 1:
				cph.sw = flt
			case 2:
				cph.pl = flt
			case 3:
				cph.pw = flt
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}
		if cph.ee == nil {
			parsed = append(parsed, cph)
		}
	}
	return parsed
}
