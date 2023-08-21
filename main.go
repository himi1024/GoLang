package main

// Package import
import (
	"fmt"
	"golang/helper"
)

// variables
var remainingTickets uint = 20

const confereneTickets int = 50

var conferenceName = "Go Conference"

// Arrays
// Slices (Dynamic Arrays)
var bookings = make([]UserData, 0)

// Object
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUser()

	// for {
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names %v\n", firstNames)

		// Tickets sold
		if remainingTickets == 0 {
			fmt.Println("SOLD OFF")
		}
	} else {
		if !isValidName {
			fmt.Println("firstname or lastname you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email should contain a @ symbol")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	// }

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
}
func greetUser() {
	fmt.Printf("Welcome to %v", conferenceName)
	fmt.Printf("There is %v tickets and %v are still available.\n", confereneTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
