package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]

	fmt.Println("months -> ", months)
	fmt.Println("slice1 -> ", slice1)
	fmt.Println("len -> ", len(slice1))
	fmt.Println("Cap -> ", cap(slice1))

	months[5] = "Dirubah"

	fmt.Println("slice1 -> ", slice1)

	//slice will also change the main array
	slice1[0] = "May-Dirubah"
	fmt.Println("months -> ", months)

	//cap 2, pointer 10, length 2
	slice2 := months[10:]

	//cap 10, pointer 2, length 2
	//slice2 := months[2:4]

	fmt.Println("slice2 -> ", slice2)

	//months[10:] -> this slice3 will create new array because slice not have enough cap from 2, need to add 1 to be cap (3)
	//months[2:4] -> this slice3 will not create new array because slice still has enough cap, and will changes parent array if there's a change on child slice
	slice3 := append(slice2, "Surya")

	fmt.Println("slice3 -> ", slice3)

	slice3[1] = "Bukan-December"
	fmt.Println("slice3 -> ", slice3)
	fmt.Println("slice2 -> ", slice2)
	fmt.Println("months -> ", months)

	//create fresh slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Surya"
	newSlice[1] = "Widiarto"

	fmt.Println(newSlice)

	//when copy'ing slice, make sure the length is same, if not the data will be terpotong
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	//this is array
	iniArray1 := [...]int{1, 2, 3, 4, 5}
	iniArray2 := [5]int{1, 2, 3, 4, 5}
	//this is slice
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray1, " - ", iniArray2, " - ", iniSlice)

}
