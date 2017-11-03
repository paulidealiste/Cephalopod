package cephalolinreg

import (
	"math"
	"testing"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephalorandom"
)

// Does hypothesis work
func TestHypothesis(t *testing.T) {
	a := 2.0
	b := 4.0
	x := 2.0
	test := hypothesis(a, b, x)
	if test != 10 {
		t.Error("Hypothesis function not working")
	}
}

// Does least squares return full and accurrate model parameters and r2
func TestLeastSquares(t *testing.T) {
	input, _ := cephalorandom.GenerateRandomDataStore(120, 3, 0.5)
	test := leastSquares(input.Basic)
	if test.A == 0.0 && test.B == 0 && test.R2 == 0.0 || math.IsNaN(test.A) || math.IsNaN(test.B) || math.IsNaN(test.R2) {
		t.Error("Least squares did not return any meaningful results")
	}
}

// Wheter cost function calculates non-zero costs and returns re-calculated coeffs in a leanrning step as well as coeffs after descent loop
func TestCostLearnDescent(t *testing.T) {
	input, _ := cephalorandom.GenerateRandomDataStore(120, 3, 0.5)
	var testPar cephalobjects.ModelSummary
	testPar.A = 0.65
	testPar.B = 1.34
	test := cost(input.Basic, testPar)
	test2 := learn(input.Basic, testPar)
	test3 := gradientDescent(input.Basic)
	if test == 0.0 || math.IsNaN(test) {
		t.Error("Cost function does not work properly")
	}
	if test2.A == testPar.A && test2.B == testPar.B {
		t.Error("Learning step not performed properly")
	}
	if test3.A == 0.0 && test3.B == 0.0 {
		t.Error("Gradient descent not performed properly")
	}
}
