package main

import "fmt"

func main () {
	// if else statement in go
	age := 18

	// 1st method
	if age >= 28{
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

	// 2nd method
	if age >= 28{
		fmt.Println("You are eligible to vote.")
	} else if age >= 18 {
		fmt.Println("You are eligible to vote but not eligible for driving license.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

	// 3rd method
	if age >= 28{
		fmt.Println("You are eligible to vote.")
	} else if age >= 18 && age < 28 {
		fmt.Println("You are eligible to vote but not eligible for driving license.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

	// 4th method
	if age >= 28{
		fmt.Println("You are eligible to vote.")
	} else if age >= 18 && age < 28 {
		fmt.Println("You are eligible to vote but not eligible for driving license.")
	} else if age >= 16 && age < 18 {
		fmt.Println("You are eligible for driving license but not eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

	// 5th method
	var role = "admin"
	var permission = true

	if role == "admin" && permission {
		fmt.Println("You have admin access.")
	} else {
		fmt.Println("You do not have admin access.")
	}
}