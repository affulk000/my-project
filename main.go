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
	stack := exercise.NewStack()

	fmt.Println("=== Testing Stack ===")

	// Test 1: Push items
	fmt.Println("\n1. Pushing items: 10, 20, 30")
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	// fmt.Printf(" Stack push: %d\n", stack.Size())

	// Test: Peek at top
	if top, ok := stack.Top(); ok {
		fmt.Printf("2. Top item (peek): %d\n", top)
	}

	// Test 3: Pop items
	fmt.Println("\n3. Popping items:")
	for !stack.Empty() {
		if pop, ok := stack.Pop(); ok {
			fmt.Printf(" Remove item (pop) %d\n", pop)
		}
	}

}
