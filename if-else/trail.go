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

	fmt.Printf("Welcome to %v booking application.\nWe have a total of %v tickets and %v are still available.\nGet your tickets here to attend!\n", conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//ask to user to input
		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)
		
		fmt.Printf("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Printf("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		//validate user input
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@gmail.com")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			//print only first names
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Split(booking, " ")
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first name %v\n", firstNames)

			//exit application when tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The name you entered is too short.")
			}

			if !isValidEmail {
				fmt.Println("The email address you entered is invalid.")
			}

			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid.")
			}
			continue
		}
	}
}
