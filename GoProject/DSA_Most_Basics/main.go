package main

import (
	"fmt"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func swapM1(var1 int, var2 int) {
	tmp := var1
	var1 = var2
	var1 = tmp
	fmt.Printf("\n \n \n The swapped numbers by Method 1 are %v %v \n", var1, var2)
}

func swapM2(int1 int, int2 int) {
	int1 = int1 + int2
	int2 = int1 - int2
	int1 = int1 - int2
	fmt.Printf("\nThe swapped numbers by Method 2 on current numbers are %v %v \n \n", int1, int2)
}

func swapArrM1(arr []int, n int) {
	for i := 0; i < n/2; i++ {
		f := i
		l := n - 1 - i
		temp := arr[f]
		arr[f] = arr[l]
		arr[l] = temp
	}
	for i := 0; i < n; i++ {
		fmt.Printf("The swapped array by method 1 is %v \n", arr[i])
	}
	fmt.Printf("The full current array is %v \n \n", arr)
}

func swapArrM2(arr []int, n int) {
	swpdarr := arr
	for i := n - 1; i >= 0; i-- {
		fmt.Println("The swapped array by method 2 on current array is", swpdarr[i])
	}
}

func BubblesortArr(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[i] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Printf("\nThe Bubble sort on array is %v \n", arr)
}

func InsertionsortArr(arr []int, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	fmt.Printf("\nThe Insertion sort on current array is %v \n \n", arr)
}

func LinearSearch(arr []int, val int) {
	n := len(arr)
	for i := 0; i <= n-1; i++ {
		if arr[i] == val {
			fmt.Printf("Linear Search***** The position of your value %v in array is %v\n \n", val, i+1)
		}
	}
}

func BinarySearch(arr []int, val int) {
	n := len(arr)
	var l int = n
	var f int = 0
	var mid int = (l + f) / 2

	for l > f {
		if val == arr[mid] {
			fmt.Printf("Binary Search*****  ...Iteration Completed!! The position of your value %v in array is %v\n \n \n", val, mid+1)
			break
		} else if val < arr[mid] {
			fmt.Println("Binary Search*****  Iterating left")
			l = mid
		} else {
			fmt.Println("Binary Search*****  Iterating right")
			f = mid + 1
		}
		mid = (l + f) / 2
	}
}

type stack struct {
	items []int
}

func (s *stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *stack) Pop() {
	if len(s.items) == 0 {
		fmt.Println("The stack is empty", s)
	}

	item := s.items[len(s.items)-1]    //retrieving the top element of the stack
	s.items = s.items[:len(s.items)-1] //slice function to delete the last element which is retrieved
	fmt.Println(item)
}

type queue struct {
	items []int
}

func (q *queue) enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *queue) dequeue() {
	if len(q.items) == 0 {
		fmt.Println("empty queue!!")
		return
	}
	item := q.items[0]
	q.items = q.items[1:]
	fmt.Printf("\nItem to be removed from queue is %v", item)
}

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
}

func (ll *linkedList) prepend(node *node) { //we use pointers in structure reference when we want to directly modify something in the original data rather than the copy of data
	temp := ll.head
	ll.head = node
	ll.head.next = temp
	ll.length++
}

func (ll linkedList) PrintListdata() {
	if ll.length == 0 {
		fmt.Println("\nThe List is empty\n")
		return
	}
	fmt.Println("Your Linked List is - ")
	for ll.head != nil {
		fmt.Printf("%d ", ll.head.data)
		ll.head = ll.head.next
	}
}

func (ll *linkedList) deleteValuefromList(val int) {
	if ll.head == nil { //case: if list is empty
		fmt.Println("\nThe List is Empty")
		return
	}

	if ll.head.data == val { //case: if head data is same as value
		ll.head = ll.head.next
		ll.length--
		ll.PrintListdata()
		fmt.Println("\n")
		return
	}

	//for values present anywhere in the list but not in the head value. or if value is simply Not present in the list
	for ll.head.next.data != val {
		if ll.head.next.next == nil {
			fmt.Println("value is not present in the list")
			return
		}
		ll.head = ll.head.next
	}
	ll.head.next = ll.head.next.next
	ll.length--
}

func main() {

	//*******Take input from User*********
	var1, var2, n, arr, val := userInput()
	//*******Take input from User Done*********//

	//*******Call Functions********
	swapM1(var1, var2)

	swapM2(var1, var2)

	swapArrM1(arr, n)

	swapArrM2(arr, n)

	BubblesortArr(arr, n)

	InsertionsortArr(arr, n)

	LinearSearch(arr, val)

	BinarySearch(arr, val)

	StacksInputs()

	QueueInputs()

	UserInputLinkedList()

	//*******Call Functions Done********

}

func QueueInputs() {
	queue := queue{}
	queue.enqueue(543)
	queue.enqueue(643)
	queue.enqueue(5)
	queue.enqueue(80)
	fmt.Println(queue)

	queue.dequeue()
	fmt.Println(queue)
	queue.dequeue()
	fmt.Println(queue)
	queue.dequeue()
	fmt.Println(queue)
}

func StacksInputs() {
	stack := stack{}
	stack.push(54)
	stack.push(78)
	stack.push(90)
	stack.push(3)
	fmt.Println(stack)

	stack.Pop()
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
}

func UserInputLinkedList() {
	var len int
	var val int
	ll := linkedList{}
	fmt.Println("Enter the length of Linked List. Minimum 2")
	_, err := fmt.Scan(&len)
	checkError(err)
	if len < 2 {
		fmt.Println("Enter Number Greater than 2")
	}
	if len >= 2 {
		for i := 0; i < len; i++ {
			fmt.Println("Enter the elments of linked list one by one")
			_, err := fmt.Scan(&val)
			checkError(err)
			ll.prepend(&node{data: val})
		}
		ll.PrintListdata()
		var value int
		fmt.Println("\nEnter the value you want to Delete from the List")
		fmt.Scan(&value)
		ll.deleteValuefromList(value)
		ll.PrintListdata()
		fmt.Println("\n")
	}

}

func userInput() (int, int, int, []int, int) {
	var var1 int
	var var2 int
	fmt.Println("Enter 2 numbers to be swapped")
	_, err := fmt.Scan(&var1, &var2)
	checkError(err)

	var n int
	fmt.Println("How many array elements you want to enter in Array")
	_, err = fmt.Scan(&n)
	checkError(err)
	arr := make([]int, n)
	fmt.Println("Enter an Array to be swapped and Sorted, each element in next line")
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&arr[i])
		checkError(err)
	}
	fmt.Printf("The array you entered is %v \n \n", arr)

	var val int
	fmt.Println("What is the array element or value you want to search position for")
	_, err = fmt.Scan(&val)
	checkError(err)

	return var1, var2, n, arr, val
}
