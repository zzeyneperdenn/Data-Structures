package main

import (
	"arrays/internal"
	"log"
)

func main() {

	ll := internal.NewLinkedList(5)
	ll.Push(10)
	ll.Push(15)
	//ll.Head = 15
	//ll.Head.Next = 10
	//ll.Head.Next.Next = 5
	ll.InsertAfter(ll.Head, 3)
	ll.Append(9999)
	//ll.Print()

	dll := internal.NewDoublyLinkedList(1)
	dll.Push(7)
	dll.DoublyAppend(3)
	dll.DoublyInsertAfter(dll.Head.Next, 2)
	dll.Print()
}

func RunArray() {

	arr := []int{3, 7, 2, 9, 12, 6, 2394, 123, 456, 45634987652, 2134}
	output := BubbleSort(arr)

	log.Printf("Result: %d", output)

}

func FindSmallest(a []int) int {
	temp := a[0]
	for _, v := range a {
		if temp > v {
			temp = v
		}
	}
	return temp
}

func FindSecondSmallest(a []int) int {
	temp := a[0]
	var currentIndex int
	for i, v := range a {
		if temp > v {
			temp = v
			currentIndex = i
		}
	}

	newArr := RemoveIndex(a, currentIndex)

	temp = newArr[0]
	for _, v := range newArr {
		if temp > v {
			temp = v
		}
	}
	return temp
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func FindGivenSmallest(a []int, n int) int {
	temp := a[0]
	for i := 0; i < n-1; i++ {
		var currentIndex int
		for i, v := range a {
			if temp > v {
				temp = v
				currentIndex = i
			}
		}

		newArr := RemoveIndex(a, currentIndex)

		temp = newArr[0]
		for _, v := range newArr {
			if temp > v {
				temp = v
			}
		}
	}
	return temp
}

func BubbleSort(a []int) []int {

	// get the first index element and compare with next element
	// compare them; if first greater then second swap elements

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}

	return a
}
