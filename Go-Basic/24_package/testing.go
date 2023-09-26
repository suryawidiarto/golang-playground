package main

//by standard go lang can only access file within the same package
//if we want to access other file go outside the package, we can use import

import (
	"fmt"
	// _ to ignore import if the package was not used, otherwise it will need to be used if we are importing package
	_ "testing/database"
	"testing/helper"
)

func main() {
	helper.SayHello("Sury")
	//will call by package name not file name inside package
	fmt.Println(helper.VersionApplication) // can be accessed from outside its package if the first letter is capital

	//fmt.Println(database.GetDatabase())
}
