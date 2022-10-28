package main

import (
	"booking-app/helper"
	"fmt"
)

/* Package variables */
const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets = conferenceTickets

/* slice of maps with key and value as strings
var bookings = make([]map[string]string, 0)
*/
/* slice of structs */
var bookings = make([]UserData, 0)

type UserData struct {
	firstname   string
	lastname    string
	email       string
	ticketCount uint
}

func main() {
	greetUsers()

	for remainingTickets > 0 {
		/* Get user input */
		userFirstName, userLastName, userEmail, userTickets := getUserInput()

		/* Validate user input */
		isValidName, isValidEmail, isValidTicketCount :=
			helper.ValidateUserInput(userFirstName, userLastName, userEmail, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketCount {
			bookTickets(userFirstName, userLastName, userEmail, userTickets)
			firstNames := getFirstNames()
			fmt.Printf("List of people who booked so far %v\n\n", firstNames)
		} else {
			if !isValidName {
				fmt.Printf("Firstname %v or Lastname %v you entered is too short. Try again.\n", userFirstName, userLastName)
			}
			if !isValidEmail {
				fmt.Printf("Email %v is invalid, doesn't contain @. Try again.\n", userEmail)
			}
			if !isValidTicketCount {
				fmt.Printf("Number of tickets %v exceeds available tickets %v. Try again.\n", userTickets, remainingTickets)
			}
		}
	}

	if remainingTickets == 0 {
		fmt.Println("Conference is completely booked! Come again next time")
	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking site!\n", conferenceName)
	fmt.Printf("We have %v tickets left out of %v tickets available.\n", remainingTickets, conferenceTickets)
	fmt.Printf("You can book the conference tickets here.\n\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, user := range bookings {
		firstNames = append(firstNames, user.firstname)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint

	fmt.Printf("Please enter your first name\n")
	fmt.Scan(&userFirstName)
	fmt.Printf("Please enter your last name\n")
	fmt.Scan(&userLastName)

	fmt.Printf("Please enter your email\n")
	fmt.Scan(&userEmail)
	fmt.Printf("Please enter the number of tickets you want to reserve\n")
	fmt.Scan(&userTickets)

	return userFirstName, userLastName, userEmail, userTickets
}

func bookTickets(userFirstName string, userLastName string, userEmail string, userTickets uint) {
	remainingTickets -= userTickets
	var userData = UserData{
		firstname:   userFirstName,
		lastname:    userLastName,
		email:       userEmail,
		ticketCount: userTickets,
	}
	/*
		var userData = make(map[string]string)
		userData["firstname"] = userFirstName
		userData["lastname"] = userLastName
		userData["email"] = userEmail
		userData["ticketCount"] = strconv.FormatUint(uint64(userTickets), 10)
	*/

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v shortly.\n",
		userFirstName, userLastName, userTickets, userEmail)
	fmt.Printf("There are %v tickets that are remaining\n", remainingTickets)
}
