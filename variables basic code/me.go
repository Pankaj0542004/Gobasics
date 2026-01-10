package main

import "fmt"

func main() {
	const conferenceTickets int = 10
	var remainingTickets uint = 5
	consferenceName := "Go conference"

	fmt.Printf("Welcome to %v booking application. \nWe have total of %v tickets and %v are still available. \nGet your ticket here to attend \n",
		conferenceTickets, remainingTickets, consferenceName)

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v \n", remainingTickets, consferenceName)

}
