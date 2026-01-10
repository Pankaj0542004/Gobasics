package main

import "fmt"

func main(){
	const conferenceTickets int =  50
	var remaningTickets uint = 50
	conferenceName := "Go conference" //string type

	//create [] to store multiple values and initialize it to empty in {} empty slice and booking will store multiple bookings names 
	bookings := []string{}

	//print welcome message
	fmt.Printf("Welcome to %v booking application.\n We have total of %v tickets and %v are still available.\n Get your tickets here to attend\n" , 
	conferenceName, conferenceTickets, remaningTickets)
	
	var firstName string 
	var lastName string
	var email string
	var userTickets uint

	//ask to user input
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remaningTickets = remaningTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName) // append does not modify it returns a new slice with the added value We store in back in bookings slice

	fmt.Printf("Thank You %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, userTickets, email)
	fmt.Printf("%v tickets are remaning for %v\n", remaningTickets, conferenceName)
	
	fmt.Printf("These are all our bookings: %v\n", bookings)

}