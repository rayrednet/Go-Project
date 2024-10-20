package main
import "fmt"

func main() {
  conferenceName  := "Go Conference"
  const conferenceTickets uint = 50
  var remainingTickets uint = 50

  fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

  fmt.Printf("Welcome to %v booking application\n", conferenceName)
  fmt.Printf("We have the total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
  fmt.Println("Get your tickets here to attend")

  var userName string
  var userTickets int
  // asking for user name

  userName = "John Doe"
  userTickets = 2
  fmt.Printf ("Hello %v, you have booked %v tickets\n", userName, userTickets)
   
}
