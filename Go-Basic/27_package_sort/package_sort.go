package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

// alias
type UserSlice []User

// create contract interface sort with alias UserSlice, to be able using sort on golang
func (value UserSlice) Len() int {
	return len(value)
}
func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}
func (value UserSlice) Less(i, j int) bool {
	//compare by age
	return value[i].Age < value[j].Age
}

func main() {
	users := []User{
		{"AAA", 30},
		{"BBB", 29},
		{"CCC", 12},
		{"DDD", 50},
	}

	//UserSlice(users) -> convert it to alias to prevent error because we created interface implementation on alias
	sort.Sort(UserSlice(users))
	//error
	//sort.Sort(users)

	fmt.Println(users)
}
