package main

//documentation of go built in package
//? https://pkg.go.dev/std
import (
	"fmt"
	"strconv"
)

// Define variable at package scope
var valPackage = 86

// Another way by decalring variable block
var (
	actorName    string = "Surya Widiarto"
	companion    string = "Adi"
	doctorNumber int    = 9
	season       int    = 11
)

var (
	book string = "Books"
)

// redeclaring variable behavior
var redeclare int = 27

// if the naming use PascalCase it tell that the variable is on Global scope and used for export outside of package
var GlobalScope int = 100

// if the naming use camelCase and at package position it tell that the variable is on package scope and used pn internal package
var packageScope int = 100

func main() {

	//we can use it like this
	//conversion way int > int64 > float32
	//but not like this, as the lower dataType can handle the maximum value of higher dataType and we will lose data
	//! This will have a probability of losing data information | float32 > int64
	//! var firstVal float32 = 100.55
	//! secondVal = int(firstVal) // the value will be converted to 100 and we losing the information value of .55
	var firstVal int = 100
	fmt.Printf("%v, %T\n", firstVal, firstVal) // 100, int

	var secondVal float32
	secondVal = float32(firstVal)                //variable conversion
	fmt.Printf("%v, %T\n", secondVal, secondVal) // 100, float32

	var convertNumber int = 155
	var convertedString string
	//convert int to string
	convertedString = strconv.Itoa(convertNumber)
	fmt.Printf("%v, %T\n", convertedString, convertedString) // 155, string

	//if the naming use camelCase and at block position it tell that the variable is on block scope
	var blockScope int = 100
	fmt.Println(GlobalScope, blockScope)

	fmt.Println(redeclare) // will print 27

	//if there's different level redeclaring variable, go can accept it and will use the redeclared variable on func scope rather than the package scope
	//this will be called shadowing
	var redeclare int = 32

	//if we try to redeclare it here, it will throw error
	// [ERR] no new variables on left side of
	//redeclare := 1

	fmt.Println(redeclare) // will print 32

	//How to define variable
	//1st Way
	var val1 int
	val1 = 42

	fmt.Println(val1)

	//2nd Way
	var val2 int = 1
	fmt.Println(val2)

	//3rd Way
	val3 := 4
	fmt.Println(val3)

	//4th Way
	fmt.Println(valPackage)
}
