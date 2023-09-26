package main

import "fmt"

/**
struct is template data that used for merging 0 or more data type as single entity
simply, struct is bunch of field
*/

//stuct is data type
type Customer struct {
	Name, Address string
	Age           int
	Location      map[string]string
	ArrayLoc      [5]int
	SliceLoc      []int
}

//struct method/function
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (c Customer) sayHoo() {
	fmt.Println("Hooo from", c.Name)
}

func main() {
	//1st way define struct
	var cust Customer
	cust.Name = "Surya"
	cust.Address = "JKT"

	fmt.Println(cust)

	//2nd way define struct
	cust2 := Customer{
		Name:    "Joko",
		Address: "TGR",
		Age:     40,
	}

	fmt.Println(cust2)

	//3rd way define struct, but if we want to define it like this, we need to input all data on struct
	cust3 := Customer{"Budi", "JKT", 39, make(map[string]string), [5]int{1, 2, 3, 4, 5}, make([]int, 5)}

	fmt.Println(cust3)

	//struct method/function
	cust3.sayHello("Jono")
	cust3.sayHoo()

}
