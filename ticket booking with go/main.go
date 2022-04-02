package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type data struct {
	name               string
	email              string
	tickets            int
	optedFornewsletter bool
}

var wg = sync.WaitGroup{}

func main() {
	name := "LOL GANG"      // this alt syntax cannot be used if u want to create a constant
	const tickets int = 100 // immutable , cannot change the value
	var remainingTickets int = 100
	var booking = make([]data, 0) // slice

	fmt.Println("WELCOME TO THE", name, " BOOKING CENTER")
	fmt.Printf("we have a total of %v tickets and %v are left\n", tickets, remainingTickets)
	fmt.Println("GO HERE TO ACCESS THE TICKETS")

	for remainingTickets > 0 && len(booking) <= 100 {

		var username string
		var usertickets int
		var email string
		//asking for the user for thier name

		fmt.Println("Enter your name:")
		fmt.Scan(&username)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets for booking:")
		fmt.Scan(&usertickets)

		isValid := len(username) >= 2
		Validemail := strings.Contains(email, "@")
		isvalidticketnumber := usertickets > 0 && usertickets <= remainingTickets

		if isValid && Validemail && isvalidticketnumber {
			wg.Add(1)
			go sendTicket(usertickets, username, email)
			remainingTickets = remainingTickets - usertickets

			var userdata = data{
				name:    username,
				email:   email,
				tickets: usertickets,
			}

			booking = append(booking, userdata)
			fmt.Println(booking)
			fmt.Printf("the slice:%v\n", booking)

			fmt.Printf("Thank you %v for booking %d tickets.You will recieve a confirmation email at %v", username, usertickets, email)
			fmt.Printf("\n %v tickets remaining for %v\n", remainingTickets, name)
			if remainingTickets == 0 {
				fmt.Println("All the tickets are booked,come back when its available")
				break
			}

		} else {
			if !isValid {
				fmt.Println(" Please Enter a valid name, try again")
			}
			if !isvalidticketnumber {
				fmt.Println(" Please Enter a valid number, try again")
			}
			if !Validemail {
				fmt.Println("Enter a valid email address")
			}
			continue

		}

	}
	wg.Wait()

}
func sendTicket(usertickets int, username string, email string) {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v ", usertickets, username)
	fmt.Println("################################")
	fmt.Printf("Sending ticket:\n %v to email address \n %v", ticket, email)
	fmt.Println("THANK YOU FOR BOOKING!!!")
	fmt.Println("################################")
	wg.Done()
}
