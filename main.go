package main

import (
	"Go-Project/helper"
	"fmt"
  "sync"
)

const conferenceTickets uint = 50
var conferenceName  = "Go Conference"
var remainingTickets uint = 50

var wg = sync.WaitGroup {}

func main() {
  helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)
  firstName, lastName, email, userTickets := helper.GetUserInput()
  isValidName, isValidEmail, isValidTicketsNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

  if isValidName && isValidEmail && isValidTicketsNumber {
    helper.BookTickets(userTickets, firstName, lastName, email, &remainingTickets, conferenceName)
    
    wg.Add(1)
    go helper.SendTicket(userTickets, firstName, lastName, email, &wg)

    firstNames := helper.GetFirstNames()
    fmt.Printf ("The first names of the attendees are %v\n", firstNames)     
    
    if remainingTickets == 0 {
      fmt.Println("All tickets are sold out. Thank you for booking")
      
    }
  } else {
    if !isValidName {
      fmt.Println("Invalid name. First name and last name should be at least 2 characters long")
    }
    if !isValidEmail {
      fmt.Println("Invalid email. Email should contain @ and . sign")
    }
    if !isValidTicketsNumber {
      fmt.Println("Invalid number of tickets. Please enter a number between 1 and", remainingTickets)
    }
  }
  wg.Wait()
  } 