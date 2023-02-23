package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("ConferenceName is %T , ConferenceTickets is %T , RemainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their name and tickets

	fmt.Println("Enter Your First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your Last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scan(&email)

	fmt.Println("Enter Your Number of Tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank You %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
