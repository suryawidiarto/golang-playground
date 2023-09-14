package main

import (
	"fmt"
	"reflect"
)

// defining named struct as Doctor
type Doctor struct {
	number    int
	actorName string
	companion []string
}

// ? embedding in struct, create behavior like extends on java
type Animal struct {
	Name   string `required	max:"100"`
	Origin string
}

type Bird struct {
	Animal // embed Animal into Bird
	Speed  float32
	CanFly bool
}

func main() {

	//! Map
	statePopulations := make(map[string]int) // we can create map using make func
	statePopulations = map[string]int{
		"Purwokerto": 9000000,
		"Jakarta":    12312312312,
		"Semarang":   1000000,
	}

	fmt.Println(statePopulations) // map[Jakarta:12312312312 Purwokerto:9000000 Semarang:1000000]
	// get map value by key
	fmt.Println(statePopulations["Purwokerto"]) // 9000000

	statePopulations["Jogja"] = 99999999
	fmt.Println(statePopulations) // map[Jakarta:12312312312 Jogja:99999999 Purwokerto:9000000 Semarang:1000000]
	delete(statePopulations, "Jakarta")
	fmt.Println(statePopulations) // map[Jogja:99999999 Purwokerto:9000000 Semarang:1000000]

	value, check := statePopulations["Purwokerto"]
	value2, check2 := statePopulations["NONE"]
	fmt.Println(value, check)   // 0 true
	fmt.Println(value2, check2) // false

	newMap := statePopulations // this will do a reference and not creating a new map, so when there's change on newMap, it will also change the statePopulations
	delete(newMap, "Jogja")
	fmt.Println(newMap)
	fmt.Println(statePopulations)

	//! Struct
	// create new named Doctor struct -> this really similiar like object
	aDoctor := Doctor{
		number:    3,
		actorName: "Surya Widiarto",
		companion: []string{
			"Adi",
			"Arto",
		},
	}

	fmt.Println(aDoctor)                                 // {3 Surya Widiarto [Adi Arto]}
	fmt.Println(aDoctor.actorName, aDoctor.companion[1]) // Surya Widiarto Arto

	// create anonymous struct
	//define struct     | //initializer
	bDoctor := struct{ name string }{name: "Surya"}
	cDoctor := bDoctor // these will create a copy and not a reference
	cDoctor.name = "Adi"

	fmt.Println(bDoctor) // {Surya}
	fmt.Println(cDoctor)

	bird := Bird{}
	bird.Name = "Mew"
	bird.Origin = "Indonesia"
	bird.Speed = 58
	bird.CanFly = true
	fmt.Println(bird) // {{Mew Indonesia} 58 true}

	bird2 := Bird{
		Animal: Animal{Name: "Mew2", Origin: "Jakarta"},
		Speed:  90,
		CanFly: false,
	}
	fmt.Println(bird2) // {{Mew2 Jakarta} 90 false}

	tags := reflect.TypeOf(Animal{})
	field, _ := tags.FieldByName("Name")
	fmt.Println(field.Tag)

}
