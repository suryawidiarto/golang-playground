package main

import "fmt"

/**
by default all variable is passed by value and not by reference, and it will be duplicated
which mean if we send variable to a function, method, or another variable, the actual data that sended is the duplicated value

pointer is more efficient memory, because it not duplicating each time we pass the variable
*/

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address Address) {
	address.Country = "Indonesia"
}

// "*" will tell go that the variable is a reference, and do not duplicating
func ChangeCountryToIndonesiaWithPointer(address *Address) {
	address.Country = "Indonesia"
}

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

//indicate that the consumable struct is from pointer and do reference and not duplicate
func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1            // "here" | copied to address2
	var address5 Address = address1 //another declaration way

	address2.City = "Bandung"

	fmt.Println("address1", address1) //address1 not changed, beacuse when we assigning "here" it pass by value and not pass by reference
	fmt.Println("address2", address2)

	//pointer is a skill that can create reference to the data location in the same memory without duplicating the existing data
	//in short, with pointer, we can create pass by reference

	//pointer
	//& indicating that address3 is reference to address1
	// & is used for creating pointer or refence to already existing data
	address3 := &address1
	var address4 *Address = &address1 // another declaration way
	address3.City = "JKT"             //will also change address1 because address3 is reference of address1
	fmt.Println("address1", address1)
	fmt.Println("address3", address3)

	fmt.Println("address4", address4)
	fmt.Println("address5", address5)

	//pointer with operator *
	// "*" will chan
	var addressN1 Address = Address{"JKT", "JABAR", "Indonesia"}
	//indicate that addressN12 in refence of addressN1
	var addressN2 *Address = &addressN1
	var addressN3 *Address = &addressN1
	var addressN4 *Address = &addressN1

	//update addressN2 will also update addressN1 value
	addressN2.City = "Bandung"

	fmt.Println("addressN1", addressN1)
	fmt.Println("addressN2", addressN2)

	//if we refedinne the refernce like this, it will not changing addressN1, and will create new memory
	addressN2 = &Address{"Malang", "JaTim", "Indonesia"}

	fmt.Println("addressN1", addressN1) // stay the same
	fmt.Println("addressN2", addressN2) // updated with Malang, Jatim, indonesia
	//if we want to change all variable that it reference to, we can use *
	*addressN3 = Address{"JJJ", "DKI JKT", "Indonesia"}

	//addressN1 will also updated as addressN3
	fmt.Println("addressN1", addressN1) // addressN1 {JJJ DKI JKT Indonesia}
	//not updated because it has different memory location after changing it in line 53
	fmt.Println("addressN2", addressN2) // addressN2 &{Malang JaTim Indonesia}
	fmt.Println("addressN3", addressN3) // addressN3 &{JJJ DKI JKT Indonesia}
	fmt.Println("addressN4", addressN4) // addressN4 &{JJJ DKI JKT Indonesia}

	//if we can create empty pointer, we can use "new" keyword
	var addressN5 *Address = new(Address) // addressN5 will reference to empty pointer
	fmt.Println("addressN5", addressN5)   // addressN5 &{  }

	//Pointer in function
	var alamat = Address{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "",
	}

	alamat2 := &alamat

	//wil not change alamat, because the data that passed to the function is the duplicated one and not refer to alamat variable
	ChangeCountryToIndonesia(alamat)
	fmt.Println("alamat", alamat) // alamat {Bandung Jawa Barat }

	//if we want to tell that it reference and not duplicating, we can use & when passing value if the data passed is existing actual data and not pointer, and use * in function definition
	ChangeCountryToIndonesiaWithPointer(&alamat)
	//just use it like this if the variable passed to the funciton is a pointer
	ChangeCountryToIndonesiaWithPointer(alamat2)
	fmt.Println("alamat", alamat) // alamat {Bandung Jawa Barat Indonesia}

	//Point in struct Method
	sury := Man{"Surya"}
	sury.Married()                      // will not change sury, because the passed data is duplicated not referenced
	fmt.Println("sury.Name", sury.Name) // sury.Name Surya

	//will change the data, because it refernced and not duplicate
	sury.MarriedPointer()
	fmt.Println("sury.Name", sury.Name) //sury.Name Mr. Surya
}
