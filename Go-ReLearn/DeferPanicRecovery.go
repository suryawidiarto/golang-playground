package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//! DEFER
	fmt.Println("start")
	//defer tell go that it will be executed after the main() finish executing any other statement and skip something that has defer keyword on it at first
	// after main() finish executing, go then trying to look for defer, which then execute defer before returning
	// defer will execute in LIFO (Last In First Out)
	defer fmt.Println("middle")
	defer fmt.Println("end")

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close() // allow to define stopping behavior on end of used code, but it best not be used on for, as this will only closed the conection on end of executable code, and it there's thousand of loop with each of it opening connection, it would be bad to leave open all of the connection and only close it on end of code, and will lead to bugs
	// need to use it in the best way possible, for example use it on separate func
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)

	//? another example
	a := "start here"
	defer fmt.Println(a) // will print "start", because even if defer will execute the function at end of code it will used the same logic where it was positioned and only call it last
	a = "end here"

	//! PANIC
	//pA, pB := 1, 0
	//anw := pA / pB // will throw error
	//fmt.Println(anw)

	//? FLOW OF GO execute, func -> defer -> panic
	//? for example this was useful for stopping any opening connection when there's a panic occur
	fmt.Println("START PANIC")
	defer fmt.Println("This was Deffered")
	panicker()
	fmt.Println("END PANIC")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Goooo!!"))
	})
	errors := http.ListenAndServe(":8080", nil)
	if errors != nil {
		panic(errors.Error())
	}
}

func panicker() {
	fmt.Println("About to panicc")
	//defining anonymous func
	defer func() {
		//recover() -> will return nil if the application is panic, but if it isn't nil its going to return the error that is the cause of the application
		//recover() was used to recover from a panic and properly used in deferred function as the GO flow, func -> defer -> panic
		//current function will not attempt to continue, but higher function in call stack will continue
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err) // we need to rethrow the error so the upper func stop as well
		}
	}() // () is tell to execute it after defining
	panic("Something bad happened !!!") // will throw an error, but the error wil only for panicker()
	fmt.Println("Done panicking")
}
