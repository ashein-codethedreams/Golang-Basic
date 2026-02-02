package main

import (
	"fmt"
)

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("i is:", i)
	// }

	names := []string{"buddy", "cliff", "tom", "sam"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println("name of each room is:", names[i])
	// }

	for index, value := range names {
		fmt.Printf("the value at index %v is: %v \n", index, value)
	}

	for _, value := range names {
		fmt.Printf("the value is: %v \n", value)
		value = "new string"
	}
	fmt.Println(names)
}
