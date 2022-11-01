package controllers

import (
	"fmt"
	"log"
)

var UserPin string = "1234"

func validPin(pin string) (string, error) {
	if len(pin) != 4 {
		return "", fmt.Errorf("\nInvalid Pin:Four digits required")

	}
	return pin, nil
}

func VerifyUserPin(pin string) bool {
	pin, err := validPin(pin)
	if err != nil {
		log.Println(err)
		return false
	}
	if pin != UserPin {
		return false
	}

	return true
}

func ChangeUserPin() {
	var oldPin string
	var newPin string

	fmt.Println("Enter your current or default PIN:")
	fmt.Scan(&oldPin)
	oldPin, err := validPin(oldPin)
	if err != nil {
		log.Println("Enter a four digit pin")

	}

	fmt.Println("Enter your new PIN:")
	fmt.Scan(&newPin)
	if oldPin != UserPin {
		log.Println("PIN is incorrect")
	}

	newPin, err = validPin(newPin)
	if err != nil {
		log.Println("Enter a four digit pin")
	}
	UserPin = newPin

	log.Println("User PIN changed")

}

func Login() bool{
	fmt.Println("Default PIN is 1234")
	for {
		var pin string 
		fmt.Println("Enter your pin")
		fmt.Scan(&pin)
		if VerifyUserPin(pin){
			return true
		}
		break
	}
	return false
}
