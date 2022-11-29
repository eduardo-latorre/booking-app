package maps

import (
	"fmt"
	"strconv"
)

// This is a emty Slice of Maps, so it's necesary to add size
var bookings = make([]map[string]string, 0)

// Maps only uses a defined type value
func AddingUserMap(firstName string, lastName string, email string, userTickets uint, city string) {
	var userMap = make(map[string]string)
	userMap["firstName"] = firstName
	userMap["lastName"] = lastName
	userMap["email"] = email
	// As this maps only uses string, we need to convert different values
	userMap["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	userMap["city"] = city

	bookings = append(bookings, userMap)
}

func GettingMapValues() {
	fmt.Println("Map length: ", len(bookings))

	for _, bookingObject := range bookings {
		fmt.Printf("Values from Map: %v\n", bookingObject)
	}
}
