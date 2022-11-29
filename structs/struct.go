package structs

import "fmt"

// Same as Class in Java
// Different data types
type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var bookings = make([]UserData, 0)

func AddValuesToStruct(firstName string, lastName string, email string, userTickets uint) {
	userData := UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
}

func PrintStructValues() {
	fmt.Printf("Values from Struct: %v\n", bookings)
}
