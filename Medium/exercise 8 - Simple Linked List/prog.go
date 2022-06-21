package main

import (
	"errors"
	"fmt"
)

type Element struct {
	data int
	next *Element
}
type List struct {
	head *Element
	size int
}

func (l List) String() string {
	listString := ""

	n := l.head
	for n != nil {
		listString += fmt.Sprintf("%d ", n.data)
		n = n.next
	}

	return "List looks: " + listString
}

func New(nums []int) *List {
	l := &List{}

	for i := range nums {
		l.Push(nums[i])
	}

	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	el := &Element{
		data: element,
		next: nil,
	}

	if l.head == nil {
		l.head = el
	} else {
		n := l.head
		for n.next != nil {
			n = n.next
		}

		n.next = el
	}

	l.size++
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return -1, errors.New("Invalid operation")
	}

	val := 0

	if l.size == 1 {
		val = l.head.data
		l.head = nil
	} else if l.size == 2 {
		val = l.head.next.data
		l.head.next = nil
	} else {
		n := l.head
		for n.next.next != nil {
			n = n.next
		}

		val = n.next.data
		n.next = nil
	}

	l.size--

	return val, nil
}

func (l *List) Array() []int {
	data := []int{}

	n := l.head
	for n != nil {
		data = append(data, n.data)
		n = n.next
	}

	return data
}

func (l *List) Reverse() *List {
	lNew := &List{}
	arr := l.Array()

	for i := len(arr) - 1; i >= 0; i-- {
		lNew.Push(arr[i])
	}

	return lNew
}

func main() {
	l := New([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(l)
	lNew := l.Reverse()
	fmt.Println(lNew)

	fmt.Println(l)
	l.Pop()
	fmt.Println(l)
	l.Pop()
	fmt.Println(l)

	l.Pop()
	fmt.Println(l)
	l.Pop()
	fmt.Println(l)

	l.Pop()
	fmt.Println(l)
	l.Pop()
	fmt.Println(l)
	fmt.Println(l.Pop())
}
