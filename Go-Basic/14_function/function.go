package main

import "fmt"

//function declaration

//without parameter
func sayHello() {
	fmt.Println("sayHello() is Running")
}

//with multiple parameter
func sayHello2(param1 string, param2 string) {
	fmt.Println("Hi, ", param1, " Nice see you ", param2)
}

//with return value string
func getHello(name string) string {
	return "Hello " + name
}

//with return value more than 1 value
func getData(name2 string, age int) (string, int) {
	return name2, age
}

//with return value map
func getData2(name string, age int) []string {
	slice := []string{}
	slice = append(slice, name)
	return slice
}

//with return value map
func getData3(name string, age int) map[int]string {
	mapp := make(map[int]string)
	mapp[age] = name

	return mapp
}

// return named value
//this will add firstName, and secondName with string as their data type without declaring it first
func getCompleteName() (firstName, secondName, Thirdname string) {
	firstName = "Surya"
	secondName = "Adi"
	Thirdname = "Widiarto"

	//Go will automatically return the 3 variable
	return

	// return above will be work like this
	//return firstName, secondName, Thirdname
}

//variadic function -> function that have ...args argument / can accept multiple param like spread operator in js that became slice
func sumAll(numbers ...int) int {
	fmt.Println(numbers) // [10 10 10 10 10]

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

//function as variable value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

//function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

//with type declaration
type Filter func(string) string

func sayHelloWithFilter2(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Guguk" {
		return "..."
	} else {
		return name
	}
}

//anonymous function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//factorial for loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

//recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {

	//get each returned value from function
	r1, r2, r3 := getCompleteName()
	fmt.Println(r1, r2, r3)

	resultName2, resultAge := getData("Surya", 22)
	fmt.Println(resultName2, " - ", resultAge)

	//_ underscore will tell GO to ignore the second returned value
	resultName, _ := getData("Surya", 22)
	fmt.Println(resultName, " - ")

	result := getHello("Surya")
	fmt.Println(result)

	sayHello2("Sunny", "Surya")

	for i := 0; i < 2; i++ {
		sayHello()
	}

	fmt.Println(getData2("Surya", 30))
	fmt.Println(getData3("Surya", 50))

	fmt.Println(sumAll(10, 10, 10, 10, 10))

	sliceTest := []int{10, 2, 5, 8, 9, 22}

	//we can user variable... to spread slice
	fmt.Println(sumAll(sliceTest...))

	//function as varible value, we can assign the function into variable
	variableFunc := getGoodBye
	fmt.Println(variableFunc("Surya"))

	//function as parameter
	sayHelloWithFilter("Guguk", spamFilter)

	//or

	spamFunc := spamFilter
	sayHelloWithFilter("Guguk", spamFunc)

	//anonymous function
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("surya", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})

	//recursive function
	recursive := factorialRecursive(3)
	loop := factorialLoop(3)
	fmt.Println(recursive, loop)
}
