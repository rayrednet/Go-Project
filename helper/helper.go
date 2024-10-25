package helper

import (
  "strings"
  "fmt"
  "time"
  "sync"
)

var bookings = make([]UserData, 0)

type UserData struct {
  FirstName string
  LastName string
  Email string
  NumberOfTickets uint
}

func GreetUsers(conferenceName string, conferenceTickets uint, remainingTickets uint) {
  fmt.Printf("Welcome to %v booking application\n", conferenceName)
  fmt.Printf("We have the total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
  fmt.Println("Get your tickets here to attend")
}

func GetUserInput() (string, string, string, uint) {
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

  return firstName, lastName, email, userTickets
}

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
    isValidName := len(firstName) >= 2 && len(lastName) >=2
    isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
    isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets
    return isValidName, isValidEmail, isValidTicketsNumber
  }

func BookTickets (userTickets uint, firstName string, lastName string, email string, remainingTickets *uint, conferenceName string) {
  *remainingTickets -= userTickets

  var userData = UserData{
    FirstName: firstName,
    LastName: lastName,
    Email: email,
    NumberOfTickets: userTickets,
  }

  bookings = append(bookings, userData)
  fmt.Printf("List of bookings: %v\n", bookings)

  fmt.Printf ("Thank you %v %v for booking %v tickets for %v. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, conferenceName, email)
  fmt.Printf("There are %v tickets remaining for %v\n", *remainingTickets, conferenceName)
}

func GetFirstNames() []string {
  firstNames := []string{}
      for _, booking := range bookings {
        firstNames = append(firstNames, booking.FirstName)
      }
      return firstNames
}

func SendTicket(userTickets uint, firstName string, lastName string, email string, wg *sync.WaitGroup) {
  time.Sleep(50 * time.Second)
  var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
  fmt.Println("###########################")
  fmt.Printf("Sending ticket:\n%v\nhas been sent to email address %v\n", ticket, email)
  fmt.Println("###########################")
  wg.Done()
}