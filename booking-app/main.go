package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var bookings []string

	var userFirstName string
	var userLastName string
	var userEmail string
	var userNumberOfTicket uint
	fmt.Print("First Name: ")
	fmt.Scan(&userFirstName)
	fmt.Print("Last Name: ")
	fmt.Scan(&userLastName)
	fmt.Print("Email: ")
	fmt.Scan(&userEmail)
	fmt.Print("Number of tickets: ")
	fmt.Scan(&userNumberOfTicket)
	fmt.Println("Welcome", userFirstName, userLastName)
	fmt.Printf("You have booked %d tickets. \n", userNumberOfTicket)
	bookings = append(bookings, userFirstName)
	remainingTickets -= userNumberOfTicket
	fmt.Println("Bookings: ", bookings)
	fmt.Printf("We have total of %v tickets and %v are left.\n", conferenceTickets, remainingTickets)
}
