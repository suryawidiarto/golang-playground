package helper

import "fmt"

/**
in golang, we can set access modifier to a variable or function with its naming
SayHello -> with capital on the first letter indicating that it can be accessed from outside package
sayHello -> camelCase indicating that it cannot be accessed from outside package
*/

var version = "1.0.0"            //cannot be accessed from outside package
var VersionApplication = "1.2.2" //can be accessed from outside package

//cannot be accessed from outside package
func sayHello(name string) string {
	return "Hello " + name
}

//can be accessed from outside package
func SayHello(name string) {
	fmt.Println("Hello", name)
}
