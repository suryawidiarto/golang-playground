package main

import "fmt"

func main() {
	//will create alias for string
	type NoKTP string
	type Married bool

	var noKTPSurya NoKTP = "01823012398123"
	var isMarried Married = false
	fmt.Println(noKTPSurya)
	fmt.Println(isMarried)
}
