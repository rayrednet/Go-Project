package main

import (
	"Go-Project/helper"
	"fmt"
	"sync"
)

const conferenceTickets uint = 50
var conferenceName  = "Go Conference"
var remainingTickets uint = 50

var wg = sync.WaitGroup{}

func main() {
  for {
	  helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

		firstName, lastName, email, userTickets := helper.GetUserInput()
		isValidName, isValidEmail, isValidTicketsNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketsNumber {
			helper.BookTickets(userTickets, firstName, lastName, email, &remainingTickets, conferenceName)
			
			wg.Add(1)
			go helper.SendTicket(userTickets, firstName, lastName, email, &wg)

			firstNames := helper.GetFirstNames()
			fmt.Printf("The first names of the attendees are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("\n🎉 All tickets are sold out! 🎉")
				fmt.Println("🙏 Thank you for your interest and for booking with us!")
				fmt.Println("==========================================")
				break
			}
		} else {
			fmt.Println()
			if !isValidName {
				fmt.Println("❌ Invalid Name: First and last name should be at least 2 characters long.")
			}
			if !isValidEmail {
				fmt.Println("❌ Invalid Email: Email should contain both '@' and '.' symbols.")
			}
			if !isValidTicketsNumber {
				fmt.Printf("❌ Invalid Ticket Number: Please enter a number between 1 and %v.\n", remainingTickets)
			}

			fmt.Println("🚫 Please correct the above errors and try again.")
		}
		wg.Wait()
	}
}
