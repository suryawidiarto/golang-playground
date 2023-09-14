package main

import "fmt"

func main() {
	var name string
	var name2 = "aku"
	var age = 30
	name3 := "Suryaaa"

	var (
		firstName = "Surya"
		lastName  = "Widiarto"
	)

	const namaDepan = "Surya"
	const namaBelakang = "Surya"

	const (
		namaku1 = "Surya"
		umurku1 = 30
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)

	fmt.Println(firstName)
	fmt.Println(lastName)

	fmt.Println(name3)
	name = "Surya Adi Widiarto"

	fmt.Println(name)

	name = "Surya"

	fmt.Println(name2)
	fmt.Println(age)

}
