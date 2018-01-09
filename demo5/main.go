package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	os.Exit(run(os.Stdin, os.Stdout))

}

func run(in io.Reader, out io.Writer) int {

	n1 := askNumber(in)
	n2 := askNumber(in)
	sum := add(n1, n2)
	printSum(out, sum)

	return 0
}

func add(n1, n2 float64) float64 {
	return n1 + n2
}

func askNumber(in io.Reader) float64 {
	fmt.Println("Type a number:")
	n := 0.0
	fmt.Fscanf(in, "%f\n", &n)

	return n
}

func printSum(out io.Writer, sum float64) {
	fmt.Fprintln(out, "The sum is:", sum)
}
