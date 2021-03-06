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
	in := bytes.NewBufferString("5")

	var exp float64 = 5

	res, err := askNumber(in)

	if err != nil {
		t.Errorf("Unexpected error %s\n", err.Error())
	}

	if exp != res {
		t.Errorf("Expected %f, got %f\n", exp, res)
	}
}

func Test_PrintSum(t *testing.T) {
	out := bytes.NewBuffer([]byte{})

	var sum float64 = 5
	exp := fmt.Sprintln("The sum is:", sum)
	printSum(out, sum)

	if exp != out.String() {
		t.Errorf("Expected '%s', got '%s'\n", exp, out.String())
	}
}

func Test_Run(t *testing.T) {
	tests := []struct { // HL
		name       string  // HL
		n1         string  // HL
		n2         string  // HL
		expRetCode int     // HL
		expSum     float64 // HL
	}{ // HL
		{"2 ints", "5", "5", 0, 10},        // HL
		{"2 floats", "3.2", "2.4", 0, 5.6}, // HL
		{"blank", "", "", 1, 0},            // HL
		{"first text", "r", "4", 1, 0},     // HL
		{"second text", "4", "r", 1, 0},    // HL
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) { // HL
			in := bytes.NewBufferString(fmt.Sprintf("%s\n%s\n", test.n1, test.n2))
			out := bytes.NewBuffer([]byte{})
			exp := fmt.Sprintln("The sum is:", test.expSum)
			if test.expRetCode != 0 {
				exp = ""
			}
			retCode := run(in, out)

			t.Log("Message:", out.String())
			if test.expRetCode != retCode {
				t.Errorf("Expected %d, got %d\n", test.expRetCode, retCode)
			}
			if exp != out.String() {
				t.Errorf("Expected '%s', got '%s'\n", exp, out.String())
			}
		})
	}
}
