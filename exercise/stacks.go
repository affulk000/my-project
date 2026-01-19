package exercise

// Step 1: UNDERSTAND
// Problem: We need to store items where last-in is first-out (LIFO)
// Operations needed: Push (add), Pop (remove), Peek (look at top)
// Think: Like a stack of plates - you add/remove from the top

// Step 2: VISUALIZE
// Stack: [1, 2, 3, 4] <- top is on the right
// Push 5: [1, 2, 3, 4, 5]
// Pop: removes 5, returns [1, 2, 3, 4]

// Step 3: PICK STRUCTURE
// A slice works perfectly for a stack in Go

// Step 4: IMPLEMENT INCREMENTALLY
type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: []int{},
	}
}

// Next: Add Add item data (Push)
func (s *Stack) Push(items int) {
	s.items = append(s.items, items)
}

// then: Remove items data (Pop)
func (s *Stack) Pop() (int, bool) {
	// Checking the case edge first
	if len(s.items) == 0 {
		return 0, false
	}

	// Get the last items
	index := len(s.items) - 1
	items := s.items[index]

	// remove the items
	s.items = s.items[:index]

	return items, true
}

// Finally: Look at top without removing (Peek)
func (s *Stack) Top() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

// Helper: Check if empty
func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

// Helper: Get size
func (s *Stack) Size() int {
	return len(s.items)
}
