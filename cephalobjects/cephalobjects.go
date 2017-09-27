// Package cephalobjects define global data structures
package cephalobjects

// DataPoint is the basic xy analytic data type with a simple annotation
type DataPoint struct {
	X, Y float64
	A, G string
}

// DataStore is a bit complex annotated slice-like data type (other properties to be added)
type DataStore struct {
	Basic []DataPoint
}
