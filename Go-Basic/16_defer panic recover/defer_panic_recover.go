package main

import "fmt"

func logging() {
	fmt.Println("Done Calling Function")
}

func endApp() {
	//recover is function that can be used to catch panic, with recover panic will be catched and the program will not stopped
	message := recover()
	if message != nil {
		fmt.Println("There's error occured ->", message)
	}
	fmt.Println("End Application")
}

func runApp(error bool) {
	defer endApp()
	if error {
		//panic is used for stopping program
		panic("Application Error")
	}
	fmt.Println("Application Running")
}

func runApplication(value int) {
	fmt.Println("Run Application")
	//defer will execute as last called execution after function is executed, it will also execute even if the function crashed
	defer logging()
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	//runApplication(0)
	runApp(true)

	fmt.Println("this will be not stopped because panic is catched by recover")
}
