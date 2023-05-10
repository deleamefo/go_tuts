package main

import "fmt"

func main() {
	hotelName := "The Gopher Hotel"

	const rooms = 100
	var roomsAvailable uint = 50

	fmt.Printf("Welcome to %v\n", hotelName)
	fmt.Println("Please book a room to stay")

	//roomsAvailable = rooms - roomsOccupied
	fmt.Printf("We have total of %v rooms and %v rooms are still available\n", rooms, roomsAvailable)

	// Get user user input
	var firstName string
	var lastName string
	var email string
	var userRooms int

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of rooms: ")
	fmt.Scan(&userRooms)

	fmt.Printf("User %v booked %v rooms. \n", firstName, userRooms)
	fmt.Print("Thank you for booking a room here")
	fmt.Print("Do have a nice stay")
}
