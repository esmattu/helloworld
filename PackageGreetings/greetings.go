package PackageGreetings

import "fmt"

//Hellow returns a greeting for the name person
func Greeting(name string) string {
	//Return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message

}
