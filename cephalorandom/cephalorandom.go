// Package cephalorandom provides utilities for random data generation
package cephalorandom

import (
	"errors"
	"math"
	"math/rand"
	"time"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type randomGroup struct {
	mean, length int
}

// GenerateRandomDataStore generates random DataStore for mock input data (rho tba)
func GenerateRandomDataStore(length int, groups int, rho float64) (cephalobjects.DataStore, error) {
	var randstore cephalobjects.DataStore
	if length <= 0 || groups <= 0 || rho < -1 || rho > 1 {
		return randstore, errors.New("check input parameters")
	}
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	randstore.Basic = randomDPSlice(random, length, groups, rho)
	return randstore, nil
}

func randomDPSlice(r *rand.Rand, l int, g int, rho float64) []cephalobjects.DataPoint {
	randbasic := make([]cephalobjects.DataPoint, l)
	groupseeds := generateGroups(r, l, g)
	imer := 0
	for _, group := range groupseeds {
		temp := RandStringBytes(r, 5)
		for j := 0; j < group.length; j++ {
			iter := imer + j
			randbasic[iter].X = math.Abs(r.NormFloat64()) + float64(group.mean)
			randbasic[iter].Y = math.Abs(r.NormFloat64()) + float64(group.mean)
			randbasic[iter].A = temp
			randbasic[iter].UID = RandStringBytes(r, 4)
		}
		imer += group.length
	}
	return randbasic
}

func generateGroups(r *rand.Rand, l int, g int) []randomGroup {
	groupseeds := make([]randomGroup, g)
	lm := l
	for i := range groupseeds {
		groupseeds[i].mean = r.Intn(3)
		groupseeds[i].length = int(l / g)
		lm -= groupseeds[i].length
		if i == g-2 {
			groupseeds[i].length = lm
		}
	}
	return groupseeds
}

// RandStringBytes returns random letter string - used for group labels and IDs
func RandStringBytes(r *rand.Rand, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}

// RandomID returns 8-digit integer IDs
func RandomID() int {
	source := rand.NewSource(time.Now().UnixNano())
	driver := rand.New(source)
	return 10000000 + driver.Intn(99999999-10000000)
}

// func generateCorrelated(a float64, b float64, rho float64) float64 {
// 	return rho*b + math.Sqrt(1-math.Pow(rho, 2.0))*a
// }
