package main

import "fmt"

func main(){
	const conferenceTickets int = 80
	var remaningTickets uint = 80
	conferenceName := "flight ticket booking" //string type
	bookings := []string{}

	// print welcome message
	fmt.Printf("Welcome to %v booking application.\n We have total of %v tickets and %v are still available.\n Get your tickets here to attend\n" , conferenceName, conferenceTickets, remaningTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask for user input
	fmt.Printf("Enter your first Name: ")
	fmt.Scanln(&firstName)

	fmt.Printf("Enter Your Last Name: ")
	fmt.Scanln(&LastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	remaningTickets = remaningTickets - userTickets
	bookings = append(bookings, firstName+ " " +lastName)

	fmt.Printf("Thank You %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, userTickets, email)
}