package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch

	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	// multiple condition switch
 
	currentTime := time.Now()
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its weekend")
	default:
		fmt.Println("its a work day")
		fmt.Println("Current time:", currentTime)
	}

	// type switch
	whoAmI := func (i interface{}) {	 	// Important syntax to use interface{} to accept any type of value
		switch t := i.(type){
		case int:
			fmt.Println("its a integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Printf("i dont know type ", t)
		}
	}

	whoAmI(8.33)
}