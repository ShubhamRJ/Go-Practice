package helper

import "strings"

func IsValidInput(first string, last string, email string, tickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(first) >= 2 && len(last) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTickets = tickets > 0 && tickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}
