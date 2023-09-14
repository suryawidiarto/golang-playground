package main

import "fmt"

func main() {
	name := "Surya"

	switch name {
	case "Surya":
		fmt.Println("Hi, Surya")
	case "Adi":
		fmt.Println("Hi, Adi")
	case "Widiarto":
		fmt.Println("Hi, Widiarto")
	default:
		fmt.Println("Default")
	}

	//switch with short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Hi, Surya")
	case false:
		fmt.Println("Hi, Adi")
	}

	length2 := len(name)
	switch {
	case length2 > 4:
		fmt.Println("Suryaa")
	case length2 > 5:
		fmt.Println("Kebanyakan")
	}
}
