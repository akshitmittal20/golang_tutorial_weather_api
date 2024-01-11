package main

import "fmt"

type PetrolEngine struct {
	petrolspent      uint
	distsancecovered uint
	owner
}

type owner struct {
	name string
}

func (pe PetrolEngine) calculatemileage() uint {
	mileage := pe.distsancecovered / pe.petrolspent
	return mileage
}

func (cng cngengine) calculatemileage() uint {
	mileage := cng.gasusageperkm / cng.gallonused
	return mileage
}

type engine interface { //we can use this function to use calculatemethod anywhere now, without using name of structure berfore it
	calculatemileage() uint
}

func (Pe PetrolEngine) canyoumakeit(distanceoftrip uint, currentpetrolqty uint) {
	if distanceoftrip <= currentpetrolqty*Pe.calculatemileage() {
		fmt.Println("You can make it there")
	} else {
		fmt.Println("You cant make it sorry!!")
	}
}

type cngengine struct {
	gasusageperkm uint
	gallonused    uint
}

func main() {
	var details PetrolEngine = PetrolEngine{4, 34, owner{"akshit"}}
	fmt.Printf("Average is %v \n", details.calculatemileage())

	var dist uint
	var qty uint
	var err error

	fmt.Println("Enter distance you have to cover")
	_, err = fmt.Scan(&dist)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Enter current petrol qty you have")
	_, err = fmt.Scan(&qty)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	details.canyoumakeit(dist, qty)
}
