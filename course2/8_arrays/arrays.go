package main

import "fmt"

func main() {
	// numbered sequence of specific length

	var a [10]int // array is declared with length and type of elements, in this case 10 integers
	
	// len is inbuilt function to get length of array
	fmt.Println(len(a)) // length of array is 10.

	// Condition 1) if we want to check the number on 0 index
	fmt.Println(a[0]) // default value of int is 0, so it will print 0
	    // condition 1 a) if we assign a value to 0 index
	a[0] = 100
	fmt.Println(a[0]) // now it will print 100

	//condition 2) if we want to print the whole array
	fmt.Println(a) // it will print the whole array with 100 at 0 index and rest 0s
	

}