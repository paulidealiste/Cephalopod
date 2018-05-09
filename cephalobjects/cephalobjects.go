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
	MeanX, MeanY, VarX, VarY, SdX, SdY float64
}

// ModelSummary holds the usual result structure from a linear regression a (intercept), b (slope) and R squared
type ModelSummary struct {
	A, B, R2 float64
}

// AnovaSummary represents basic analysis of variance table
type AnovaSummary struct {
	SSM, SST, SSE float64
	Dfm, Dft, Dfe float64
	MSM, MST, MSE float64
	F             float64
	P             float64
}
