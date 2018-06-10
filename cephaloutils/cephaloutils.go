// Package cephaloutils provides various utility functions (i.e. min, max, range, ...)
package cephaloutils

import (
	"errors"
	"math"
	"math/rand"
	"time"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephalorandom"
)

// ExtremesRange returns X and Y coordinates range
func ExtremesRange(input *cephalobjects.DataStore) []cephalobjects.DataPoint {
	extremes := make([]cephalobjects.DataPoint, 2)
	var minX, minY float64 = input.Basic[0].X, input.Basic[0].Y
	var maxX, maxY float64 = 0.0, 0.0
	for _, dp := range input.Basic {
		minX = math.Min(minX, dp.X)
		minY = math.Min(minY, dp.Y)
		maxX = math.Max(maxX, dp.X)
		maxY = math.Max(maxY, dp.Y)
	}
	extremes[0].X = minX
	extremes[0].Y = minY
	extremes[1].X = maxX
	extremes[1].Y = maxY
	return extremes
}

// CalculateDescriptors returns means and SDs of a DataPoint slice
func CalculateDescriptors(input []cephalobjects.DataPoint) cephalobjects.Descriptors {
	var meanX, meanY, sdX, sdY float64
	var sumX, sumY, ssX, ssY, sX, sY, sXsY float64 = 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	li := float64(len(input))
	for _, dp := range input {
		sumX += dp.X
		sumY += dp.Y
	}
	meanX = sumX / li
	meanY = sumY / li
	for _, dp := range input {
		sX += dp.X - meanX
		ssX += sX * sX
		sY += dp.Y - meanY
		ssY += sY * sY
		sXsY += sX * sY
	}
	sdX = math.Sqrt(ssX / (li - 1))
	sdY = math.Sqrt(ssY / (li - 1))
	descs := cephalobjects.Descriptors{
		MeanX:   meanX,
		MeanY:   meanY,
		VarX:    ssX / (li - 1),
		VarY:    ssY / (li - 1),
		SdX:     sdX,
		SdY:     sdY,
		CovarXY: sXsY / (li - 1),
	}
	return descs
}

// CovarianceMatrix transforms descriptosrs to a matrix notation
func CovarianceMatrix(desc cephalobjects.Descriptors) cephalobjects.DataMatrix {
	var dmc cephalobjects.DataMatrix
	dmc.Variables = []string{"X", "Y"}
	dmc.Matrix = make([][]float64, len(dmc.Variables))
	dmc.Grep = make(map[string]int)
	for i, name := range dmc.Variables {
		dmc.Matrix[i] = make([]float64, len(dmc.Variables))
		dmc.Grep[name] = i
		for j := range dmc.Variables {
			if i == j {
				if name == "X" {
					dmc.Matrix[i][j] = desc.VarX
				} else {
					dmc.Matrix[i][j] = desc.VarY
				}
			} else {
				dmc.Matrix[i][j] = desc.CovarXY
			}
		}
	}
	return dmc
}

// InverseMatrix calculates the matrix inverse 2x2 matrix only
func InverseMatrix(datma cephalobjects.DataMatrix) cephalobjects.DataMatrix {
	var determinant float64
	diag1, diag2 := 1.0, 1.0
	var imc cephalobjects.DataMatrix
	imc.Variables = datma.Variables
	imc.Grep = datma.Grep
	for i, mp := range datma.Matrix {
		for j := range mp {
			if i == j {
				diag1 *= datma.Matrix[i][j]
			} else {
				diag2 *= datma.Matrix[i][j]
			}
		}
	}
	determinant = diag1 - diag2
	datma.Matrix[0][0], datma.Matrix[1][1] = datma.Matrix[1][1], datma.Matrix[0][0]
	datma.Matrix[0][1] = datma.Matrix[0][1] * -1
	datma.Matrix[1][0] = datma.Matrix[1][0] * -1
	imc.Matrix = datma.Matrix
	for i, mp := range imc.Matrix {
		for j := range mp {
			imc.Matrix[i][j] *= 1 / determinant
		}
	}
	return imc
}

// DotProduct returns the product of matrix and vector mutliplication
func DotProduct(datma cephalobjects.DataMatrix, datve []float64) []float64 {
	var dotprod []float64
	for _, row := range datma.Matrix {
		var rowProd float64
		for i, rowVal := range row {
			rowProd += datve[i] * rowVal
		}
		dotprod = append(dotprod, rowProd)
	}
	return dotprod
}

// TruncatedNormal generates truncated random normals
func TruncatedNormal(desc cephalobjects.Descriptors, l int) []cephalobjects.DataPoint {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	truncgen := make([]cephalobjects.DataPoint, l)
	upperBoundX := desc.MeanX + desc.SdX
	lowerBoundX := desc.MeanX - desc.SdX
	upperBoundY := desc.MeanY + desc.SdY
	lowerBoundY := desc.MeanY - desc.SdY
	for i := range truncgen {
		for {
			truncgen[i].X = math.Abs(random.NormFloat64())*desc.SdX + desc.MeanX
			truncgen[i].Y = math.Abs(random.NormFloat64())*desc.SdY + desc.MeanY
			truncgen[i].G = cephalorandom.RandStringBytes(random, 5)
			truncgen[i].A = truncgen[i].G
			if truncgen[i].X > lowerBoundX && truncgen[i].X < upperBoundX && truncgen[i].Y > lowerBoundY && truncgen[i].Y < upperBoundY {
				break
			}
		}
	}
	return truncgen
}

// EuclideanDistance returns the L2 norm of two DataPoints
func EuclideanDistance(p1 cephalobjects.DataPoint, p2 cephalobjects.DataPoint) float64 {
	ed := math.Sqrt(math.Pow((p1.X-p2.X), 2) + math.Pow((p1.Y-p2.Y), 2))
	return ed
}

// MinSliceIndex returns the position of the slice's smallest element
func MinSliceIndex(input []float64) int {
	m := input[0]
	var mi int
	for i, e := range input {
		if e < m {
			m = e
			mi = i
		}
	}
	return mi
}

// CheckIfSame checks if all the values in the input arrays are equal
func CheckIfSame(s1 []cephalobjects.DataPoint, s2 []cephalobjects.DataPoint) (bool, error) {
	if len(s1) != len(s2) {
		err := errors.New("input slices must be of the same length")
		return false, err
	}
	var counter int
	for i := range s1 {
		if s1[i].X == s2[i].X && s1[i].Y == s2[i].Y {
			counter++
		}
	}
	if counter == len(s1) {
		return true, nil
	}
	return false, nil
}

// PluckStringValues returns a list of string-based values from a list of DataPoint maps
func PluckStringValues(s []cephalobjects.DataPoint, keyCode cephalobjects.GroupType) []string {
	strval := make([]string, 0)
	if keyCode == cephalobjects.Actual {
		for _, dp := range s {
			strval = append(strval, dp.A)
		}
	}
	if keyCode == cephalobjects.Grouped {
		for _, dp := range s {
			strval = append(strval, dp.G)
		}
	}
	return strval
}

// CheckAllTrue checks wheter all members of a boolean list are true
func CheckAllTrue(b []bool) bool {
	allTrue := true
	for _, ob := range b {
		if ob == false {
			allTrue = false
			break
		}
	}
	return allTrue
}

// RandomID returns 8-digit integer IDs
func RandomID() int {
	source := rand.NewSource(time.Now().UnixNano())
	driver := rand.New(source)
	return 10000000 + driver.Intn(99999999-10000000)
}
