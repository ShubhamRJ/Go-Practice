package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var bookings []string

	for {
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

		var isValidName = len(userFirstName) >= 2 && len(userLastName) >= 2
		var isValidEmail = strings.Contains(userEmail, "@")
		var isValidTickets = userNumberOfTicket > 0 && userNumberOfTicket <= remainingTickets

		if !isValidTickets {
			fmt.Printf("We only have %v tickets, so you cannot book %v tickets", remainingTickets, userNumberOfTicket)
			continue
		}
		if !isValidName || !isValidEmail {
			fmt.Println("Please input correct name/email address")
		}
		fmt.Println("Welcome", userFirstName, userLastName)
		fmt.Printf("You have booked %d tickets. \n", userNumberOfTicket)
		bookings = append(bookings, userFirstName+" "+userLastName)
		remainingTickets -= userNumberOfTicket
		fmt.Println("Bookings: ", bookings)
		fmt.Printf("We have total of %v tickets and %v are left.\n", conferenceTickets, remainingTickets)

		var firstnames []string
		for _, booking := range bookings {
			var name = strings.Fields(booking)
			firstnames = append(firstnames, name[0])
		}
		fmt.Println(firstnames)
		if remainingTickets == 0 {
			fmt.Println("All tickets are sold now!")
			break
		}
	}
}
