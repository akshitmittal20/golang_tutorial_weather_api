package main

import (
	"fmt"
)

func main() {
	mystring := "Akshit"
	fmt.Printf("The first word of string is - %v %T \n", mystring[1], mystring[1])
	for i, v := range mystring {
		fmt.Println(i, v) //prints position (index) and then ascii value of character
	}
	fmt.Printf("\n The length of my string is %v", len(mystring))

	for _, char := range mystring {
		fmt.Println("\n", string(char)) //prints position (index) and then ascii value of character
	}

	var MyRune = 'A'
	fmt.Printf("\n my Rune = %v \n", MyRune)

	mylongstring := []string{"a", "b", "c", "d", "e"}
	for i, v := range mylongstring {
		fmt.Println(i, v)
		fmt.Printf("%v %v \n", i, v)

	}
}
