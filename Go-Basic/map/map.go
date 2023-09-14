package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Surya",
		"address": "Purwokerto",
	}

	person["title"] = "Developer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)

	book["title"] = "Belajar Golang"
	book["author"] = "Surya"
	book["publisher"] = "S Media"

	fmt.Println(book)
	delete(book, "publisher")
	fmt.Println(book)
}
