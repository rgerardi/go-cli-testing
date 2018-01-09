package main

import (
	"fmt"
	"io"
	"os"
)

// START OMIT
func main() {
	n1 := askNumber(os.Stdin) // HL
	n2 := askNumber(os.Stdin)
	sum := add(n1, n2)
	printSum(os.Stdout, sum)
}

func askNumber(in io.Reader) float64 { // HL
	fmt.Println("Type a number:")
	n := 0.0
	fmt.Fscanf(in, "%f\n", &n)

	return n
}

// END OMIT

func add(n1, n2 float64) float64 {
	return n1 + n2
}

func printSum(out io.Writer, sum float64) { // HL
	fmt.Fprintln(out, "The sum is:", sum)
}
