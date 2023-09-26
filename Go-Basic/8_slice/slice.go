package main

import "fmt"

func main() {
	//spread operator like in go
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

	//array[low:high] -> create slice from index low until index before high in array
	//array[low:] -> create slice from index low until last index in array
	//array[:high] -> create slice from index 0 until index before high in array
	//array[:] -> create slice from index 0 until last index in array

	//slice create reference to array, so if we update value in slice it will affect the array and the other slice that reference the same array
	var slice1 = months[4:7]

	//len -> size filled value in slice
	//Cap -> total size from sliced index to last index in array

	fmt.Println("months -> ", months)   //[January February March April May June July August September October November December]
	fmt.Println("slice1 -> ", slice1)   //[May June July]
	fmt.Println("len -> ", len(slice1)) //3 -> get length slice
	fmt.Println("Cap -> ", cap(slice1)) //8 -> get capacity slice

	months[5] = "Dirubah"

	//if we change referenced array, slice will also change
	fmt.Println("slice1 -> ", slice1) //[May Dirubah July]

	//if we change slice it will also change the referenced array
	slice1[0] = "May-Dirubah"
	// will update months May to May-Dirubah
	fmt.Println("months -> ", months) //[January February March April May-Dirubah Dirubah July August September October November December]

	//cap 2, pointer 10, length 2
	slice2 := months[10:]

	//cap 10, pointer 2, length 2
	//slice2 := months[2:4]

	fmt.Println("slice2 -> ", slice2) //[November December]

	//months[10:] -> this slice3 will create new array because slice not have enough cap from 2, need to add 1 to be cap (3)
	//months[2:4] -> this slice3 will not create new array because slice still has enough cap, and will changes parent array if there's a change on child slice
	//[November December Surya]
	slice3 := append(slice2, "Surya") // add new data to slice based on the capacity of referenced array, if slice has enough capacity then it will add new data to slice referenced array (existing array)
	//if capacity is not enough, it will create new array, and it will not affect the referenced slice array

	fmt.Println("slice3 -> ", slice3) //[November December Surya]

	slice3[1] = "Bukan-December"
	fmt.Println("slice3 -> ", slice3) //[November Bukan-December Surya]
	fmt.Println("slice2 -> ", slice2) //[November December] // not affected as- slice3 has create new array because when appending it has not enough cap
	fmt.Println("months -> ", months) //[January February March April May-Dirubah Dirubah July August September October November December] //the referenced array also not affected

	//create empty slice
	//make(arrayType, len, cap)
	newSlice := make([]string, 2, 5) // create new slice

	newSlice[0] = "Surya"
	newSlice[1] = "Widiarto"

	fmt.Println(newSlice) // [Surya Widiarto]

	//when copy'ing slice, make sure the length is same, if not the data will be capped
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	//this is array
	iniArray1 := [...]int{1, 2, 3, 4, 5}
	iniArray2 := [5]int{1, 2, 3, 4, 5}

	//this is slice
	iniSlice := []int{1, 2, 3, 4, 5}
	iniSlice2 := []int{}

	fmt.Println(iniArray1, " - ", iniArray2, " - ", iniSlice, " - ", iniSlice2) //[1 2 3 4 5]  -  [1 2 3 4 5]  -  [1 2 3 4 5]

}
