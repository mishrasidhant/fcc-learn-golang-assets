package main

import "fmt"

func main() {
	// You can use multiple _'s, it has syntactical significance in go
	firstName, _, _:= getNames()
	fmt.Println("Welcome to Textio,", firstName)
}

// don't edit below this line

func getNames() (string, string, string) {
	return "John", "Doe", "testing"
}
