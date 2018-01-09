package main

import "fmt"

func main() {

	fmt.Println("Type the first number:")
	n1 := 0.0
	fmt.Scanf("%f\n", &n1)

	fmt.Println("Type the second number:")
	n2 := 0.0
	fmt.Scanf("%f\n", &n2)

	fmt.Println("The sum is:", n1+n2)
}
