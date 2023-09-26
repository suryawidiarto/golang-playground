package main

import "fmt"

//type assertion is used for changing data type
//this feature frequently used when we handling empty interface{}

func random() interface{} {
	return "UpsiUpsiUlala"
}

func main() {
	result := random()
	fmt.Println("result", result) //result is interface{}
	//type assertions result.(string)
	resultString := result.(string) //cast interface{} to string with type assertion
	fmt.Println(resultString)

	//resultInt := result.(int) //cast interface{} to int --> will result panic, because returned value from random() is string, and cannot be casted to int

	//get interface{} type
	//safe measure to cast interface{} with switch statement
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown Type")
	}

}
