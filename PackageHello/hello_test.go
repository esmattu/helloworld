package PackageHello_test

import (
	"testing"

	"example.com/helloworld/PackageHello"
)

func TestHello(t *testing.T) {

	if PackageHello.Hello() != "Hello World" {
		t.Fatal("Wrong greeting :(")
	}

}
