package main

import "fmt"

func main() {
	menu := map[string]float64{ //key-string & value-float
		"pie":   34.6,
		"soda":  12.5,
		"pizza": 41,
	}
	fmt.Println(menu)
	fmt.Println(menu["pizza"])

	//lopping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//int as key type
	postalCodes := map[int]string{
		11111: "Sanchaung",
		11051: "Hlaing",
		11421: "Dagon",
	}
	fmt.Println(postalCodes)

	//update values
	postalCodes[11421] = "Dagon Myo Thit"
	fmt.Println("updated -", postalCodes)

}
