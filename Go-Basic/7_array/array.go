package main

import "fmt"

func main() {
	//set array capacity
	var name [3]string // array in go is fixed length

	name[0] = "Surya"
	name[1] = "Adi"
	name[2] = "Widiarto"

	var value = [5]int{
		1, 2, 3, 4, 5,
	}

	fmt.Println(name[0])
	fmt.Println(value)
	fmt.Println(len(value))

	var defaultArray [10]string

	fmt.Println(len(defaultArray)) // 10 -> because the length are already fixed

}
