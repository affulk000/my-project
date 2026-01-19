package main

import (
	"fmt"
	"main/exercise"
)

func main() {
	// var manager = exercise.Manager{}
	// manager.AddEmployee(exercise.Employee{ID: 1, Name: "John", Age: 30, Salary: 5000})
	// manager.AddEmployee(exercise.Employee{ID: 2, Name: "Jane", Age: 25, Salary: 4000})
	// fmt.Println(manager.GetAverageSalary())
	// fmt.Println(manager.GetEmployeeByID(1))
	//
	// stack := exercise.NewStack()

	// fmt.Println("=== Testing Stack ===")

	// // Test 1: Push items
	// fmt.Println("\n1. Pushing items: 10, 20, 30")
	// stack.Push(10)
	// stack.Push(20)
	// stack.Push(30)
	// // fmt.Printf(" Stack push: %d\n", stack.Size())

	// // Test: Peek at top
	// if top, ok := stack.Top(); ok {
	// 	fmt.Printf("2. Top item (peek): %d\n", top)
	// }

	// // Test 3: Pop items
	// fmt.Println("\n3. Popping items:")
	// for !stack.Empty() {
	// 	if pop, ok := stack.Pop(); ok {
	// 		fmt.Printf(" Remove item (pop) %d\n", pop)
	// 	}
	// }
	//

	list := exercise.NewLinkedList()

	fmt.Println("=== Testing Linked List ===")

	// Test 1: Add to front
	fmt.Println("\n1. Adding to front: 30, 20, 10")
	list.Add(30)
	list.Add(20)
	list.Add(10)
	list.Print()
	fmt.Printf("   Size: %d\n", list.Sizes())

	// Test 2: Add to back
	fmt.Println("\n2. Adding to back: 40, 50")
	list.AddBack(40)
	list.AddBack(50)
	list.Print()

	// Test 4: Remove from front
	fmt.Println("\n4. Removing from front:")
	if val, ok := list.Remove(); ok {
		fmt.Printf("   Removed: %d\n", val)
	}
	list.Print()
}
