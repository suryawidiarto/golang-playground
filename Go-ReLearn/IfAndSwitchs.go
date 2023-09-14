package main

import (
	"fmt"
	"math"
)

func main() {
	//! IF STATEMENT
	if true {
		fmt.Println("Running")
	}

	statePopulations := map[string]int{
		"Purwokerto": 9000000,
		"Jakarta":    12312312312,
		"Semarang":   1000000,
	}

	// we can define if with initializer, this is similiar like try with param on java
	if value, ok := statePopulations["Purwokerto"]; ok {
		fmt.Println(value)
	}

	number := 50
	guess := 30
	//? guess := -5 if we use this returnTrue will not be called

	//? its because when the first condition is already satisfied, then the other condition will not get tested on OR || operator. this is called short circuiting
	if guess < 1 || returnTrue() {
		fmt.Println("Guess must be above 1")
	} else if guess > 100 {
		fmt.Println("Guess must be under 100")
	} else {
		fmt.Println("hwehwehwehw")
	}

	//? if one of the first parameter is return false, then it will not testing the next parameter for AND && operator
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Guess Too Low")
		}
		if guess > number {
			fmt.Println("Guess Too High")
		}
		if guess == number {
			fmt.Println("Good Job, You Got It!")
		}
	}

	fmt.Println(number <= guess, number >= guess, number != guess)

	myNum := 0.123

	// myNum == math.Pow(math.Sqrt(myNum), 2) -> bad way before as it will not calculating precisely
	// but this just more a good way a bit as it is also didnt really precise, and need to use other precise condition
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 { //divide the number and if they're within a 10% of each other then we will consider it is the same
		fmt.Println("myNum is same")
	} else {
		fmt.Println("myNum is different")
	}

	//! SWITCH STATEMENT

	//switch with tags, each case cannot overlap with each other and need to be unique
	// we can also use initializer in switch
	switch i := 2 + 3; i {
	//in go we can use multiple condition on one case statement separated by ,
	// but each case need to be unique, if not it will throw error as the condition is duplicated
	case 1, 4, 6:
		fmt.Println("one, four, six")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("no value match")
	}

	//another way to define switch
	j := 10
	//switch without tags can be overlapped with each other
	switch {
	case j <= 10:
		fmt.Println("value less than equal 10")
		fallthrough // will execute the next case wether its true or false. the default behavior in switch is after case condition is fulfilled it will break and stop
	case j <= 20:
		fmt.Println("value less than equal 20")
	default:
		fmt.Println("value greater than 20")
	}

	var k interface{} = "strings"
	switch k.(type) {
	case int:
		fmt.Println("k is a int")
		break // we can add break in here to force fully stop the switch
		fmt.Println("dont want to print this if the case condition is true")
	case float32:
		fmt.Println("k is a float32")
	case string:
		fmt.Println("k is string")
	default:
		fmt.Println("k is another type")
	}
}

func returnTrue() bool {
	return true
}
