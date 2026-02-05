package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float32
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 25.3, "coffee": 12.65, "ham": 23},
		tip:   0,
	}
	return b
}

//format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-20v ...$%v \n", k+":", v)
		total += v
	}

	//total
	fs += fmt.Sprintf("%-20v ...$%0.2f", "total:", total)
	return fs
}
