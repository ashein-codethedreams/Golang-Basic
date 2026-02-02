package main

import "fmt"

func main() {
	age := 26
	name := "codethedreams"
	//printting
	fmt.Print("Hello ,")
	fmt.Print("World \n")
	fmt.Print("new line! \n")

	//print line
	fmt.Println("Hello world !")
	fmt.Println("new line ...")
	fmt.Println("I am", name, ",age of", age)

	//print formatting %_ =format specifier
	fmt.Printf("my age is %v and name is %v \n", age, name)
	fmt.Printf("age is type of %T \n", age)
	fmt.Printf("my score is %f points! \n", 89.62)    //89.620000
	fmt.Printf("my score is %0.1f points! \n", 89.62) //89.6

	//Sprint (save formatted strings)
	var str = fmt.Sprintf("my age is %v and name is %v \n", age, name)
	fmt.Println("save string is:", str)

}
