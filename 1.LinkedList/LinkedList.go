package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

type LinkedListInterface interface {
	insert(data interface{})
	display()
	delete(data interface{}) bool
	reversePrint(node *Node)
}

func (ll *LinkedList) insert(data interface{}) {
	newNode := &Node{data: data}

	if ll.head == nil {
		ll.head = newNode
	} else {
		currentNode := ll.head

		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}

func (ll *LinkedList) display() {

	node := ll.head

	for node != nil {
		fmt.Printf("%v -> ", node.data) // %v to handle any data type
		node = node.next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) delete(data interface{}) bool {
	// condition - if node is empty
	if ll.head == nil {
		fmt.Println("Linked List Is Empty")
		return false
	}

	// condition - if data is in first node
	if ll.head.data == data {
		ll.head = ll.head.next
		return true
	}

	// condition - if data is some other node
	currentNode := ll.head

	for currentNode.next != nil && currentNode.next.data != data {
		currentNode = currentNode.next
	}

	// it is not last node , now remove it
	if currentNode.next != nil {
		currentNode.next = currentNode.next.next
		return true
	} else {
		fmt.Println("Node not found")
		return false
	}
}

func (ll *LinkedList) reversePrint(node *Node) {
	if node == nil {
		return
	}

	ll.reversePrint(node.next)
	fmt.Printf("%v -> ", node.data)
}

func main() {
	ll := &LinkedList{}

	ll.insert(10)
	ll.insert(20)
	ll.insert("Hello")
	ll.insert(30)

	ll.display()

	ll.reversePrint(ll.head)
	fmt.Println("nil")

	ll.delete(30)
	ll.delete("Hello")
	ll.display()
}
