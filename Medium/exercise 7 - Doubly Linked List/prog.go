package main

import (
	"errors"
	"fmt"
)

type List struct {
	head, tail *Node
}

type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

func (l List) String() string {
	var result string

	node := l.head
	for node != nil {
		result += fmt.Sprintf("%v ", node.Value)
		node = node.next
	}

	return result
}

func NewList(args ...interface{}) *List {
	list := &List{}

	for i := range args {
		list.Push(args[i])
	}

	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	node := &Node{
		Value: v,
		next:  nil,
		prev:  nil,
	}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.First()
		l.head.prev = node
		l.head = node
	}

}

func (l *List) Push(v interface{}) {
	node := &Node{
		Value: v,
		next:  nil,
		prev:  nil,
	}

	if l.First() == nil {
		l.head = node
	} else {
		node.prev = l.tail
		l.tail.next = node
	}

	l.tail = node
}

func (l *List) Shift() (interface{}, error) {
	if l.First() == nil {
		return 0, errors.New("Invalid operation")
	}

	popElem := l.First()

	if l.First().Next() == nil {
		l.head = nil
		l.tail = nil
		return popElem.Value, nil
	}

	l.head = l.First().Next()
	l.head.prev = nil

	return popElem.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.First() == nil {
		return 0, errors.New("Invalid operation")
	} else if l.First() == l.Last() {
		popElem := l.Last()

		l.head = nil
		l.tail = nil

		return popElem.Value, nil
	}

	popElem := l.Last()

	l.tail = l.Last().Prev()
	l.tail.next = nil

	return popElem.Value, nil
}

func (l *List) Reverse() {
	l.head, l.tail = l.Last(), l.First()

	node := l.First()
	for node != nil {
		node.next, node.prev = node.Prev(), node.Next()
		node = node.Next()
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func main() {
	l := NewList(2, 3)
	fmt.Println(l)
	l.Unshift(1)
	fmt.Println(l)
	l.Push(4)
	fmt.Println(l)
	l.Shift()
	fmt.Println(l)
	l.Shift()
	fmt.Println(l)
	l.Pop()
	fmt.Println(l)
	l.Pop()
	fmt.Println(l)
	l.Unshift(8)
	fmt.Println(l)
	l.Push(7)
	fmt.Println(l)
	l.Unshift(9)
	fmt.Println(l)
	l.Push(6)
	fmt.Println(l)
}
