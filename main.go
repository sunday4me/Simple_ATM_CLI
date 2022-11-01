package main

import (
	"atm/controllers"
	"fmt"
	"log"
	"os"
	"strings"
)

func welcome() {
	welcMess := "*********{Welcome to the ALT_SCHOOL ATM************"
	fmt.Println(welcMess)
}

func exitProgram() {
	fmt.Println("Thanks for banking with Alt_School")
	os.Exit(0)
}

func performAnotherOperation() {
	var newOption string
	fmt.Println("Do want perform another transaction?")
	fmt.Scan(&newOption)
	if strings.ToLower(newOption) == "yes" {
		menu()

	}
	exitProgram()

}

func menu() {
	fmt.Println("Select Operation")
	fmt.Println("1. Change Pin \n 2. Check Account Balance \n 3.Withdrawal \n 4.Deposit \n 5. Exit program")

	var SelectedOp int

	_, err := fmt.Scan(&SelectedOp)

	if err != nil {
		log.Println("Please the correct menu item")

	}
	switch SelectedOp {
	case 1:
		controllers.ChangeUserPin()
		performAnotherOperation()
	case 2:
		controllers.CheckBalnce()
		performAnotherOperation()

	case 3:
		controllers.Withdraw()
		performAnotherOperation()
	case 4:
		controllers.Deposit()
		performAnotherOperation()
	case 5:
		exitProgram()
	default:
		fmt.Println("Invalid Input")
		performAnotherOperation()
	}

}

func main() {
	welcome()

	if controllers.Login() {
		menu()

	}
	exitProgram()
}
