package main

import "fmt"

func main(){
    // type of passing string variable in go

	// type 1
	var mes = "Hello, Go!"
	fmt.Println(mes)
	
	// type 2
	var name string = "Pankaj"
	fmt.Println(name)

	// type 3
	message := "Hello"
	fmt.Println(message)

	// Same we can do for other data types as well like int, float, bool etc.

	// example for int
	var age int = 30
	fmt.Println(age)

	var age3 = 35
	fmt.Println(age3)

	age2 := 25
	fmt.Println(age2)

	var age4 int
	age4 = 50
	fmt.Println(age4)


	//example for float
	var price float64 = 19.99
	fmt.Println(price)

	var price2 = 29.99
	fmt.Println(price2)

	price3 := 39.99
	fmt.Println(price3)

	var price4 float64
	price4 = 49.99
	fmt.Println(price4)

	// example for bool
	var isAvailable bool
	isAvailable = true
	fmt.Println(isAvailable)

	var isAvailable2 = false
	fmt.Println(isAvailable2)
	
	isAvailable3 := true
	fmt.Println(isAvailable3)
	
	var demo bool = true
	fmt.Println(demo)
}