package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_Add(t *testing.T) {
	sum := add(3.2, 6.8)

	if sum != 10.0 {
		t.Errorf("Sum should be 10. Got %f\n", sum)
	}
}

func Test_askNumber(t *testing.T) {
	in := bytes.NewBufferString("5") // HL

	var exp float64 = 5

	res := askNumber(in)

	if exp != res {
		t.Errorf("Expected %f, got %f\n", exp, res)
	}
}

func Test_PrintSum(t *testing.T) {
	out := bytes.NewBuffer([]byte{}) // HL

	var sum float64 = 5
	exp := fmt.Sprintln("The sum is:", sum)
	printSum(out, sum)

	if exp != out.String() {
		t.Errorf("Expected '%s', got '%s'\n", exp, out.String())
	}
}
