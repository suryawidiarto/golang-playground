package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Surya"
	name[1] = "Adi"
	name[2] = "Widiarto"

	var value = [5]int{
		1, 2, 3, 4, 5,
	}

	fmt.Println(name[0])
	fmt.Println(value)
	fmt.Println(len(value))

}
