package main

import "fmt"

func main() {
	//! Looping
	//standard loop
	//i is in for scope only
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	//with 2 variable
	// in go ++ is not an expression so we can't use it as part of the statement
	//? i, j = i++, j++ -> will throw error
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	//this also valid
	i := 0 // i is in main func scope
	for ; i < 3; i++ {
		fmt.Println(i)
	}

	//this also valid, and will work like a while
	j := 0 // i is in main func scope
	for j < 3 {
		fmt.Println(j)
		j++
	}

	//in case we need infinite loop and only stop on specific condition
	for {
		fmt.Println(i)
		if i == 5 {
			break // will stop the loop
		}
		i++
	}

	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("odd number only -", i)
	}

Loop: // define labels, and tell go where to break to. as Loop is defined on the outer part of for i := 1; i <= 3; i++ it will break to outer of it
	// this was useful to automatically break to the upper level loop without need to break the outer part from inside the child loop
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	//loop through an slice or array
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	//use _ to tell go to skip the variable because it is not used, but needed for param positioning to get v
	for _, v := range s {
		fmt.Println(v)
	}

	//loop through a map
	statePopulations := map[string]int{
		"Purwokerto": 9000000,
		"Jakarta":    12312312312,
		"Semarang":   1000000,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	//loop through string
	strings := "Learn!"
	for k, v := range strings {
		fmt.Println(k, string(v))
	}
}
