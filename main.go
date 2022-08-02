package main

import "fmt"

func main() {
	conferenceName := "Beedoz Conference"
	var ticketsAmount int = 50
	var remainingTickets uint = 50

	fmt.Printf("\nWelcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still remaining!\n", ticketsAmount, remainingTickets)
	fmt.Println("Get your tickets here!")

	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint
	var bookings []string

	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&userEmail)
	fmt.Print("How many tickets do you want to book? ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v!\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("There are %v tickets remaining for %v!\n", remainingTickets, conferenceName)
	fmt.Printf("All bookings: %v\n", bookings)
}
