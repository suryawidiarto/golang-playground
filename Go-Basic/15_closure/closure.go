package main

import "fmt"

//scope in go
func main() {
	counter := 0
	increment := func() {
		//will shadowing variable
		counter := 2
		fmt.Println("Increment")
		//and this counter will refer to variable inside function
		counter++ // this will update upper bracket variable, because it can access variable 1 bracket outside
	}

	increment()
	increment()

	fmt.Println(counter)
}

//variable outside function can be updated by inner function
//we need to redeclare if we want the variable is only for the function, so it not changing outside variable scope
