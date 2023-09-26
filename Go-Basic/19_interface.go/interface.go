//interface is abstract data type
//interface in go like interface in java
//interface is like a contract, where the inheritance must have the same interface method

package main

import "fmt"

type HasName interface {
	GetName() string
}

//define interface, anyone that want to use this method, need to inherit interface condition
func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

//sample inherit interface
type Person struct {
	Name string
}

//fulfill interface condition by having GetName() that return string
func (person Person) GetName() string {
	return person.Name
}

//2nd sample inherit interface
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

//empty interface
//with interface{} we can return all datatype or consume all data type
type Apapun interface {
}

//another declaration interface{}
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

//as parameter
func Ups2(i int, something interface{}) {

}

func main() {

	var cust Person
	cust.Name = "Surya"
	//insert struct as interface, but it need fulfill all condition on interface, if not fulfilled it will throw error
	SayHello(cust)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)

	//data wil be interface{} data type and can assigned different data type as their value
	var data interface{} = Ups(1)
	fmt.Println(data)
	data = Ups(2)
	fmt.Println(data)
	data = Ups(3)
	fmt.Println(data)
}
