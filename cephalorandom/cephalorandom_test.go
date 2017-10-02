package cephalorandom

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// whether random generator outputs non-empty DataStore struct with a predefined length
func TestGeneratesRadnomDataStore(t *testing.T) {
	testlen := 12
	test, _ := GenerateRandomDataStore(testlen, 4, 0.3)
	if reflect.TypeOf(test).Name() != "DataStore" {
		t.Error("Generated data is not a DataStore object")
	}
	if len(test.Basic) <= 0 {
		t.Error("Generated data must not be of zero length")
	}
	if len(test.Basic) != testlen {
		t.Error("Generated data did not conform to the projected length")
	}
}

// wheter random generator propagates ok errors
func TestGeneratesRadnomErrors(t *testing.T) {
	_, err := GenerateRandomDataStore(0, 0, 2)
	if err.Error() != "check input parameters" {
		t.Error("Errors not propagated correctly")
	}
}

// are generated groups of a desired length
func TestGeneratedGroupsLength(t *testing.T) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	testlen := 566
	test := generateGroups(r, testlen, 4)
	var sum int
	for _, gr := range test {
		sum += gr.length
	}
	if sum != testlen {
		t.Errorf("These lengths must match but one is %d ant the other %d", sum, testlen)
	}
}

// generates string of a predifined length
func TestGenerateRandomString(t *testing.T) {
	l := 10
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	test := RandStringBytes(r, l)
	if len(test) != l {
		t.Error("Generated string length doesn't match the required length")
	}
}
