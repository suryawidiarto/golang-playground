package main

import (
	"fmt"
)

const shadowConst int16 = 24

// we can define const iota as like this, and go will apply the same pattern to another const for us
const (
	a = iota // iota is a counter that we can use when we're creating what are we called enumerated constant
	b        // will do the same pattern as above
	c        // will do the same pattern as above
)

const (
	cA = iota // iota will reset to 0
	cB = iota
)

// example
const (
	isAdmin = 1 << iota // we can implement basic arithmatic, bitwise, and bitshifting on const declaration as long as its not a advance function declaration
	isHeadquarter
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b, %T\n", roles, roles)                                                               // 100101, uint8
	fmt.Printf("%v, %T\n", isAdmin&roles == isAdmin, isAdmin&roles == isAdmin)                         // true, bool
	fmt.Printf("%v, %T\n", isHeadquarter&roles == isHeadquarter, isHeadquarter&roles == isHeadquarter) // false, bool

	fmt.Printf("%v, %T\n", a, a)   // 0, int
	fmt.Printf("%v, %T\n", b, b)   // 1, int
	fmt.Printf("%v, %T\n", c, c)   // 2, int
	fmt.Printf("%v, %T\n", cA, cA) // 0, int
	fmt.Printf("%v, %T\n", cB, cB) // 1, int

	const constA = 10 // we can also define const as this
	var valB int32 = 2
	// as we can see before if we calculate different type of numeric data type, go will throw an error but this will not
	// this is because go will replace constA with 10 and it do an implicit conversion, and so we can get the result
	//? fmt.Printf("%v, %T\n", 10+valB, 10+valB)
	fmt.Printf("%v, %T\n", constA+valB, constA+valB) // 12, int32

	fmt.Printf("%v, %T\n", shadowConst, shadowConst) // 24, int16
	//we can still shadow a const, so we need to be makesure it when coding
	const shadowConst int = 12
	const anotherConst = 10                            // we can also define const as this
	fmt.Printf("%v, %T\n", shadowConst, shadowConst)   // 12, int
	fmt.Printf("%v, %T\n", anotherConst, anotherConst) // 10, int

	// const need to be defined at compile time and not a run time
	const myConst int = 32
	//! will throw error
	//! const errConst float64 = math.Ceil(68.88) // math.Ceil(68.88) (value of type float64) is not constant
	fmt.Printf("%v, %T\n", myConst, myConst) // 32, int
}
