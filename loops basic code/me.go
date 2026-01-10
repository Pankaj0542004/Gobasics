package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	bookings := []string{}

	// print welcome message
	fmt.Printf("Welcome to %v booking application.\nWe have a total of %v tickets and %v are still available.\nGet your tickets here to attend!\n", conferenceName, conferenceTickets, remainingTickets) // Keep only one main function, remove or rename the duplicates

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask for user input

		fmt.Printf("Enter your first name: ")
		fmt.Scanln(&firstName)

		fmt.Printf("Enter your last name: ")
		fmt.Scanln(&lastName)

		fmt.Printf("Enter your email address: ")
		fmt.Scanln(&email)

		fmt.Printf("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		// book ticket in system
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName) //Displays updated system state

		firstNames := []string{}

		// (Blank Identifier): _ is index ignored variable and booking is value variable
		for _, booking := range bookings {
			var name = strings.Fields(booking)       //Fields() splits string by spaces and returns slice of words
			firstNames = append(firstNames, name[0]) //name[0] is first name and appended to firstNames slice
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}

}
