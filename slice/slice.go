package slice

import (
	"fmt"
	"strings"
)

// Slice - no length needed
// var bookings = []string{}
var bookings = []string{}

func AddingNamesToSlice(firstName string, lastName string) {
	// Adding values to Slice
	bookings = append(bookings, firstName+" "+lastName)
}

func PrintSliceItems() {
	firstNames := []string{}
	//For each (index, var with content, range)
	//_ stands for a blank variable
	for _, booking := range bookings {
		name := strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}
	fmt.Printf("Values form Slice: %v\n", firstNames)
	fmt.Printf("Slice lenght: %v\n", len(bookings))
}
