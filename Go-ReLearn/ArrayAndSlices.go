package main

import "fmt"

func main() {

	//! Array
	//[array_size]array_data_type
	grades := [3]int{97, 85, 93}
	//or like this, this will tell go to create and array with the size of its value
	grades2 := [...]int{97, 85, 93}
	//or empty array like this
	var grades3 [3]string
	fmt.Printf("%v, %T\n", grades, grades)   // [97 85 93], [3]int
	fmt.Printf("%v, %T\n", grades2, grades2) // [97 85 93], [3]int
	fmt.Printf("%v, %T\n", grades3, grades3) // [], [3]int
	grades3[0] = "93"
	grades3[2] = "90"
	fmt.Printf("%v, %T\n", grades3, grades3)        // [93  90], [3]string
	fmt.Printf("grades3 size : %v\n", len(grades3)) // grades3 size : 3

	var identityMatrix [3][3]int = [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	fmt.Printf("%v, %T\n", identityMatrix, identityMatrix) // [[1 0 0] [0 1 0] [0 0 1]], [3][3]int

	a := [...]int{1, 2, 3} // defining array
	b := a                 // will copy a and insert it into b
	b[1] = 10              // change b on index 1 to 10
	fmt.Println(a)         // [1 2 3]
	fmt.Println(b)         // [1 10 3]

	//! Slice
	//in short, slice is similiar with array and can do what array do, with a bit different behavior
	d := []int{1, 2, 3}                                              // defining slice
	e := d                                                           // will reference d to e
	e[1] = 10                                                        // change d on index 1 to 10
	fmt.Println("d -> ", d, "length -> ", len(d), "cap -> ", cap(d)) // d ->  [1 10 3] length ->  3 cap ->  3 -> as e was reference from d, and when we change e, d will changes as the slice is the same
	fmt.Println("e -> ", e, "length -> ", len(e), "cap -> ", cap(e)) // e ->  [1 10 3] length ->  3 cap -> 3

	sliceNum := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sliceNumB := sliceNum[:]                                                                          //slice of all element of sliceNum
	sliceNumC := sliceNum[3:]                                                                         //slice from index 3 element to end
	sliceNumD := sliceNum[:6]                                                                         //slice from index 0 to element 6th -> the first 6 element
	sliceNumE := sliceNum[3:6]                                                                        //slice from index 3 to to element 6th
	sliceNumE[2] = 100                                                                                // will change all sliceNum as all of them has the same array reference
	fmt.Println("sliceNum -> ", sliceNum, "| len -> ", len(sliceNum), "| cap -> ", cap(sliceNum))     // sliceNum ->  [1 2 3 4 5 100 7 8] | len ->  8 | cap ->  8
	fmt.Println("sliceNumB -> ", sliceNumB, "| len -> ", len(sliceNumB), "| cap -> ", cap(sliceNumB)) // sliceNumB ->  [1 2 3 4 5 100 7 8] | len ->  8 | cap ->  8
	fmt.Println("sliceNumC -> ", sliceNumC, "| len -> ", len(sliceNumC), "| cap -> ", cap(sliceNumC)) // sliceNumC ->  [4 5 100 7 8] | len ->  5 | cap ->  5
	fmt.Println("sliceNumD -> ", sliceNumD, "| len -> ", len(sliceNumD), "| cap -> ", cap(sliceNumD)) // sliceNumD ->  [1 2 3 4 5 100] | len ->  6 | cap ->  8
	fmt.Println("sliceNumE -> ", sliceNumE, "| len -> ", len(sliceNumE), "| cap -> ", cap(sliceNumE)) // sliceNumE ->  [4 5 100] | len ->  3 | cap ->  5

	builtSlice := make([]int, 8, 100)
	fmt.Println("builtSlice -> ", builtSlice, "| len -> ", len(builtSlice), "| cap -> ", cap(builtSlice)) // builtSlice ->  [0 0 0 0 0 0 0 0] | len ->  8 | cap ->  8

	reSlice := []int{}
	fmt.Println("reSlice -> ", reSlice, "| len -> ", len(reSlice), "| cap -> ", cap(reSlice)) // reSlice ->  [] | len ->  0 | cap ->  0
	reSlice = append(reSlice, 1)
	fmt.Println("reSlice -> ", reSlice, "| len -> ", len(reSlice), "| cap -> ", cap(reSlice)) // reSlice ->  [1] | len ->  1 | cap ->  1
	reSlice = append(reSlice, []int{2, 3, 4, 5}...)                                           //we can use some spread like operator in js with go
	fmt.Println("reSlice -> ", reSlice, "| len -> ", len(reSlice), "| cap -> ", cap(reSlice)) // reSlice ->  [1 2 3 4 5] | len ->  5 | cap ->  6

	reSlice = append(reSlice[:2], reSlice[3:]...)                                             // will remove index 3, but it will affect the entire reSlice as the reference is same. ensure that the behavior is working safely
	fmt.Println("reSlice -> ", reSlice, "| len -> ", len(reSlice), "| cap -> ", cap(reSlice)) // reSlice ->  [1 2 4 5] | len ->  4 | cap ->  6
}
