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

package main
import "fmt"

func main(){
	// var conferenceName = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

    // fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still remaining\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// var bookings = [50]string{"Dipanshu","Nicole","Patrick"}
	// empty array
	var bookings [50]string
	
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

	  remainingTickets = remainingTickets - userTickets
	  bookings[0] = firstName + " " + lastName

	  fmt.Printf("The whole array: %v\n",bookings)
	  fmt.Printf("The first value: %v\n",bookings[0])
	  fmt.Printf("Array type: %T\n",bookings)
	  fmt.Printf("Array length: %v\n", len(bookings))
	
    fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a comfirmation email at %v \n ",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets are remaining for %v\n",remainingTickets,conferenceName)
	
}


