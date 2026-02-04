package main

import "fmt"

func updateName(x string) string {
	x = "Code the dreams"
	return x
}

func updateMenu(y map[string]float32) {
	y["coffee"] = 4.6
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "Nine"
	name = updateName(name)
	fmt.Println(name)

	// group B types -> slices, maps, functions
	menu := map[string]float32{
		"lipo":  3.2,
		"shark": 2.7,
	}
	updateMenu(menu) //merge menu
	fmt.Println(menu)
}
