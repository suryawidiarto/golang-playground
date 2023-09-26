package main

import (
	"errors"
	"fmt"
)

//golang has interface that is used as contract for creating error
//in short, built in error

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	var containError error = errors.New("Ups Error")
	fmt.Println(containError.Error())

	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

}
