package main

import "fmt"

type List[T comparable] struct {
	first *Node[T]
}

type Node[T comparable] struct {
	val T
	n   *Node[T]
}

func NewList[T comparable]() List[T] {
	return List[T]{}
}

func (l *List[T]) Add(val T) {
	if l.first == nil {
		l.first = &Node[T]{val, nil}
		return
	}
	var last *Node[T] = l.first
	for last.n != nil {
		last = last.n
	}
	last.n = &Node[T]{val, nil}
}

// assume that i is valid index (0, len>, I would rather return error here otherwise
func (l *List[T]) Insert(val T, i int) {
	if i == 0 {
		n := &Node[T]{val, l.first}
		l.first = n
		return
	}

	c := 1
	for it := l.first; it != nil; it = it.n {
		if c == i {
			n := &Node[T]{val, it.n}
			it.n = n
			return
		}
		c++
	}
}

func (l *List[T]) Index(val T) int {
	i := 0
	for it := l.first; it != nil; it = it.n {
		if it.val == val {
			return i
		}
		i++
	}
	return -1
}

func main() {
	l := &List[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	for curNode := l.first; curNode != nil; curNode = curNode.n {
		fmt.Println(curNode.val)
	}

	l.Insert(300, 10)
	for curNode := l.first; curNode != nil; curNode = curNode.n {
		fmt.Println(curNode.val)
	}

	l.Add(400)
	for curNode := l.first; curNode != nil; curNode = curNode.n {
		fmt.Println(curNode.val)
	}

	l.Insert(500, 6)
	for curNode := l.first; curNode != nil; curNode = curNode.n {
		fmt.Println(curNode.val)
	}
}
