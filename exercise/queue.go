package exercise

// Step 1: UNDERSTAND
// Problem: We need to store items where first-in is first-out (FIFO)
// Operations needed: Enqueue (add to back), Dequeue (remove from front)
// Think: Like a line at a store - first person in line is served first

// Step 2: VISUALIZE
// Queue: [1, 2, 3, 4] <- front is left, back is right
// Enqueue 5: [1, 2, 3, 4, 5]
// Dequeue: removes 1, returns [2, 3, 4, 5]

// Step 3: PICK STRUCTURE
// A slice works, but we'll track front/back for clarity

// Step 4: IMPLEMENT INCREMENTALLY
type Queue struct {
	items []int
}

// Create a new empty queue
func NewQueue() *Queue {
	return &Queue{
		items: []int{},
	}
}

// Add item to the back (enqueue)
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Remove item from the front (dequeue)
func (q *Queue) Dequeue() (int, bool) {
	// check if queue is empty
	if len(q.items) == 0 {
		return 0, false
	}

	// Get first item
	items := q.items[0]

	// Remove it (shift everything left)
	q.items = q.items[1:]
	return items, true
}

// Look at front item without removing
func (q *Queue) Peek() (int, bool) {
	// check if queue is empty
	if len(q.items) == 0 {
		return 0, false
	}

	// Get first item
	item := q.items[0]
	return item, true
}
