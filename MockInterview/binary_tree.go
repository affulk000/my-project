package mockinterview

type Node struct {
	Value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) insert(value int) {
	newNode := &Node{Value: value}

	if t.root == nil {
		t.root = newNode
		return
	}

	queue := []*Node{t.root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.left == nil {
			current.left = newNode
			return
		} else {
			queue = append(queue, current.left)
		}

		if current.right == nil {
			current.right = newNode
			return
		} else {
			queue = append(queue, current.right)
		}
	}
}

func search(node *Node, target int) bool {
	if node == nil {
		return false
	}
	if node.Value == target {
		return true
	}

	return search(node.left, target) || search(node.right, target)
}
