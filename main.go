package main

import (
	"fmt"

	"example.com/helloworld/PackageGreetings"
	"example.com/helloworld/PackageHello"
	"github.com/esmattu/dateformat"
	"rsc.io/quote"
)

func main() {

	fmt.Println(PackageHello.Hello())
	fmt.Println(quote.Go())

	//get a message from another package
	message := PackageGreetings.Greeting("Gildri")
	fmt.Println(message)
	//print the current date with the sent format, from another module
	dateFormatted := dateformat.DateFormatterDB("01/02/06")
	fmt.Println(dateFormatted)
}
