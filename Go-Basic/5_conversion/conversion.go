package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64) // not enough size and will repeating to its maximum value to minimum value

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Surya"
	// get char will turn it into byte
	var char byte = name[0]

	var toString = string(char) // convert byte to string

	fmt.Println(name)
	fmt.Println(char)
	fmt.Println(toString)

	//int32 -> int64
	//int64 -> int32 | beware not enough size, so it will repeat
}
