package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hey there friends!"

	// fmt.Println(strings.Contains(greeting, "hey")) //boolean
	// fmt.Println(strings.ReplaceAll(greeting, "hey", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "th")) //4
	fmt.Println(strings.Split(greeting, " "))

	// the original phase value is unchanged
	fmt.Println("Original string value is =", greeting)

	//sort
	ages := []int{23, 65, 17, 54, 41}
	sort.Ints(ages) //changed original value
	fmt.Println(ages)

	index := sort.SearchInts(ages, 65)
	fmt.Println(index)

	names := []string{"mario", "bunny", "lumos", "regnok"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bunny"))
}
