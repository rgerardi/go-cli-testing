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

	n1, err := askNumber(in)
	if err != nil {
		return 1
	}

	n2, err := askNumber(in)
	if err != nil {
		return 1
	}

	sum := add(n1, n2)
	printSum(out, sum)

	return 0
}

func add(n1, n2 float64) float64 {
	return n1 + n2
}

func askNumber(in io.Reader) (float64, error) {
	fmt.Println("Type a number:")
	n := 0.0
	_, err := fmt.Fscanf(in, "%f\n", &n)

	if err != nil {
		return 0.0, err
	}

	return n, nil
}

func printSum(out io.Writer, sum float64) {
	fmt.Fprintln(out, "The sum is:", sum)
}
