package calculations

import "fmt"

func RunTheProgram() {

	// choice of users
	var userChoice int

	// welcome messages
	promptMessages()

	// getting user choice
	fmt.Print("Please select one of operations: 1)sum - 2)subtraction - 3)multiplication - 4)division => ")
	fmt.Scan(&userChoice)

	switch userChoice {

	case 1:
		finalAddNumber := addInputs()
		fmt.Printf("The sum is: %v", finalAddNumber)

	case 2:
		finalSubNumber := subInputs()
		fmt.Printf("The substraction is: %v", finalSubNumber)

	case 3:
		finalMulNumber := mulInputs()
		fmt.Printf("The multiplication is: %v", finalMulNumber)

	case 4:
		finalDivNumber := divInputs()
		fmt.Printf("The division is: %v", finalDivNumber)

	default:
		fmt.Println("You have eo enter a number between 1 to 4.")
	}

}

func promptMessages() {

	fmt.Println("Hello! welcome to calculator program. hope you enjoy it!")
	fmt.Println("----------------------------------------------------------")
}
