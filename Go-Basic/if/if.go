package main

import "fmt"

func main() {
	name := "Suryaa"

	if name == "Surya" && name != "" {
		fmt.Println("HALLO SURYA")
	} else if name == "Suryaa" {
		fmt.Println("SURYAA")
	} else {
		fmt.Println("FALSE")
	}

	//short statemen on if
	if length := len(name); length > 5 {
		fmt.Println("YUHUUU")
	}
}
