package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	//container/list
	data := list.New()

	data.PushFront("1")
	data.PushBack("3")
	data.PushBack("4")

	fmt.Println(data.Front(), data.Back())

	//loop back to front
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

	fmt.Println("---------------")

	//loop front to back
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	//container/ring
	rData := ring.New(5)
	//insert data to container ring
	for i := 0; i < rData.Len(); i++ {
		rData.Value = "Value " + strconv.FormatInt(int64(i), 10)
		rData = rData.Next()
	}

	//loop through
	rData.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
