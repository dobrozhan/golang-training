package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// // Assign
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"

	emails := map[string]string{"Bob":"bob@gmail.com", "Sharon":"sharon@gmail.com"}

	fmt.Println(emails)
	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
