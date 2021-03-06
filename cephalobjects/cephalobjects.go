// Package cephalobjects define global data structures
package cephalobjects

// GroupType - Possible values for grouping keys in a DataPoint struct
type GroupType int

// Possible values for grouping keys in a DataPoint struct => enum-like
const (
	Actual GroupType = iota
	Grouped
)

// DistanceMetric - possible distances for calculating the distance matrix of a DataStore
type DistanceMetric int

// Possible distance metrics
const (
	Euclidean DistanceMetric = iota + 1
	SquaredEuclidean
	Manhattan
	Maximum
	Mahalanobis
)

// LinkageCriteria - Hierarchical clustering linkage criteria
type LinkageCriteria int

// Possible linkage criteria
const (
	Complete LinkageCriteria = iota + 1
	Single
	Average
)

// GrepFold shows row/column position for any entry in the DataMatrix Matrix
type GrepFold struct {
	Row int
	Col int
}

// DataMatrix represents a simple matrix like structure with variable labels on cols and rows
type DataMatrix struct {
	Variables []string
	Matrix    [][]float64
	Grep      map[string]GrepFold
}

// DataMatrixExtreme represents a single extreme value with the info on row and column
// of the extreme value, as well as represntative column/row grep
type DataMatrixExtreme struct {
	Value      float64
	Row        int
	Col        int
	RowName    string
	ColName    string
	Cumulative int
}

// DataPoint is the basic xy analytic data type with a simple annotation (A - actual, G - groupped)
type DataPoint struct {
	UID  string
	X, Y float64
	A, G string
}

// DataStore is a bit complex annotated slice-like data type (other properties to be added)
type DataStore struct {
	Basic    []DataPoint
	Distance DataMatrix
}

// Descriptors represent basic statistics from an array of DataPoints by X and Y coordinates
type Descriptors struct {
	MeanX, MeanY, VarX, VarY, SdX, SdY, CovarXY float64
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

//TimeSeriesDataLike default output of timeseries for json
type TimeSeriesDataLike struct {
	ID   int                   `json:"ID"`
	Data []TimeSeriesDataPoint `json:"series_data"`
}

// TimeSeriesDataPoint default output of timeseries data for json
type TimeSeriesDataPoint struct {
	ID       int     `json:"point_id"`
	Datetime string  `json:"date_time"`
	Data     float64 `json:"data_value"`
}
