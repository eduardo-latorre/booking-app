package main

import (
	"booking-app/common"
	"fmt"
	"strings"
)

// Package level variable
// They must be created using var keyword
const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50 //uint no negative as int

func main() {

	//Pointers
	fmt.Println(conferenceName)  //Variable value
	fmt.Println(&conferenceName) //Memory address or special variable

	greetUser()

	//Arrays - specific length
	// bookings = [50]string{}
	// bookings = [50]string{"value1", "value2"}
	// var bookings [50]string
	// bookings[0] = firstName + " " + lastName

	//Slices - no length needed
	//var bookings = []string{}
	bookings := []string{}

	//Loops, there's only for loop in GO
	//for <condition(s)>{ }
	//for (endless loop)
	//for true (endless loop)
	for remainingTickets > 0 {

		firstName, lastName, email, userTickets, city := scaningUserInfo()
		isNameValid, isLastName, isEmailValid, isTicketsValid := common.ValidateUserInput(firstName, lastName, email, userTickets)

		//Conditionals
		//Simple way just: if remainingTickets < userTickets{ }
		//Or the one bellow
		if !isNameValid {
			fmt.Println("Firstname entered are invalid")
			continue
		} else if !isLastName {
			fmt.Println("Lastname entered is invalid")
			continue
		} else if !isEmailValid {
			fmt.Println("Email entered is invalid")
			continue
		} else if !isTicketsValid {
			fmt.Println("TIckes entered are invalid")
			continue
		}

		//Switch validation
		switch city {
		case "Singapour":
			//Add logic here
		case "Hongkong", "Berlin":
			//Add logic here
		default:
			fmt.Println("City entered is not correct")
			continue
		}

		//Remaining tickets validation
		enoughTickets := remainingTickets < userTickets
		if enoughTickets {
			fmt.Printf("We're sorry we have only %v tickets left\n", remainingTickets)
			continue
		}

		//Using returning value function
		remainingTickets = getRemainingTickets(userTickets)

		fmt.Printf("Thank %v %v for booking %v tickets, you'll get your recieve to %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("Remaining tickets: %v\n", remainingTickets)

		//Adding values to Slice
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Whole array or slice: %v\n", bookings)
		fmt.Printf("First array or slice position: %v\n", bookings[0])
		fmt.Printf("Array or slice lenght: %v\n", len(bookings))

		//Getting the type of variables
		fmt.Printf("userName var is type %T\n", firstName)
		fmt.Printf("userName var is type %T\n", userTickets)

		printFirstNames(bookings)
	}

	//Specifying var types, used when specific types and automatically validated, there so many
	var int8 int8 = 99
	fmt.Printf("Valor de int8 %v\n", int8)
	var int16 int16 = 9999
	fmt.Printf("Valor de int8 %v\n", int16)
	var int32 int32 = 999999999
	fmt.Printf("Valor de int8 %v\n", int32)
	var int64 int64 = 999999999999999999
	fmt.Printf("Valor de int8 %v\n", int64)

	//two ways of implicit definying var
	var var1 = "value1"
	var2 := "value2"

	fmt.Printf("variable 1: %v y variable 2: %v\n", var1, var2)
}

func greetUser() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName) //Printf (printformat) using variables inside
	fmt.Println("We have a total of", conferenceTickets, "and there are", remainingTickets, "remaining")
	fmt.Println("Please get your tickets here to attend")
}

func printFirstNames(bookings []string) {
	firstNames := []string{}
	//For each (index, var with content, range)
	//_ stands for a blank variable
	for _, booking := range bookings {
		name := strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}
	fmt.Printf("Names booked are: %v\n", firstNames)
}

// Function with returning value
func getRemainingTickets(userTickets uint) uint {
	updatedRemainingTickets := remainingTickets - userTickets
	return updatedRemainingTickets
}

// Function with multiple values
func scaningUserInfo() (string, string, string, uint, string) {
	//Types of variables
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var city string

	//Scanning values values, pointers needed to store incoming values
	fmt.Println("Enter your firstname: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your lastname: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	fmt.Println("Enter city location: ")
	fmt.Scan(&city)

	return firstName, lastName, email, userTickets, city
}
