package calculations

import "fmt"

func subInputs() int {

	var numberOne, numberTwo int

	fmt.Print("Please enter number one: ")
	fmt.Scan(&numberOne)

	fmt.Print("Please enter number two: ")
	fmt.Scan(&numberTwo)

	sum := numberOne - numberTwo
	return sum
}
