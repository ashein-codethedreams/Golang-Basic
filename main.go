package main

import "fmt"

func updateName(n *string) {
	*n = "luffy"
}
func main() {
	name := "zed"

	// & gets the memory address of the value (pointer)
	// * gets the value at the specified memory address
	m := &name //pointer

	//fmt.Println("memory address:", m) //0x...
	//fmt.Println("value at memory address:", *m) //zed

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

}
