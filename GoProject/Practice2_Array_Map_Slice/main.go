package main

import (
	"fmt"
)

func main() {
	var arr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println("array 1 - ", arr)

	var arr2 []int32 = []int32{1, 2, 3, 4, 3}
	fmt.Println("array2 - ", arr2)

	arr3 := []int32{543, 124, 2365, 3645}
	fmt.Println("arra 3 -", arr3)

	arr4 := [...]int32{3, 4, 6, 1, 1, 4, 53, 124}
	fmt.Println("array4 - ", arr4)

	arr5 := []int32{3, 4, 6, 1, 1, 34}
	fmt.Println("array 5 - ", arr5, len(arr5), cap(arr5))

	arr5 = append(arr5, 23, 234, 254634)
	fmt.Println("array5 after append - ", arr5, len(arr5), cap(arr5))

	arr3 = append(arr5, arr3...)
	fmt.Println("appending array5 and array3- ", arr3, len(arr3), cap(arr3))

	arr6 := make([]int32, 3, 4)
	fmt.Println("Array 6 - ", arr6)

	//map is userd to to store ky value pair, unlie arrays, they can use json
	var mymap map[string]uint32 = make(map[string]uint32) //initialise a map
	fmt.Println(mymap)

	map1 := map[string]uint16{"Ana": 20, "Aku": 15}
	fmt.Println(map1["Aku"])
	fmt.Println(map1["Ana"])

	map2 := map[string]uint16{"Radha": 200, "shyam": 150, "Mor": 100}
	fmt.Println(map2)

	delete(map2, "Mor")
	fmt.Println(map2)

	age, ok1 := map2["Radha"] //find age in the jsnon file
	//age, ok1 := map2["Ram"]
	if ok1 {
		//fmt.Println("The age for Ram is", age)
		fmt.Println("The age Radha is - ", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name, age := range map2 {
		fmt.Printf("Name- %v , Age- %v \n", name, age)
	}

	for index, value := range arr2 {
		fmt.Printf("Index %v  Value %v \n", index, value)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
