# Go-Project

This Go project is part of my course, **Pemrograman Berbasis Kerangka Kerja (D)**. The repository will be continuously updated as I follow the video tutorials provided by the YouTube channel linked with this course.

## Identity

- **Name:** Rayssa Ravelia  
- **NRP:** 5025211219  
- **Class:** Pemrograman Berbasis Kerangka Kerja (D)

---

## Project Overview

The project is based on a tutorial from the YouTube video: [Golang Tutorial for Beginners | Full Go Course by TechWorld with Nana](https://youtu.be/yyUHQIec83I?si=s1RoXRBlNdCXgo8g). The video is divided into two sections for clarity.

- **Section 1:** Introduction to Go, basic syntax, and control structures (00:00 - 1:54:02)
- **Section 2:** Functions and advanced topics (1:58:37 - 3:23:51)

You can follow each commit to see how the project evolves. View screenshots demonstrating the code output in each section.

---

## ⭐ Section 1 - Introduction to Go (00:00 - 1:54:02)

### Key Topics Covered:

1. **Introduction to Go**  
   - Overview of Go language and its advantages.
   
2. **Go Setup**  
   - Installation and setup of Go and the text editor.

3. **Basic Syntax & First Program**  
   - Structure of a Go file, writing the first program.

4. **Variables & Constants**  
   - Using variables and constants in Go.

5. **Formatted Output (printf)**  
   - Formatting and displaying output using `printf`.

6. **Data Types**  
   - Overview of Go data types.

7. **User Input**  
   - How to get user input using `fmt.Scan`.

8. **Pointers**  
   - Introduction to pointers in Go.

9. **Logic for Booking Tickets**  
   - Writing logic for booking tickets, using arrays and slices.

10. **Loops & Conditionals**  
    - Looping and conditionals with `if`, `else`, and `switch`.

11. **User Validation**  
    - Validating user input for names, email, and ticket count.

---

### Final Code for Section 1
This code below can be viewed in the `/main.go` file

```go
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

  fmt.Printf("Welcome to %v booking application\n", conferenceName)
  fmt.Printf("We have the total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
  fmt.Println("Get your tickets here to attend")

  for {
      var firstName, lastName, email string
      var userTickets uint

      fmt.Println("Please enter your first name:")
      fmt.Scan(&firstName)

      fmt.Println("Please enter your last name:")
      fmt.Scan(&lastName)

      fmt.Println("Please enter your email:")
      fmt.Scan(&email)

      fmt.Println("Please enter the number of tickets you want to book:")
      fmt.Scan(&userTickets)

      isValidName := len(firstName) >= 2 && len(lastName) >= 2
      isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
      isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

      if isValidName && isValidEmail && isValidTicketsNumber {
        remainingTickets -= userTickets
        bookings = append(bookings, firstName+" "+lastName)

        fmt.Printf("Thank you %v %v for booking %v tickets for %v. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, conferenceName, email)
        fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)

        firstNames := []string{}
        for _, booking := range bookings {
          var names = strings.Fields(booking)
          firstNames = append(firstNames, names[0])
        }
        fmt.Printf("The first names of the attendees are %v\n", firstNames)

        if remainingTickets == 0 {
          fmt.Println("All tickets are sold out. Thank you for booking!")
          break
        }
      } else {
        if !isValidName {
          fmt.Println("Invalid name. First and last name must be at least 2 characters long.")
        }
        if !isValidEmail {
          fmt.Println("Invalid email. Email must contain an '@' and '.'")
        }
        if !isValidTicketsNumber {
          fmt.Printf("Invalid number of tickets. Please enter a number between 1 and %v\n", remainingTickets)
        }
        continue
      }
  }
}
```

### Output Example:

- **Correct Input:**  
   - Two users, John and Jane, have booked tickets. The application shows a summary of the attendees' names after each booking.
   ![Correct Input](img/section1-correct_input.png)

- **Incorrect Input:**  
   - For incorrect inputs (e.g., name with less than 2 characters, invalid email, or exceeding available tickets), the application will display specific error messages and ask the user to try again.
   ![Wrong Input](img/section1-wrong_input.png)

---

## Summary

In this section, we:
- Covered the basics of Go, including syntax, variables, and user input.
- Implemented logic for booking conference tickets.
- Added validation for user input, ensuring valid names, emails, and ticket numbers.

Further updates will be made as we progress through Section 2, which covers functions and more advanced Go topics.

