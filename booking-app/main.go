package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getUserInput() (string, string, string, uint) {
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
	return userFirstName, userLastName, userEmail, userNumberOfTicket
}

func bookTickets(first string, last string, tickets uint) {
	fmt.Println("Welcome", first, last)
	fmt.Printf("You have booked %d tickets. \n", tickets)
	bookings = append(bookings, first+" "+last)
	remainingTickets -= tickets
	fmt.Println("Bookings: ", bookings)
	fmt.Printf("We have total of %v tickets and %v are left.\n", conferenceTickets, remainingTickets)
}

func describeBookings() {
	var firstnames []string
	for _, booking := range bookings {
		var name = strings.Fields(booking)
		firstnames = append(firstnames, name[0])
	}
	fmt.Println(firstnames)
}

func main() {
	greetUsers()
	for {
		var userFirstName, userLastName, userEmail, userNumberOfTicket = getUserInput()
		var isValidName, isValidEmail, isValidTickets = helper.IsValidInput(userFirstName, userLastName, userEmail, userNumberOfTicket, remainingTickets)

		if !isValidTickets {
			fmt.Printf("We only have %v tickets, so you cannot book %v tickets\n", remainingTickets, userNumberOfTicket)
			continue
		}
		if !isValidName || !isValidEmail {
			fmt.Println("Please input correct name/email address")
			continue
		}
		bookTickets(userFirstName, userLastName, userNumberOfTicket)
		describeBookings()
		if remainingTickets == 0 {
			fmt.Println("All tickets are sold now!")
			break
		}
	}
}
