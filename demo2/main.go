package main

import "fmt"

// START OMIT
func main() {

	fmt.Println("Type the first number:")
	n1 := 0.0
	fmt.Scanf("%f\n", &n1)

	fmt.Println("Type the second number:")
	n2 := 0.0
	fmt.Scanf("%f\n", &n2)

	fmt.Println("The sum is:", add(n1, n2))
}

func add(n1, n2 float64) float64 { // HL
	return n1 + n2 // HL
} // HL
// END OMIT
