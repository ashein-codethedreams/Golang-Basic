package main

import (
	"fmt"
	"math"
)

func greetingName(n string) {
	fmt.Printf("Hello %v! \n", n)
}

func calArea(r float64) float64 {
	return math.Pi * r * r
}

func cycleNames(n []string,f func(string)){
	for _, v := range n{
		f(v)
	}
}

func main() {
	greetingName("AMM")

	areaVal := calArea(1)
	fmt.Println(areaVal)

	cycleNames([]string{"cloud","tim","lucy"},greetingName)
}
