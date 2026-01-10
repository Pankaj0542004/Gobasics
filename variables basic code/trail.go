package main //main is package to execute the main function

import "fmt" //fmt stands for format and provides I/O functions

func main() { //main function is the entry point of the program
	const conferenceTickets int = 50  //conferenceTickets is named constant
	var remainingTickets uint = 50    // var is varible declaration as remainingTickets
	conferenceName := "Go conference"

	fmt.Printf(" Welcome to %v booking application. \n We have total of %v tickets and %v are still available. \nGet your ticket here to attend  \n", 
	conferenceTickets, remainingTickets, conferenceName)
	        // here %v is a placeholder for variable value of any type (conferenceTickets, remainingTickets, conferenceName)
	

	var firstName string
	var lastName string
	var email string 
	var userTickets uint

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTickets)


	// book tickets in system
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n",
	firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v \n", remainingTickets, conferenceName)
}
