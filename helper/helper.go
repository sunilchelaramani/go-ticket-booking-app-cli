package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, emailAddress string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := len(emailAddress) > 3 && strings.Contains(emailAddress, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
