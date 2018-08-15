package cephalolambdas

import (
	"testing"
)

func TestBeakdiff(t *testing.T) {
	t1 := []string{"buda", "budja", "bnumba"}
	t2 := []string{"buda", "bnumba", "bunda"}
	test := Beakdiff(t1, t2)
	if test[0] != "budja" {
		t.Error("Probably not an accurate difference!")
	}
	test2 := Beakdiff(t2, t1)
	if test2[0] != "bunda" {
		t.Error("Probably reverse not working!")
	}
}

func TestBeakfilter(t *testing.T) {
	t1 := []string{"john", "tim", "david", "rogan", "broly"}
	tfun := func(s string) bool { return s == "tim" }
	test := Beakfilter(t1, tfun)
	if test[0] != "tim" {
		t.Error("Probably not filtered properly!")
	}
}

func TestBeakkeys(t *testing.T) {
	tm := map[string]string{"omax": "kum", "lomax": "bum", "demomax": "opol"}
	test := Beakkeys(tm)
	if test[0] != "omax" {
		t.Error("Probably haven't got the right keys!")
	}
}

func TestBeakindex(t *testing.T) {
	tc := []string{"upan", "pupan", "gvadalupa", "membranoz", "kremac"}
	tt := "membranoz"
	test := Beakindex(tc, tt)
	if test != 3 {
		t.Error("Probably not really an accurate index!")
	}
	tx := "dermoman"
	test = Beakindex(tc, tx)
	if test != -1 {
		t.Error("Probably wasn't able to infer non-membership!")
	}
}
