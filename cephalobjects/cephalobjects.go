// Package cephalobjects define global data structures
package cephalobjects

// GroupType => Possible values for grouping keys in a DataPoint struct
type GroupType int

// Possible values for grouping keys in a DataPoint struct => enum-like
const (
	Actual GroupType = iota
	Grouped
)

// DataPoint is the basic xy analytic data type with a simple annotation (A - actual, G - groupped)
type DataPoint struct {
	X, Y float64
	A, G string
}

// DataStore is a bit complex annotated slice-like data type (other properties to be added)
type DataStore struct {
	Basic []DataPoint
}

// Descriptors represent basic statistics from an array of DataPoints by X and Y coordinates
type Descriptors struct {
	MeanX, MeanY, SdX, SdY float64
}
