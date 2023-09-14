package main

import "fmt"

func sayHello() {
	fmt.Println("sayHello() is Running")
}

func sayHello2(param1 string, param2 string) {
	fmt.Println("Hi, ", param1, " Nice see you ", param2)
}

func getHello(name string) string {
	return "Hello " + name
}

func getData(name2 string, age int) (string, int) {
	return name2, age
}

// return named value
//this will add firstName, and secondName with string as their data type
func getCompleteName() (firstName, secondName, Thirdname string) {
	firstName = "Surya"
	secondName = "Adi"
	Thirdname = "Widiarto"

	//Go will automatically return the 3 variable
	return

	// return above will be work like this
	//return firstName, secondName, Thirdname
}

func main() {

	r1, r2, r3 := getCompleteName()
	fmt.Println(r1, r2, r3)

	resultName2, resultAge := getData("Surya", 22)
	fmt.Println(resultName2, " - ", resultAge)

	//_ underscore will tell GO to ignore
	resultName, _ := getData("Surya", 22)
	fmt.Println(resultName, " - ")

	result := getHello("Surya")
	fmt.Println(result)

	sayHello2("Sunny", "Surya")

	for i := 0; i < 2; i++ {
		sayHello()
	}
}
