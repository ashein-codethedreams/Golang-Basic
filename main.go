package main

import (
	"fmt"
)

func main() {
	age := 17

	fmt.Println(age >= 20)
	fmt.Println(age <= 20)
	fmt.Println(age == 17)
	fmt.Println(age != 35)

	if age > 18 {
		fmt.Println("You are adult")
	} else if age > 25 {
		fmt.Println("too old man!")
	} else {
		fmt.Println(("you just a kid"))
	}

	names := []string{"mason", "ugarte", "zide", "dean"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("break at pos", index)
			break
		}
		fmt.Printf("the value at index %v is %v \n", index, value)
	}
}
