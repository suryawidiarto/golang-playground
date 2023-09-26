package main

import (
	"flag"
	"fmt"
	"os"
)

// os package is used for using system operation command
// https://pkg.go.dev/os

// flag package is bunch of function to parse command line argument on os.Args
// https://pkg.go.dev/flag
func main() {
	//os
	//to get arguments in golang
	//go run os_flag.go args1 args2 args3
	args := os.Args
	fmt.Println("Argument: ")
	fmt.Println(args)

	name, err := os.Hostname()

	if err == nil {
		fmt.Println("Hostname", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	goPathEnv := os.Getenv("GOPATH")
	fmt.Println(goPathEnv)

	//flag
	var host *string = flag.String("host", "localhost", "Put your database host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 0, "Put your number")

	//parse command line argument
	//put value by type'ing in command
	//go run os_flag.go -host=127.0.0 -user=admin -password=adminpassword

	//wrong input
	//go run os_flag.go -host=127.0.0 -user=admin -password=adminpassword -number=inistring
	/**
	invalid value "inistring" for flag -number: parse error
	Usage of C:\Users\SW85969X\AppData\Local\Temp\go-build1236497039\b001\exe\os_flag.exe:
	-host string
			Put your database host (default "localhost")
	-number int
			Put your number
	-password string
			Put your database password (default "root")
	-user string
			Put your database user (default "root")
	exit status 2
	*/

	//default
	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("User:", *user)
	fmt.Println("Password:", *password)
	fmt.Println("Number:", *number)

	//input argument in command line
	// Host: 127.0.0
	// User: admin
	// Password: adminpassword
	// Number: 0

	//default
	// Host: localhost
	// User: root
	// Password: root
	// Number: 0
}
