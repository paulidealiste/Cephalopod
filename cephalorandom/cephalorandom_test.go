package cephalorandom

import (
	"reflect"
	"testing"
)

// whether random generator outputs non-empty DataStore struct with a predefined length
func TestGeneratesRadnomDataStore(t *testing.T) {
	testlen := 5
	test, _ := GenerateRandomDataStore(testlen, 3)
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
}
