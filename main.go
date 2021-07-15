package main

import "fmt"

func main() {
	myBill := newBill("my bill")
	fmt.Println(myBill.format())
}
