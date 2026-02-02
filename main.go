package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{12, 34, 23} //arr type & length fixed
	var ages = [3]int{12, 34, 23}
	names := [4]string{"timber", "elen", "Nova"}
	names[3] = "hook"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices (use arrays under the hood)
	var scores = []int{100, 56} //can vary
	scores[1] = 32
	scores = append(scores, 93)

	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[1:3] //startFrom & exclude End = 1,2
	//rangeTwo := names[1:] //include last
	rangeTwo := names[:4] // start from 0

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
}
