package main

import "fmt"

func main() {
	// for loop in go
	
	//while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//infinite loop
	for {
		println("Pankaj")
		break
	}

	// classic for loop
	for i := 1; i <= 8 ; i++ {

		if i == 5{		// skip the iteration when i is 5
			continue
		}
		fmt.Println(i)
	}
}