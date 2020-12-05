package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {

	fac := map[int64]string{

		10:     "3628800",
		20:     "2432902008176640000",
		10000:  "2846259680917054518906413",
		100000: "282422940796034787429342157802453551",
	}

	for num, res := range fac {

		str := doFactorial(num)

		if str[:5] != res[:5] {
			t.Errorf("the first five number of %d! should be %s, not %s", num, res[:5], str[:5])
			continue
		}
		t.Logf("%d! had been calculated correctly", num)
	}
}

func pow(a, n int) int64 {
	ret := int64(a)
	for i := 0; i < n; i++ {
		ret *= ret
	}
	return ret
}
