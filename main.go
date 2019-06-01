// executable package
package main

import "fmt"

// will be called when we execute this file
func main() {
	// create a variable to turn into a byte slice
	greeting := "Hi there!"
	// use type conversion to print greeting to console
	// as byte slice
	fmt.Println([]byte(greeting))
}
