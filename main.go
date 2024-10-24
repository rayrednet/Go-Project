package main
import (
  "fmt"
  "strings"
)

func main() {
  conferenceName  := "Go Conference"
  const conferenceTickets uint = 50
  var remainingTickets uint = 50
  bookings := [] string{}

  fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

  fmt.Printf("Welcome to %v booking application\n", conferenceName)
  fmt.Printf("We have the total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
  fmt.Println("Get your tickets here to attend")

  for {
      var firstName string
      var lastName string
      var email string
      var userTickets uint

      fmt.Println("Please enter your first name:")
      fmt.Scan(&firstName)

      fmt.Println("Please enter your last name:")
      fmt.Scan(&lastName)

      fmt.Println("Please enter your email:")
      fmt.Scan(&email)

      fmt.Println("Please enter the number of tickets you want to book:")
      fmt.Scan(&userTickets)

      isValidName := len(firstName) >= 2 && len(lastName) >=2
      isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
      isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

      if isValidName && isValidEmail && isValidTicketsNumber {
        remainingTickets -= userTickets
        bookings = append(bookings, firstName + " " + lastName)

        fmt.Printf ("Thank you %v %v for booking %v tickets for %v. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, conferenceName, email)
        fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)

        firstNames := []string{}
        for _, booking := range bookings {
          var names = strings.Fields(booking)
          firstNames = append(firstNames, names[0])
        }
        fmt.Printf ("The first names of the attendees are %v\n", firstNames)
        
        if remainingTickets == 0 {
          fmt.Println("All tickets are sold out. Thank you for booking")
          break
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
        continue
      }
  }
  }
