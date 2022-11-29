package common

import "strings"

// In Java this is like a public variable
var VariableScope string = "This is a variable scope from common package"

// Using capital letter allow to export the function (Public as in Java)
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool, bool) {
	//Conditionals
	isNameValid := len(firstName) >= 2
	isNameLastname := len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isTicketsValid := userTickets > 0

	return isNameValid, isNameLastname, isEmailValid, isTicketsValid
}
