package array

import "fmt"

// Arrays - specific length
// bookings = [50]string{}
// bookings = [50]string{"value1", "value2"}
// bookings[0] = firstName + " " + lastName
var bookings [50]string
var index uint = 0

func AddingValuesToArray(firstName string) {
	bookings[index] = firstName
}

func GettingArrayValues() {
	fmt.Printf("Array index %v value %v\n", index, bookings[index])
}
