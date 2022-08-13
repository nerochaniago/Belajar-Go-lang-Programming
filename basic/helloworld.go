package main

import "fmt"

func main() {
	//fmt.Println("hello world")
	var name string
	name = "Nero Chaniago"
	fmt.Println(name)

	//atau
	country := "Indonesia" //untuk deklarasi awal
	fmt.Println(country)

	//atau
	var (
		firstName = "Nero Chaniago"
		title     = "junior Programmer"
	)

	fmt.Println(firstName)
	fmt.Println(title)

	//const
	const firstName_2 string = "Nero"
	const lastName_2 = "Chaniago"
}
