package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke - ", counter)
		counter++
	}

	//  init statement, condition, post statement
	for count := 1; count <= 10; count++ {
		fmt.Println("# Perulangan ke - ", count)
	}

	slice := []string{"Surya", "Adi", "Widiarto"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, v := range slice {
		fmt.Println("index -> ", i, " - value -> ", v)
	}

	// _ underscore will tell GO that the variable is not used
	for _, v := range slice {
		fmt.Println("value -> ", v)
	}

	person := make(map[string]string)
	person["K1"] = "value 1"
	person["K2"] = "value 2"
	person["K3"] = "value 3"

	for k, v := range person {
		fmt.Println("Key -> ", k, " - value -> ", v)
	}
}
