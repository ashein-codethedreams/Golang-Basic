package main

import "fmt"

func main() {

	//string
	var nameOne string = "Nine"
	var nameTwo = "John"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Lucus"
	nameThree = "Elen"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Dumbili" //short inside func
	fmt.Println(nameFour)

	//int
	var ageOne int = 26
	var ageTwo = 17
	ageThree := 21
	fmt.Println(ageOne, ageTwo, ageThree)

	//bit & memory
	// var numOne int8 = 25 //8bit
	// var numTwo int8 = -128
	// var numThree uint16 = 256

	// var scoreOne float32 = 21.22
	// var scoreTwo float64 = 34534534534534.7

}
