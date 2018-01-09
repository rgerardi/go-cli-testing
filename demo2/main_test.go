package main

import (
	"testing"
)

func Test_Add(t *testing.T) {
	sum := add(3.2, 6.8)

	if sum != 10.0 {
		t.Errorf("Sum should be 10. Got %f.", sum)
	}
}
