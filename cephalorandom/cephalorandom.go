// Package cephalorandom provides utilities for random data generation
package cephalorandom

import (
	"errors"
	"math/rand"
	"time"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
)

// GenerateRandomDataStore generates random DataStore for mock input data
func GenerateRandomDataStore(length int, groups int) (cephalobjects.DataStore, error) {
	var randstore cephalobjects.DataStore
	if length <= 0 || groups <= 0 {
		return randstore, errors.New("eh")
	}
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	randstore.Basic = randomDPSlice(random, length)
	return randstore, nil
}

func randomDPSlice(r *rand.Rand, l int) []cephalobjects.DataPoint {
	randbasic := make([]cephalobjects.DataPoint, l)
	for _, dp := range randbasic {
		dp.X = r.NormFloat64()
		dp.Y = r.NormFloat64()
	}
	return randbasic
}
