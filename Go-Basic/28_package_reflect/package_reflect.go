package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"100"`
	Ayy  int64  // will not be validated because it has no tag required
}

func IsValidStruct(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		//check tag required = true
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			//if field value is "" return false
			if value == "" {
				return false
			}
			//otherway ternary like operator in go
			//return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

// reflect is used to check field/value on at compiled / runned time
func main() {
	sample := Sample{"Sury", 40}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)

	fmt.Println(structField.Name)                   //Name
	fmt.Println(structField.Tag.Get("required"))    //true
	fmt.Println(sampleType.Field(0).Tag.Get("max")) //100
	fmt.Println(sampleType.Field(0).Tag.Get("min")) //

	//checking the value of field with its tag
	fmt.Println("sample", sample)
	fmt.Println("IsValidStruct(sample)", IsValidStruct(sample))
	sample.Name = ""
	fmt.Println("IsValidStruct(sample)", IsValidStruct(sample))
}
