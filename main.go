package main

import (
	"fmt"

	"example.com/helloworld/PackageGreetings"
	"example.com/helloworld/PackageHello"
	"rsc.io/quote"
)

func main() {

	fmt.Println(PackageHello.Hello())
	fmt.Println(quote.Go())

	//get a message from another package
	message := PackageGreetings.Greeting("Gildri")
	fmt.Println(message)

}
