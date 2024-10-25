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
  fmt.Println()
  fmt.Println()
  fmt.Println("=============================================================")
  fmt.Printf("🎉 Welcome to the %v Booking Application 🎉\n", conferenceName)
  fmt.Println("=============================================================")

  fmt.Printf("🎟️  Total Tickets: %v\n", conferenceTickets)
  fmt.Printf("🎫  Tickets Remaining: %v\n\n", remainingTickets)

  fmt.Println("✨ Get your tickets here to attend and enjoy an unforgettable experience! ✨")
  fmt.Println("-------------------------------------------------------------")
}

func GetUserInput() (string, string, string, uint) {
  var firstName string
  var lastName string
  var email string
  var userTickets uint

  fmt.Println()
  fmt.Println("👤 Let's get to know you! Please enter your details below to order your ticket👇")

  fmt.Println("\n📝 First Name:")
  fmt.Print("👉 ")
  fmt.Scan(&firstName)

  fmt.Println("\n📝 Last Name:")
  fmt.Print("👉 ")
  fmt.Scan(&lastName)

  fmt.Println("\n📧 Email Address:")
  fmt.Print("👉 ")
  fmt.Scan(&email)

  fmt.Println("\n🎟️ Number of Tickets:")
  fmt.Print("👉 ")
  fmt.Scan(&userTickets)

  fmt.Println("\n✅ Thank you! You've entered your information successfully")

  return firstName, lastName, email, userTickets
}

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
    isValidName := len(firstName) >= 2 && len(lastName) >=2
    isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
    isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets
    return isValidName, isValidEmail, isValidTicketsNumber
  }

  func BookTickets(userTickets uint, firstName string, lastName string, email string, remainingTickets *uint, conferenceName string) {
    fmt.Println() 

    *remainingTickets -= userTickets

    var userData = UserData{
        FirstName:      firstName,
        LastName:       lastName,
        Email:          email,
        NumberOfTickets: userTickets,
    }

    bookings = append(bookings, userData)
    fmt.Printf("📝 List of Bookings: %v\n", bookings)

    fmt.Printf("🎉 Thank you, %v %v, for booking %v tickets for %v!\n", firstName, lastName, userTickets, conferenceName)
    fmt.Printf("📧 A confirmation email will be sent to: %v\n", email)
    fmt.Printf("🎟️ Remaining Tickets: %v for %v\n", *remainingTickets, conferenceName)
     fmt.Println("-------------------------------------------------------------")
}


func GetFirstNames() []string {
  firstNames := []string{}
      for _, booking := range bookings {
        firstNames = append(firstNames, booking.FirstName)
      }
      return firstNames
}

func SendTicket(userTickets uint, firstName string, lastName string, email string, wg *sync.WaitGroup) {
  fmt.Println() 

  time.Sleep(10 * time.Second)
  var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

  fmt.Println("###########################")
  fmt.Println("🎟️  Sending Ticket... 🎟️")
  fmt.Printf("📧 Ticket Details:\n%v\nConfirmation email has been sent to: %v\n", ticket, email)
  fmt.Println("###########################")
  fmt.Println("✅ Ticket sent successfully!")

  wg.Done()
}
