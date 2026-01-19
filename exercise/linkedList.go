package exercise

import "fmt"

// Step 1: UNDERSTAND
// Problem: We need a list where adding/removing from ends is fast (O(1))
// Operations: AddFront, AddBack, RemoveFront, RemoveBack, Search
// Think: Chain of boxes, each box points to the next one

// Step 2: VISUALIZE
// [10] -> [20] -> [30] -> nil
//  ^                ^
// head            tail

// Step 3: PICK STRUCTURE
// Singly linked list - each node points to next
// We'll track both head and tail for efficiency

// Step 4: IMPLEMENT INCREMENTALLY

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

// Add item to the front - O(1)
func (ll *LinkedList) Add(value int) {
	newNode := &Node{Value: value, Next: ll.Head}

	ll.Head = newNode

	// If list was empty, tail should also point to this node
	if ll.Tail == nil {
		ll.Tail = newNode
	}

	ll.Size++
}

// Add item to the back - O(1)
func (ll *LinkedList) AddBack(value int) {
	newNode := &Node{Value: value, Next: nil}

	// If list is empty
	if ll.Tail == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}

	ll.Size++
}

// Remove from front - O(1)
func (ll *LinkedList) Remove() (int, bool) {
	if ll.Head == nil {
		return 0, false
	}

	value := ll.Head.Value
	ll.Head = ll.Head.Next

	if ll.Head == nil {
		ll.Tail = nil
	}

	ll.Size--

	return value, true
}

// Remove from back - O(n)
func (ll *LinkedList) RemoveBack() (int, bool) {
	if ll.Tail == nil {
		return 0, false
	}

	// If only one element
	if ll.Head == ll.Tail {
		value := ll.Head.Value
		ll.Head = nil
		ll.Tail = nil
		ll.Size--
		return value, true
	}

	// Find second-to-last node
	current := ll.Head
	for current.Next != ll.Tail {
		current = current.Next
	}

	value := ll.Head.Value
	current.Next = nil
	ll.Tail = nil
	ll.Size--
	return value, true
}

// Search for a value - O(n)
func (ll *LinkedList) Contains(value int) bool {
	current := ll.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

// Get size
func (ll *LinkedList) Sizes() int {
	return ll.Size
}

// Print the list (helper for visualization)
func (ll *LinkedList) Print() {
	if ll.Head == nil {
		fmt.Println("Empty List")
		return
	}

	current := ll.Head
	fmt.Print("List: ")
	for current.Next != nil {
		fmt.Println(current.Value)
		if current.Next != nil {
			fmt.Println("->")
		}
		current = current.Next
	}
	fmt.Println()
}
