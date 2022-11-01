package controllers

import (
	"fmt"
)

var balance float64 = 0.0

func CheckBalnce() {
	fmt.Println("Your current balance is :", balance)

}
func Withdraw() {
	var Amount_to_Withdraw float64
	fmt.Println("How much do you want withdraw :")
	fmt.Scan(&Amount_to_Withdraw)
	balance = balance - Amount_to_Withdraw
	fmt.Printf("Success!, you have withdrawn %f from your account", Amount_to_Withdraw)
	fmt.Printf("New balance is %f :", balance)
	fmt.Println()

}
func Deposit (){
	var Amount_to_Deposit float64
	fmt.Println("How much do you want deposit :")
	fmt.Scan(&Amount_to_Deposit)
	balance = balance + Amount_to_Deposit
	fmt.Printf("Success!, you have deposited %f from your account", Amount_to_Deposit)
	fmt.Printf("New balance is %f :", balance)
	fmt.Println()

}