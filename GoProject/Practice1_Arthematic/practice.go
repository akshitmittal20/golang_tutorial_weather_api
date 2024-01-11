package main

import (
	"fmt"
	"log"
	"unicode"
	"unicode/utf8"
)

func printLine() {
	fmt.Println("Hi")
}

func arthematicOpr(num1, num2 int) (res1 int, res2 int, res3 float32, res4 int) {
	sub := int(num1 - num2)
	fmt.Println("Substraction -", sub)
	// if sub >

	mul := int(num1 * num2)
	fmt.Println("multiplication : ", mul)

	div := float32(num1) / float32(num2)
	fmt.Println("Division- ", div)

	remain := int(num1 % num2)
	fmt.Println("Reainder is- ", remain)

	return sub, mul, div, remain
}

func stringOperations(Mystring string) (length1, length2, length3 int) {
	length1 = len(Mystring)
	length2 = utf8.RuneCountInString(Mystring)
	for _, char := range Mystring {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			length3++
		}
	}
	return length1, length2, length3
}

func main() {
	var num1, num2 int
	var err error

	fmt.Print("Enter Number 1: ")
	_, err = fmt.Scan(&num1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Enter Number 2: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		log.Fatal(err)
	}

	printLine()
	fmt.Println(arthematicOpr(num1, num2))
	fmt.Println(stringOperations("Hello my name is Akshit @20"))
}
