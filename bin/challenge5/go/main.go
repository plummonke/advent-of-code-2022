package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"aoc2022/utility"
)

func main() {
	p, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	text, err := utility.ReadFile(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	textArr := strings.Split(text, "\r\n")
	var containers = make(map[int]*LinkedList[rune])

	for _, each := range textArr[:5] {
		for j, c := range arrayify(each) {
			if _, ok := containers[j]; !ok && c != ' ' {
				n := NewNode[rune](c, nil)
				containers[j] = NewLinkedList[rune](n, []*Node[rune]{n})
			}

			if ll, ok := containers[j]; ok && c != ' ' {
				ll.AddNode(NewNode[rune](c, nil))
			}
		}
	}

	for _, v := range containers {
		fmt.Println(string(v.Head.GetValue()))
		fmt.Println(v.Length())
	}
}

func arrayify(line string) (out []rune) {
	for i := 0; i < len(line); i += 4 {
		var container string

		if i+4 > len(line) {
			container = line[i:]
		} else {
			container = line[i : i+4]
		}

		var c rune

		for _, char := range container {
			if char >= 'a' || char <= 'Z' {
				c = char
				break
			}
		}

		out = append(out, c)
	}

	return
}

type LinkedList[T any] struct {
	Head  *Node[T]
	Nodes []*Node[T]
}

func NewLinkedList[T any](head *Node[T], nodes []*Node[T]) *LinkedList[T] {
	return &LinkedList[T]{head, nodes}
}

func (ll *LinkedList[T]) AddNode(n *Node[T]) {
	if len(ll.Nodes) == 0 {
		ll.Head = n
		ll.Nodes = append(ll.Nodes, n)
	} else {
		ll.Head.AddNode(n)
		ll.Nodes = append(ll.Nodes, n)
	}
}

func (ll *LinkedList[T]) Length() int {
	return len(ll.Nodes)
}

func (ll *LinkedList[T]) Sever(x int) (*LinkedList[T], *LinkedList[T]) {
	if x >= ll.Length() {
		return ll, nil
	}

	front := NewLinkedList(nil, []*Node[T]{})
	end := NewLinkedList(nil, []*Node[T]{})

	var node *Node[T] = ll.Head
	for i := 0; i < x; i++ {
		front.AddNode(node)
		node = node.NextNode()
	}

	end.AddNode(node)
	node = node.NextNode()
	for node != nil {
		end.AddNode(node)
	}

	return front, end
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func NewNode[T any](value T, next *Node[T]) *Node[T] {
	return &Node[T]{value, next}
}

func (n *Node[T]) AddNode(new *Node[T]) {
	if n.Next == nil {
		n.Next = new
	} else {
		n.Next.AddNode(new)
	}
}

func (n *Node[T]) GetValue() T {
	return n.Value
}

func (n *Node[T]) NextNode() *Node[T] {
	return n.Next
}
