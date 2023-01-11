package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

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

func bookTickets(first string, last string, email string, tickets uint) {
	fmt.Println("Welcome", first, last)
	fmt.Printf("You have booked %d tickets. \n", tickets)
	var userData = make(map[string]string)
	userData["firstName"] = first
	userData["lastName"] = last
	userData["email"] = email
	userData["tickets"] = strconv.FormatUint(uint64(tickets), 10)
	bookings = append(bookings, userData)
	remainingTickets -= tickets
	fmt.Println("Bookings: ", bookings)
	fmt.Printf("We have total of %v tickets and %v are left.\n", conferenceTickets, remainingTickets)
}

func describeBookings() {
	var firstnames []string
	for _, booking := range bookings {
		var name = booking["firstName"]
		firstnames = append(firstnames, name)
	}
	fmt.Println(firstnames)
}

var wg = sync.WaitGroup{}

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
		bookTickets(userFirstName, userLastName, userEmail, userNumberOfTicket)
		wg.Add(1)
		go sendTicket(userNumberOfTicket, userFirstName, userLastName, userEmail)
		describeBookings()
		if remainingTickets == 0 {
			fmt.Println("All tickets are sold now!")
			break
		}
	}
	wg.Wait()
}

func sendTicket(tickets uint, firstname string, lastname string, email string) {
	time.Sleep(5 * time.Second)
	fmt.Println("############ Sending Email ############")
	fmt.Printf("%v tickets for %v %v to %v\n", tickets, firstname, lastname, email)
	fmt.Println("#######################################")
	wg.Done()
}
