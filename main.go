package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")
	mybill.addItem("onion soup", 4.50)
	mybill.addItem("pizza", 23.76)

	mybill.updateTip(10)
	fmt.Println(mybill.format())
}
