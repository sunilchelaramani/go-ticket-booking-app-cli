package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

//var bookings [50]string <= array
//using slice to store data
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	emailAddress    string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//	for {

	firstName, lastName, emailAddress, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(firstName, lastName, userTickets, emailAddress)
		wg.Add(1)
		go sendTicket(firstName, lastName, userTickets, emailAddress)
		firstNames := getFirstNames()

		if remainingTickets == 0 {
			fmt.Printf("All tickets sold out\n")
			//break
			// end program
		}

		fmt.Printf("firstname of all members: %s\n", firstNames)
	} else {
		if !isValidName {
			fmt.Println("You entered invalid data in name, please try again")
		}
		if !isValidEmail {
			fmt.Println("You email address is invalid, please try again")
		}
		if !isValidTicketNumber {
			fmt.Printf("We only have %v tickets left, You cannot book %v tickets\n", remainingTickets, userTickets)
		}
	}
	//	}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v ticket booking system\n", conferenceName)
	fmt.Println("Total ticket available -", conferenceTickets, "/ Only", remainingTickets, "tickets are left")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint

	// Get username as input from User
	fmt.Println("Enter your Firstname: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Lastname: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&emailAddress)

	fmt.Println("How many tickets you would like to book?: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAddress, userTickets

}

func bookTicket(firstName string, lastName string, userTickets uint, emailAddress string) {
	remainingTickets = remainingTickets - userTickets

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		emailAddress:    emailAddress,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. We will send you confirmation shortly via email at %v.\n", firstName, lastName, userTickets, emailAddress)

	fmt.Printf("%v tickets are left\n", remainingTickets)
}

func sendTicket(firstName string, lastName string, userTickets uint, emailAddress string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, emailAddress)
	fmt.Println("###################")
	wg.Done()
}
