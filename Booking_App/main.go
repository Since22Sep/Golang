// go mod init<module path>
// creates a new module
// module path can correspond to a repository you plan to publish your module to.
/*
this go mod init <module.path>
-> initialized a go.mod file
-> describes the module : with name/module path and go version used in the program.
-> the module path is also the import path
*/

// All our code must belong to a package
// then missing a declaration means go needs to find where to start the program from.
// now after we got the main function we need to give the package as well .
/* Go programs are organized into packages.
   Go's standard library, provides diff core packages for us to use.
   "fmt" is one of these which we can use by importing it.
   It is a collection of source files.
*/

/*
Go is a statically typed language :- you need to tell Go compiler the data type when declaring the variable.
Type inference:- but Go can infer the type when you assign the value.
*/

/*
Scan function can now assign the user's value to the userName varibale beacause it has a pointer to its memory address.

*/

// A list that is more dynamic in size where we dont need to specify a size at the beginning , automatically expands when new elements are added
/*
-> Slice is an abstraction of an Array
-> They are more flexible and powerful: variable-length or get an sub-array of its own
-> Slices are also index-based and have a size, but is resized when needed.
-> append adds the element at the end of the slice
-> grows the slice if a greater capacity is needed and returns the updated slice value.
*/

/* Range iterates over elements for different data structures , for arrays and slices range provides the index and value for each element. */

// Blank Identifier:- to ignore a variable you dont want to use , so with go you need to make unused variables explicit

package main

import (
	"fmt"
	"strings"
)

func main(){
	// var conferenceName = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{} 

    // fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still remaining\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// var bookings = [50]string{"Dipanshu","Nicole","Patrick"}
	// empty array
	// var bookings [50]string
	// slice
	// var bookings []string

	for {
	
	  var firstName string
	  var lastName string
	  var email string
	  var userTickets uint



// ask user for their name
      fmt.Println("Enter your first name: ")
	  fmt.Scan(&firstName)

	  fmt.Println("Enter your last name: ")
      fmt.Scan(&lastName)

	  fmt.Println("Enter your email address: ")
	  fmt.Scan(&email)

	  fmt.Println("Enter the number of tickets you want: ")
      fmt.Scan(&userTickets)

	   isValidName := len(firstName) >= 2 && len(lastName) >= 2
	   isValidEmail := strings.Contains(email, "@")
	   isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	   

	  if isValidName && isValidEmail && isValidTicketNumber {
		remainingTickets = remainingTickets - userTickets
	//   bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a comfirmation email at %v \n ",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets are remaining for %v\n",remainingTickets,conferenceName)

	firstNames := []string{}

	// this loop ends when iterated over all elements of the bookings list
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("These are all our bookings: %v\n", firstNames)

	if remainingTickets == 0 {
		// end program
		fmt.Println("Our conference is booked out. Come back next year.")
		break
	}
	
	  }else
	  {
		if !isValidName {

			fmt.Println("first or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
		}

		  if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid.")
		}
       
	  }
	}
	

	//   fmt.Printf("The whole slice: %v\n",bookings)
	//   fmt.Printf("The first value: %v\n",bookings[0])
	//   fmt.Printf("Slice type: %T\n",bookings)
	//   fmt.Printf("Slice length: %v\n", len(bookings))
	
    city := 
	
}


