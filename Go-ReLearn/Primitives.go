package main

import "fmt"

func main() {
	//rune declaration is using '' singlecolon
	var rune rune = 'a' // rune is a UTF-32 and a alias for int32 // how to use rune see documentation of strings.Reader#ReadRune
	//? rune := 'a' //or can also be defined like this
	fmt.Printf("%v, %T\n", rune, rune) // 97, int32

	sString := "this is stringss" // string is a UTF-8
	byteString := []byte(sString)
	fmt.Printf("%v, %T\n", byteString, byteString) // [116 104 105 115 32 105 115 32 115 116 114 105 110 103 115 115], []uint8

	var boolean bool = true
	fmt.Printf("%v, %T\n", boolean, boolean) // true, bool

	boolean2 := 1 == 1
	boolean3 := 1 == 2
	fmt.Printf("%v, %T\n", boolean2, boolean2) // true, bool
	fmt.Printf("%v, %T\n", boolean3, boolean3) // false, bool

	//default value
	var boolean4 bool
	var string string
	var number int8                                    // start from -128 - 127
	var unsignNumber uint8                             // start from 0 - 255
	fmt.Printf("%v, %T\n", boolean4, boolean4)         // false, bool
	fmt.Printf("%v, %T\n", string, string)             // , string
	fmt.Printf("%v, %T\n", number, number)             // 0, int8
	fmt.Printf("%v, %T\n", unsignNumber, unsignNumber) // 0, uint8

	//basic arithmatic math
	a := 10                      // 1010
	b := 3                       // 0011
	fmt.Printf("%v, %T\n", a, a) // 10, int
	fmt.Printf("%v, %T\n", b, b) // 3, int
	fmt.Println(a + b)           // 13
	fmt.Println(a - b)           // 7
	fmt.Println(a * b)           // 30
	fmt.Println(a / b)           // 3 -> 10 / 3 = 3.333 as the data type we use is int, it will convert to int and show 3
	fmt.Println(a % b)           // 1
	//AND OPERATOR need each of value to be true
	// 1010
	// 0011
	// false, false, true, false
	// 0, 0, 1, 0
	fmt.Println(a & b) // AND operator will get -> 0010 -> 2
	//OR OPERATOR need one of value to be true
	// 1010
	// 0011
	// true, false, true, true
	// 1, 0, 1, 1
	fmt.Println(a | b) // OR operator will get -> 1011 -> 11
	//XOR (exclusive OR) OPERATOR need one of value to be false
	// 1010
	// 0011
	// true, false, false, true
	// 1, 0, 0, 1
	fmt.Println(a ^ b) // XOR operator will get -> 1001 -> 9
	//NAND (Not AND) OPERATOR need each of value to be false
	// 1010
	// 0011
	// false, true, false, false
	// 0, 1, 0, 0
	fmt.Println(a &^ b) // NAND operator will get -> 0100 -> 8

	var A_int int = 10
	var B_int8 int8 = 3
	fmt.Printf("%v, %T\n", A_int, A_int)   // 10, int
	fmt.Printf("%v, %T\n", B_int8, B_int8) // 3, int8
	//will throw error as go can't calculate different type of numerical
	//! fmt.Println(A_int + B_int8)            // invalid operation: A_int + B_int8 (mismatched types int and int8)
	// we need to convert it to be the same data type
	fmt.Println(int8(A_int) + B_int8) // 13

}
