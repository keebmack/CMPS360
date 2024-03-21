package main

import "fmt"

func main() {
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Bison Up"
	bookings := []string{}

	fmt.Printf("Welcome to %s booking application. \nWe have a total of %d tickets still available. \nGet your tickets here to attend\n", conferenceName, remainingTickets)

	var (
		firstName   string
		lastName    string
		email       string
		userTickets uint
	)

	fmt.Println("Enter your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of Tickets: ")
	fmt.Scanln(&userTickets)

	remainingTickets -= userTickets
	bookings = append(bookings, fmt.Sprintf("%s %s", firstName, lastName))

	fmt.Printf("Thanks %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)

	fmt.Printf("%d tickets remaining for %s\n", remainingTickets, conferenceName)
	fmt.Printf("These are all of our bookings: %v\n", bookings)
}
