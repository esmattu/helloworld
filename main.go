package main

import (
	"fmt"

	"example.com/helloworld/PackageHello"
	"rsc.io/quote"
)

func main() {

	fmt.Println(PackageHello.Hello())
	fmt.Println(quote.Go())

}
